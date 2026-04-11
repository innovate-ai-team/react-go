package main

import (
  "encoding/json"
  "log"
  "net/http"
  "time"

  "github.com/gorilla/websocket"
  "react-go-server/mathutils"
  "react-go-server/cipher"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func main() {
  // generate a symmetric key for optional cipher responses
  serverKey, err := cipher.NewKey()
  if err != nil {
    log.Println("failed to generate cipher key:", err)
  } else {
    log.Println("cipher key (base64):", serverKey)
  }

  // store key in closure for handlers
  http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
    // attach serverKey to request context by capturing it in closure
    wsHandlerWithKey(w, r, serverKey)
  })

  log.Println("Go server listening on :8080 (ws)")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

func wsHandlerWithKey(w http.ResponseWriter, r *http.Request, serverKey string) {
  conn, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    log.Println("upgrade error:", err)
    return
  }
  defer conn.Close()

  ticker := time.NewTicker(40 * time.Millisecond)
  defer ticker.Stop()

  for t := range ticker.C {
    angle := mathutils.AngleFromTime(t)

    // Build a transform that rotates and slightly scales the model
    tr := mathutils.Transform{
      Position: mathutils.NewVec3(0, 0, 0),
      Rotation: mathutils.NewVec3(0, angle, 0),
      Scale:    mathutils.NewVec3(1+0.2*mathutils.FBM3(float64(t.UnixNano())/1e9, 0.1, 0.2, 2, 2.0, 0.5), 1, 1),
    }
    m := tr.Matrix()
    // convert Mat4 to []float64
    mat := make([]float64, 16)
    for i := 0; i < 16; i++ {
      mat[i] = m[i]
    }

    // Send tick and transform
    msgTick := map[string]any{"type": "tick", "angle": angle}
    if err := conn.WriteJSON(msgTick); err != nil {
      log.Println("write json error:", err)
      return
    }

    msg := map[string]any{"type": "transform", "id": "triangle", "matrix": mat}
    if err := conn.WriteJSON(msg); err != nil {
      log.Println("write json error:", err)
      return
    }

    // read any incoming message (non-blocking with deadline)
    conn.SetReadDeadline(time.Now().Add(1 * time.Millisecond))
    _, b, _ := conn.ReadMessage()
    if len(b) > 0 {
      var req map[string]any
      if err := json.Unmarshal(b, &req); err == nil {
        if req["type"] == "cipher" {
          data, _ := req["data"].(string)
          if serverKey != "" {
            enc, err := cipher.EncryptBase64(serverKey, []byte(data))
            if err != nil {
              log.Println("cipher encrypt error:", err)
            } else {
              resp := map[string]any{"type": "cipher_resp", "data": enc}
              conn.WriteJSON(resp)
            }
          }
        }
        if req["type"] == "request_heightmap" {
          // generate a small heightmap using FBM and send
          w, h := 32, 32
          heights := make([]float64, w*h)
          for y := 0; y < h; y++ {
            for x := 0; x < w; x++ {
              fx := float64(x)/float64(w) * 8
              fy := float64(y)/float64(h) * 8
              heights[y*w+x] = mathutils.FBM3(fx, fy, float64(t.UnixNano())/1e9, 4, 2.0, 0.5)
            }
          }
          hm := map[string]any{"type": "heightmap", "w": w, "h": h, "data": heights}
          conn.WriteJSON(hm)
        }
      }
    }
  }
}

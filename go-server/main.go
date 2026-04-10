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

func wsHandler(w http.ResponseWriter, r *http.Request) {
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
    msg := map[string]any{"type": "tick", "angle": angle}
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
          enc := cipher.XOR([]byte(data), 0x7F)
          resp := map[string]any{"type": "cipher_resp", "data": enc}
          conn.WriteJSON(resp)
        }
      }
    }
  }
}

func main() {
  http.HandleFunc("/ws", wsHandler)
  log.Println("Go server listening on :8080 (ws)")
  log.Fatal(http.ListenAndServe(":8080", nil))
}

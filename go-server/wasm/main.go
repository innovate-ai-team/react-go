//go:build js
// +build js

package main

import (
  "syscall/js"
  "react-go-server/mathutils"
)

// fbm3(x,y,z,octaves,lacunarity,gain) -> float64
func fbm(this js.Value, args []js.Value) interface{} {
  if len(args) < 6 {
    return js.ValueOf(0)
  }
  x := args[0].Float()
  y := args[1].Float()
  z := args[2].Float()
  oct := int(args[3].Int())
  lac := args[4].Float()
  gain := args[5].Float()
  v := mathutils.FBM3(x, y, z, oct, lac, gain)
  return js.ValueOf(v)
}

func register() {
  js.Global().Set("fbm3", js.FuncOf(fbm))
}

func main() {
  register()
  // prevent exit
  select {}
}

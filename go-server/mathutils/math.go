package mathutils

import (
  "math"
  "time"
)

// AngleFromTime returns a smoothly varying angle in radians based on time
func AngleFromTime(t time.Time) float64 {
  s := float64(t.UnixNano())/1e9
  return math.Mod(s*0.8, math.Pi*2)
}

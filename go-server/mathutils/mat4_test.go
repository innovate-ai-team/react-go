package mathutils

import (
  "math"
  "testing"
)

func TestMat4TranslateScaleRotate(t *testing.T) {
  v := NewVec3(1, 2, 3)
  tr := Transform{Position: NewVec3(5, 0, 0), Rotation: NewVec3(0, math.Pi/2, 0), Scale: NewVec3(2, 2, 2)}
  out := tr.Apply(v)
  if out.Len() == 0 {
    t.Fatalf("transform produced zero: %v", out)
  }
}

func TestNoiseRange(t *testing.T) {
  for _, p := range []struct{ x, y, z float64 }{{0, 0, 0}, {1.3, 4.2, -0.7}, {10.1, -3.2, 5.7}} {
    n := Noise3(p.x, p.y, p.z)
    if n < -1.0001 || n > 1.0001 {
      t.Fatalf("noise out of range: %v", n)
    }
  }
}

func TestFBMVariation(t *testing.T) {
  a := FBM3(0.1, 0.2, 0.3, 4, 2.0, 0.5)
  b := FBM3(0.11, 0.2, 0.3, 4, 2.0, 0.5)
  if math.Abs(a-b) > 0.5 {
    t.Fatalf("fbm too discontinuous: %v vs %v", a, b)
  }
}

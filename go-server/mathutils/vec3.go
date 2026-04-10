package mathutils

import "math"

// Vec3 simple 3D vector
type Vec3 struct {
  X, Y, Z float64
}

func NewVec3(x, y, z float64) Vec3 { return Vec3{X: x, Y: y, Z: z} }

func (v Vec3) Add(o Vec3) Vec3   { return Vec3{v.X + o.X, v.Y + o.Y, v.Z + o.Z} }
func (v Vec3) Sub(o Vec3) Vec3   { return Vec3{v.X - o.X, v.Y - o.Y, v.Z - o.Z} }
func (v Vec3) Scale(s float64) Vec3 { return Vec3{v.X * s, v.Y * s, v.Z * s} }
func (v Vec3) Dot(o Vec3) float64 { return v.X*o.X + v.Y*o.Y + v.Z*o.Z }
func (v Vec3) Cross(o Vec3) Vec3 {
  return Vec3{
    v.Y*o.Z - v.Z*o.Y,
    v.Z*o.X - v.X*o.Z,
    v.X*o.Y - v.Y*o.X,
  }
}
func (v Vec3) Len() float64 { return math.Sqrt(v.Dot(v)) }
func (v Vec3) Normalize() Vec3 {
  l := v.Len()
  if l == 0 {
    return v
  }
  return v.Scale(1 / l)
}

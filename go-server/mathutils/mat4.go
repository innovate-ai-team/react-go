package mathutils

// Mat4 4x4 matrix in row-major order
type Mat4 [16]float64

func IdentityMat4() Mat4 {
  return Mat4{
    1, 0, 0, 0,
    0, 1, 0, 0,
    0, 0, 1, 0,
    0, 0, 0, 1,
  }
}

func MulMat4(a, b Mat4) Mat4 {
  var r Mat4
  for i := 0; i < 4; i++ {
    for j := 0; j < 4; j++ {
      var sum float64
      for k := 0; k < 4; k++ {
        sum += a[i*4+k] * b[k*4+j]
      }
      r[i*4+j] = sum
    }
  }
  return r
}

func TranslateMat4(tx, ty, tz float64) Mat4 {
  m := IdentityMat4()
  m[3] = tx
  m[7] = ty
  m[11] = tz
  return m
}

func ScaleMat4(sx, sy, sz float64) Mat4 {
  return Mat4{
    sx, 0, 0, 0,
    0, sy, 0, 0,
    0, 0, sz, 0,
    0, 0, 0, 1,
  }
}

func RotateXMat4(a float64) Mat4 {
  ca, sa := cos(a), sin(a)
  return Mat4{
    1, 0, 0, 0,
    0, ca, -sa, 0,
    0, sa, ca, 0,
    0, 0, 0, 1,
  }
}

func RotateYMat4(a float64) Mat4 {
  ca, sa := cos(a), sin(a)
  return Mat4{
    ca, 0, sa, 0,
    0, 1, 0, 0,
    -sa, 0, ca, 0,
    0, 0, 0, 1,
  }
}

func RotateZMat4(a float64) Mat4 {
  ca, sa := cos(a), sin(a)
  return Mat4{
    ca, -sa, 0, 0,
    sa, ca, 0, 0,
    0, 0, 1, 0,
    0, 0, 0, 1,
  }
}

// ApplyMat4 applies a transformation to a Vec3 using homogeneous coordinate w=1
func ApplyMat4(m Mat4, v Vec3) Vec3 {
  x := m[0]*v.X + m[1]*v.Y + m[2]*v.Z + m[3]*1
  y := m[4]*v.X + m[5]*v.Y + m[6]*v.Z + m[7]*1
  z := m[8]*v.X + m[9]*v.Y + m[10]*v.Z + m[11]*1
  w := m[12]*v.X + m[13]*v.Y + m[14]*v.Z + m[15]*1
  if w != 0 && w != 1 {
    return Vec3{x / w, y / w, z / w}
  }
  return Vec3{x, y, z}
}

// small wrappers to avoid importing math in this file repeatedly
func cos(a float64) float64 { return __cos(a) }
func sin(a float64) float64 { return __sin(a) }

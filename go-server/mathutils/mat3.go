package mathutils

// Mat3 represents a 3x3 matrix in row-major order, typically for 2D transforms (affine)
type Mat3 [9]float64

func IdentityMat3() Mat3 {
  return Mat3{
    1, 0, 0,
    0, 1, 0,
    0, 0, 1,
  }
}

func MulMat3(a, b Mat3) Mat3 {
  var r Mat3
  for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
      var sum float64
      for k := 0; k < 3; k++ {
        sum += a[i*3+k] * b[k*3+j]
      }
      r[i*3+j] = sum
    }
  }
  return r
}

// ApplyMat3 applies a Mat3 to a Vec3 (homogeneous coordinate)
func ApplyMat3(m Mat3, v Vec3) Vec3 {
  x := m[0]*v.X + m[1]*v.Y + m[2]*v.Z
  y := m[3]*v.X + m[4]*v.Y + m[5]*v.Z
  z := m[6]*v.X + m[7]*v.Y + m[8]*v.Z
  return Vec3{x, y, z}
}

package mathutils

// Transform holds basic TRS (translate, rotate euler, scale)
type Transform struct {
  Position Vec3
  Rotation Vec3 // Euler angles in radians: X (pitch), Y (yaw), Z (roll)
  Scale    Vec3
}

func NewTransform() Transform {
  return Transform{Position: NewVec3(0, 0, 0), Rotation: NewVec3(0, 0, 0), Scale: NewVec3(1, 1, 1)}
}

// Matrix builds a Mat4 from TRS (T * Rz * Ry * Rx * S)
func (t Transform) Matrix() Mat4 {
  T := TranslateMat4(t.Position.X, t.Position.Y, t.Position.Z)
  Rx := RotateXMat4(t.Rotation.X)
  Ry := RotateYMat4(t.Rotation.Y)
  Rz := RotateZMat4(t.Rotation.Z)
  S := ScaleMat4(t.Scale.X, t.Scale.Y, t.Scale.Z)

  // Order: T * Rz * Ry * Rx * S
  m := MulMat4(T, MulMat4(Rz, MulMat4(Ry, MulMat4(Rx, S))))
  return m
}

func (t Transform) Apply(v Vec3) Vec3 {
  return ApplyMat4(t.Matrix(), v)
}

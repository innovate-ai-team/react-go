package mathutils

import "testing"

func TestVec3Basics(t *testing.T) {
  a := NewVec3(1, 2, 3)
  b := NewVec3(4, -1, 0)

  if got := a.Add(b); got != NewVec3(5, 1, 3) {
    t.Fatalf("Add failed: got=%v", got)
  }

  if got := a.Sub(b); got != NewVec3(-3, 3, 3) {
    t.Fatalf("Sub failed: got=%v", got)
  }

  if d := a.Dot(b); d != 1*4+2*(-1)+3*0 {
    t.Fatalf("Dot failed: %v", d)
  }

  c := a.Cross(b)
  if c.Len() == 0 {
    t.Fatalf("Cross produced zero vector: %v", c)
  }
}

func TestMat3Apply(t *testing.T) {
  m := IdentityMat3()
  v := NewVec3(1, 2, 1)
  out := ApplyMat3(m, v)
  if out != v {
    t.Fatalf("Identity apply failed: %v", out)
  }
}

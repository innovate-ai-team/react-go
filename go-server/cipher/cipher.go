package cipher

// XOR is a toy cipher: XOR every byte with key
func XOR(data []byte, key byte) string {
  out := make([]byte, len(data))
  for i, b := range data {
    out[i] = b ^ key
  }
  return string(out)
}

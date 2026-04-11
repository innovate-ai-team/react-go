package cipher

import (
  "crypto/aes"
  "crypto/cipher"
  "crypto/rand"
  "encoding/base64"
  "fmt"
  "io"
)

// NewKey generates a 32-byte (256-bit) random key and returns it base64-encoded.
func NewKey() (string, error) {
  k := make([]byte, 32)
  if _, err := rand.Read(k); err != nil {
    return "", err
  }
  return base64.StdEncoding.EncodeToString(k), nil
}

// EncryptBase64 encrypts plaintext using AES-GCM with a base64-encoded key and
// returns a base64-encoded nonce+ciphertext.
func EncryptBase64(keyB64 string, plaintext []byte) (string, error) {
  key, err := base64.StdEncoding.DecodeString(keyB64)
  if err != nil {
    return "", err
  }
  block, err := aes.NewCipher(key)
  if err != nil {
    return "", err
  }
  gcm, err := cipher.NewGCM(block)
  if err != nil {
    return "", err
  }
  nonce := make([]byte, gcm.NonceSize())
  if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
    return "", err
  }
  ct := gcm.Seal(nonce, nonce, plaintext, nil)
  return base64.StdEncoding.EncodeToString(ct), nil
}

// DecryptBase64 decrypts a base64 nonce+ciphertext using AES-GCM and returns plaintext.
func DecryptBase64(keyB64, cipherB64 string) ([]byte, error) {
  key, err := base64.StdEncoding.DecodeString(keyB64)
  if err != nil {
    return nil, err
  }
  block, err := aes.NewCipher(key)
  if err != nil {
    return nil, err
  }
  gcm, err := cipher.NewGCM(block)
  if err != nil {
    return nil, err
  }
  ct, err := base64.StdEncoding.DecodeString(cipherB64)
  if err != nil {
    return nil, err
  }
  ns := gcm.NonceSize()
  if len(ct) < ns {
    return nil, fmt.Errorf("ciphertext too short")
  }
  nonce, ciphertext := ct[:ns], ct[ns:]
  return gcm.Open(nil, nonce, ciphertext, nil)
}

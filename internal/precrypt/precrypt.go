package precrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func Encrypt(text []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	nc := gcm.Seal(nonce, nonce, text, nil)
	b64 := make([]byte, base64.RawStdEncoding.EncodedLen(len(nc)))
	base64.RawStdEncoding.Encode(b64, nc)
	return b64, nil
}

func Decrypt(b64 []byte, key []byte) ([]byte, error) {
	nc := make([]byte, base64.RawStdEncoding.DecodedLen(len(b64)))
	if _, err := base64.RawStdEncoding.Decode(nc, b64); err != nil {
		return nil, err
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}
	nonceSize := gcm.NonceSize()
	nonce, cipherText := nc[:nonceSize], nc[nonceSize:]
	return gcm.Open(nil, nonce, cipherText, nil)
}

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
)

// GCM is the Galois/Counter encryption mode
// GCM is an implementation of crypto.Algorithm
type GCM struct {
	// key represents the key used for encryption and decryption
	key []byte
}

// NewGCM creates and returns a GCM implementation of the crypto Algorithm
func NewGCM(key []byte) Algorithm {
	return &GCM{key: key}
}

// Encrypt returns the encrypted (cipher) value from the input plain value
func (c *GCM) Encrypt(plainValue []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	cipherValue := gcm.Seal(nil, nonce, plainValue, nil)

	return cipherValue, nil
}

// Decrypt returns the plain value from the input cipher value
func (c *GCM) Decrypt(cipherValue []byte) ([]byte, error) {
	block, err := aes.NewCipher(c.key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	plainValue, err := gcm.Open(nil, nonce, cipherValue, nil)
	if err != nil {
		return nil, err
	}

	return plainValue, nil
}

package crypto

// Algorithm represents the symmetrical encryption algorithm.
// By providing a given key, the algorithm and encrypt or decrypt the input value.
type Algorithm interface {
	Encrypt(plainValue []byte) ([]byte, error)
	Decrypt(cipherValue []byte) ([]byte, error)
}

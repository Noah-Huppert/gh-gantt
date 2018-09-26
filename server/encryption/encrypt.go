package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

// AESEncrypt text using aes
func AESEncrypt(key []byte, plainText []byte) ([]byte, error) {
	// Create aes cipher
	aesCipher, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("error creating aes cipher: %s", err.Error())
	}

	// Create GCM aes cipher
	gcmCipher, err := cipher.NewGCM(aesCipher)
	if err != nil {
		return nil, fmt.Errorf("error creating aes GCM cipher: %s", err.Error())
	}

	// Generate nonce
	nonce := make([]byte, gcmCipher.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, fmt.Errorf("error generating nonce: %s", err.Error())
	}

	// Encrypt
	cipherText := gcmCipher.Seal(nil, nonce, plainText, nil)

	return cipherText, nil
}

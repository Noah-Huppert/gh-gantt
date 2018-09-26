package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"io"
)

// AESDecrypt text using aes
func AESDecrypt(key []byte, cipherText []byte) ([]byte, error) {
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

	// Check cipherText is larger than nonce
	if len(cipherText) < gcmCipher.NonceSize() {
		return nil, errors.New("cipherText must be longer than nonce")
	}

	// Generate nonce
	nonce := make([]byte, gcmCipher.NonceSize())

	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, fmt.Errorf("error generating nonce: %s", err.Error())
	}

	// Decrypt
	plainText, err := gcmCipher.Open(nil, nonce, cipherText, nil)

	return plainText, nil
}

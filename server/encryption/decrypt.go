package encryption

import (
	"crypto/aes"
	"crypto/cipher"
	"errors"
	"fmt"
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
	nonceSize := gcmCipher.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText must be longer than nonce")
	}

	// Decrypt
	plainText, err := gcmCipher.Open(nil, cipherText[0:nonceSize], cipherText[nonceSize:], nil)
	if err != nil {
		return nil, fmt.Errorf("error decrypting: %s", err.Error())
	}

	return plainText, nil
}

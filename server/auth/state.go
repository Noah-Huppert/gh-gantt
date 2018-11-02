package auth

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/crypto/ed25519"
)

// RawStateBytesLen is the number of bytes a raw state bytes array should contain
const RawStateBytesLen int = 32

// MakeState return a series of signed bytes which are base64 encoded. Used as the GH auth state.
// The returned value is in the format: `<state>.<state signature>`
func MakeState(stateSigningPrivKey []byte) string {
	// Get raw bytes to sign
	stateBytes := make([]byte, RawStateBytesLen)
	rand.Read(stateBytes)

	// Sign raw bytes
	stateSignature := ed25519.Sign(stateSigningPrivKey, stateBytes)

	// Format
	formattedState := fmt.Sprintf("%s.%s", stateBytes, stateSignature)

	// Base64 encode
	return base64.StdEncoding.EncodeToString([]byte(formattedState))
}

// VerifyState checks to see if a series of bytes was signed by the GH auth state signing key
// The provided state is expected to be in the format: `<state>.<signed state>`
func VerifyState(stateSigningPubKey []byte, state string) (bool, error) {
	// Base64 Decode
	b64DecodedState, err := base64.StdEncoding.DecodeString(state)
	if err != nil {
		return false, fmt.Errorf("error base64 decoding state: %s", err.Error())
	}

	// Separate state parts
	parts := strings.Split(string(b64DecodedState), ".")

	if len(parts) != 2 {
		return false, errors.New("state not in format: <state>.<state signature>")
	}

	stateBytes := []byte(parts[0])
	stateSignature := []byte(parts[1])

	return ed25519.Verify([]byte(stateSigningPubKey), stateBytes, stateSignature), nil
}
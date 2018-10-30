package auth

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thanhpk/randstr"
	"golang.org/x/crypto/ed25519"
)

// RawStateBytesLen is the number of bytes a raw state bytes array should contain
const RawStateBytesLen int = 32

// MakeStateBytes return a series of signed bytes used as the GH auth state
// The returned value is in the format: `<state>.<state signature>`
func MakeStateBytes(stateSigningPrivKey []byte) []byte {
	// Get raw bytes to sign
	rawState := randstr.Byte(RawStateBytesLen)

	// Sign raw bytes
	stateSignature := ed25519.Sign(stateSigningPrivKey, rawState)

	// Format
	formattedState := fmt.Sprintf("%s.%s", rawState, stateSignature)

	return []byte(formattedState)
}

// VerifyStateBytes checks to see if a series of bytes was signed by the GH auth state signing key
// The provided state is expected to be in the format: `<state>.<signed state>`
func VerifyStateBytes(stateSigningPubKey, state []byte) (bool, error) {
	// Separate state parts
	parts := strings.Split(string(state), ".")

	if len(parts) != 2 {
		return false, errors.New("state not in format: <state>.<state signature>")
	}

	state = []byte(parts[0])
	stateSignature = []byte(parts[1])

	return ed25519.Verify(stateSigningPubKey, state, stateSignature)
}

package keypairs

import (
	"crypto/sha512"
	"encoding/hex"
	"io"
	"strings"
)

type randomizer struct {
	io.Reader
}

func (r *randomizer) generateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := r.Read(b) //nolint
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Returns the first 32 bytes of a sha512 hash of a message
func sha512Half(msg []byte) []byte {
	h := sha512.Sum512(msg)
	return h[:32]
}

func formatKey(k []byte) string {
	return strings.ToUpper(hex.EncodeToString(k))
}

func deformatKey(k string) []byte {
	b, _ := hex.DecodeString(k)
	return b
}

package keypairs

import (
	"crypto/sha512"
	"io"
)

type randomizer struct {
	io.Reader
}

func (r *randomizer) generateBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := r.Read(b)
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

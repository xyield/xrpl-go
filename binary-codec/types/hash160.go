package types

import (
	"encoding/hex"
	"strings"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
)

var _ hashI = (*Hash128)(nil)

// Hash160 struct represents a 160-bit hash.
type Hash160 struct {
}

// NewHash160 is a constructor for creating a new 160-bit hash.
func NewHash160() *Hash160 {
	return &Hash160{}
}

// getLength method for hash returns the hash's length.
func (h *Hash160) getLength() int {
	return 20
}

// FromJson method for hash converts a hexadecimal string from JSON to a byte array.
// It returns an error if the conversion fails or the length of the decoded byte array is not as expected.
func (h *Hash160) FromJson(json any) ([]byte, error) {
	if _, ok := json.(string); !ok {
		return nil, ErrInvalidHashType
	}
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	if h.getLength() != len(v) {
		return nil, &ErrInvalidHashLength{Expected: h.getLength()}
	}
	return v, nil
}

// ToJson method for hash reads a certain number of bytes from a BinaryParser and converts it into a hexadecimal string.
// It returns an error if the read operation fails.
func (h *Hash160) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(h.getLength())
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

package types

import (
	"encoding/hex"
	"strings"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

var _ hashI = (*Hash128)(nil)

// Hash128 struct represents a 128-bit hash.
type Hash128 struct {
}

// NewHash128 is a constructor for creating a new 128-bit hash.
func NewHash128() *Hash128 {
	return &Hash128{}
}

// getLength method for hash returns the hash's length.
func (h *Hash128) getLength() int {
	return 16
}

// FromJson method for hash converts a hexadecimal string from JSON to a byte array.
// It returns an error if the conversion fails or the length of the decoded byte array is not as expected.
func (h *Hash128) FromJson(json any) ([]byte, error) {
	var s string
	switch json := json.(type) {
	case string:
		s = json
	case types.Hash128:
		s = string(json)
	default:
		return nil, ErrInvalidHashType
	}
	v, err := hex.DecodeString(s)
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
func (h *Hash128) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(h.getLength())
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

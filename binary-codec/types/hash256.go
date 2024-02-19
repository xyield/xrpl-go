package types

import (
	"encoding/hex"
	"strings"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

var _ hashI = (*Hash256)(nil)

// Hash256 struct represents a 256-bit hash.
type Hash256 struct {
}

// NewHash256 is a constructor for creating a new 256-bit hash.
func NewHash256() *Hash256 {
	return &Hash256{}
}

// getLength method for hash returns the hash's length.
func (h *Hash256) getLength() int {
	return 32
}

func (h *Hash256) FromJson(json any) ([]byte, error) {
	var s string
	switch json := json.(type) {
	case string:
		s = json
	case types.Hash256:
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
func (h *Hash256) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(h.getLength())
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

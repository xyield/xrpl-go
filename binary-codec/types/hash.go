package types

import (
	"encoding/hex"
	"fmt"
)

type ErrInvalidHashLength struct {
	Expected int
}

func (e *ErrInvalidHashLength) Error() string {
	return fmt.Sprintf("invalid hash length expected length %v", e.Expected)
}

type hashI interface {
	getLength()
}

type hash struct {
	Length int
}

func newHash(l int) hash {
	return hash{
		Length: l,
	}
}

func (h *hash) getLength() int {
	return h.Length
}

func (h *hash) SerializeJson(json any) ([]byte, error) {
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	if h.getLength() != len(v) {
		return nil, &ErrInvalidHashLength{Expected: h.getLength()}
	}
	return v, nil
}

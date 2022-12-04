package types

import (
	"encoding/hex"
	"fmt"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type ErrInvalidHashLength struct {
	Expected int
}

func (e *ErrInvalidHashLength) Error() string {
	return fmt.Sprintf("invalid hash length expected length %v", e.Expected)
}

type hashI interface {
	SerializedType
	getLength() int
}

type hash struct {
	Length int
}

func newHash(l int) hash {
	return hash{
		Length: l,
	}
}

func (h hash) getLength() int {
	return h.Length
}

func (h hash) FromJson(json any) ([]byte, error) {
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	if h.getLength() != len(v) {
		return nil, &ErrInvalidHashLength{Expected: h.getLength()}
	}
	return v, nil
}

func (h hash) FromParser(p *serdes.BinaryParser) ([]byte, error) {
	return nil, nil
}

package types

import (
	"encoding/hex"
	"errors"
	"strings"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

var (
	ErrNoLengthPrefix error = errors.New("no length prefix size given")
)

type Blob struct{}

func (b *Blob) FromJson(json any) ([]byte, error) {
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (b *Blob) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	if opts == nil {
		return nil, ErrNoLengthPrefix
	}
	val, err := p.ReadBytes(opts[0])
	if err != nil {
		return nil, err
	}
	return strings.ToUpper(hex.EncodeToString(val)), nil
}

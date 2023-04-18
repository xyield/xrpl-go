package types

import (
	"encoding/hex"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type Blob struct{}

func (b *Blob) FromJson(json any) ([]byte, error) {
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (b *Blob) FromParser(p *serdes.BinaryParser) (any, error) {
	return nil, nil
}

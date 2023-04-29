package types

import (
	"bytes"
	"encoding/binary"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type UInt32 struct{}

// Serializes the given json value to a 32-bit UInt byte slice.
func (u *UInt32) FromJson(value any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(value.(int)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (u *UInt32) FromParser(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(4)
	if err != nil {
		return nil, err
	}
	return int(binary.BigEndian.Uint32(b)), nil
}

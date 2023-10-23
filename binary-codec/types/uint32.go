package types

import (
	"bytes"
	"encoding/binary"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

// UInt32 represents a 32-bit unsigned integer.
type UInt32 struct{}

// FromJson converts a JSON value into a serialized byte slice representing a 32-bit unsigned integer.
// The input value is assumed to be an integer. If the serialization fails, an error is returned.
func (u *UInt32) FromJson(value any) ([]byte, error) {
	v := expandInt(value)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, v)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON integer value. This method assumes the parser contains data representing
// a 32-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt32) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(4)
	if err != nil {
		return nil, err
	}
	return int(binary.BigEndian.Uint32(b)), nil
}

func expandInt(v any) uint32 {
	switch v := v.(type) {
	case types.FlagsI:
		return v.ToUint()
	case uint:
		return uint32(v)
	default:
		return v.(uint32)
	}
}

package types

import (
	"bytes"
	"encoding/binary"

	"github.com/CreatureDev/xrpl-go/binary-codec/definitions"
	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
)

// UInt8 represents an 8-bit unsigned integer.
type UInt8 struct{}

// FromJson converts a JSON value into a serialized byte slice representing an 8-bit unsigned integer.
// If the input value is a string, it's assumed to be a transaction result name, and the method will
// attempt to convert it into a transaction result type code. If the conversion fails, an error is returned.
func (u *UInt8) FromJson(value any) ([]byte, error) {
	var u8 uint8

	switch v := value.(type) {
	case string:
		tc, err := definitions.Get().GetTransactionResultTypeCodeByTransactionResultName(v)
		if err != nil {
			return nil, err
		}
		u8 = uint8(tc)
	case uint8:
		u8 = v
	case int:
		u8 = uint8(v)
	case int32:
		u8 = uint8(v)
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, u8)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON integer value. This method assumes the parser contains data representing
// an 8-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt8) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(1)
	if err != nil {
		return nil, err
	}
	return int(b[0]), nil
}

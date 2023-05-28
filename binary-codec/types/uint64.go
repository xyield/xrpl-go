package types

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

// UInt64 represents a 64-bit unsigned integer.
type UInt64 struct{}

// FromJson converts a JSON value into a serialized byte slice representing a 64-bit unsigned integer.
// The input value is assumed to be a string representation of an integer. If the serialization fails, an error is returned.
func (u *UInt64) FromJson(value any) ([]byte, error) {

	stringToUint64, err := strconv.ParseUint(value.(string), 10, 64)

	if err != nil {
		return nil, err
	} else {
		value = stringToUint64
	}

	buf := new(bytes.Buffer)
	err = binary.Write(buf, binary.BigEndian, value)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON string value. This method assumes the parser contains data representing
// a 64-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt64) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(8)
	if err != nil {
		return nil, err
	}
	return strconv.Itoa(int(binary.BigEndian.Uint64(b))), nil
}

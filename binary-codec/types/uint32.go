package types

import (
	"bytes"
	"encoding/binary"
	"reflect"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
)

// UInt32 represents a 32-bit unsigned integer.
type UInt32 struct{}

// FromJson converts a JSON value into a serialized byte slice representing a 32-bit unsigned integer.
// The input value is assumed to be an integer. If the serialization fails, an error is returned.
func (u *UInt32) FromJson(value any) ([]byte, error) {
	v := expandInt(value)
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(v.(uint)))

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

func expandInt(v any) any {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		return rv.Elem().Interface()
	}
	return v
}

package types

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"regexp"
	"strings"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
)

var UINT64_HEX_REGEX = regexp.MustCompile("^[0-9a-fA-F]{1,16}$")

// UInt64 represents a 64-bit unsigned integer.
type UInt64 struct{}

var (
	ErrInvalidUInt64String error = errors.New("invalid UInt64 string, value should be hex encoded")
	ErrInvalidUInt64Value  error = errors.New("invalid UInt64 value, value should be an uint or a hex encoded string")
)

// FromJson converts a JSON value into a serialized byte slice representing a 64-bit unsigned integer.
// The input value is assumed to be a string representation of an integer. If the serialization fails, an error is returned.
func (u *UInt64) FromJson(value any) ([]byte, error) {

	var buf = new(bytes.Buffer)

	switch v := value.(type) {
	case uint64:
		value = v
		err := binary.Write(buf, binary.BigEndian, value)
		if err != nil {
			return nil, err
		}
	case uint:
		value = uint64(v)
		err := binary.Write(buf, binary.BigEndian, value)
		if err != nil {
			return nil, err
		}
	case string:
		if !UINT64_HEX_REGEX.MatchString(v) {
			return nil, ErrInvalidUInt64String
		}
		value = rjust(v, 16, "0") // right justify the string
		decoded, err := hex.DecodeString(value.(string))
		if err != nil {
			return nil, err
		}
		buf.Write(decoded)
	default:
		return nil, ErrInvalidUInt64Value
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
	return strings.ToUpper(hex.EncodeToString(b)), nil
}

func rjust(s string, n int, pad string) string {
	return strings.Repeat(pad, n-len(s)) + s
}

package types

import (
	"bytes"
	"encoding/binary"
	"strconv"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type UInt64 struct{}

// Serializes the given json value to a 64-bit UInt byte slice.
func (u *UInt64) FromJson(value any) ([]byte, error) {

	// convert string to uint64

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

func (u *UInt64) FromParser(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(8)
	if err != nil {
		return nil, err
	}
	return strconv.Itoa(int(binary.BigEndian.Uint64(b))), nil
}

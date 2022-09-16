package types

import (
	"bytes"
	"encoding/binary"
	"strconv"
)

type UInt64 struct{}

// Serializes the given json value to a 64-bit UInt byte slice.
func (u *UInt64) SerializeJson(value any) ([]byte, error) {

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

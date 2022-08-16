package types

import (
	"bytes"
	"encoding/binary"
)

type UInt64 struct{}

func (u *UInt64) SerializeJson(value any) ([]byte, error) {
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, value)

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

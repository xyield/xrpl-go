package types

import (
	"bytes"
	"encoding/binary"
)

type UInt8 struct{}

func (u *UInt8) SerializeJson(value any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint8(value.(int)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

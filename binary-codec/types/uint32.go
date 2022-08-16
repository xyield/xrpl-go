package types

import (
	"bytes"
	"encoding/binary"
)

type UInt32 struct{}

func (u *UInt32) SerializeJson(value any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint32(value.(int)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

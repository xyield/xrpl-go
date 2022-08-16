package types

import (
	"bytes"
	"encoding/binary"
)

type UInt16 struct{}

func (u *UInt16) SerializeJson(value any) ([]byte, error) {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint16(7))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

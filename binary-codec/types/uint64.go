package types

import (
	"bytes"
	"encoding/binary"
)

type UInt64 struct{}

func (u *UInt64) SerializeJson(value any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint64(value.(int)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

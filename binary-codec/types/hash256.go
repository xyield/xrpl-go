package types

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"encoding/hex"
	"fmt"
)

type Hash256 struct{}

func (h *Hash256) SerializeJson(value any) ([]byte, error) {
	buf := new(bytes.Buffer)

	value, _ = hex.DecodeString(value.(string))
	err := binary.Write(buf, binary.BigEndian, value)

	if err != nil {
		return nil, err
	}

	ch := CalcHash(CalcHash(buf.Bytes(), sha256.New()), sha256.New())
	fmt.Println(hex.EncodeToString(ch))
	return ch, nil
}

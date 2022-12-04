package types

import (
	"bytes"
	"encoding/binary"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type UInt8 struct{}

// Serializes the given json value to an 8-bit UInt byte slice.
func (u *UInt8) FromJson(value any) ([]byte, error) {

	if _, ok := value.(string); ok {
		tc, err := definitions.Get().GetTransactionResultTypeCodeByTransactionResultName(value.(string))
		if err != nil {
			return nil, err
		}
		value = tc
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint8(value.(int)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (u *UInt8) FromParser(p *serdes.BinaryParser) ([]byte, error) {
	return nil, nil
}

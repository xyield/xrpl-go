package types

import (
	"bytes"
	"encoding/binary"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type UInt16 struct{}

// Serializes the given json value to a 16-bit UInt byte slice.
func (u *UInt16) FromJson(value any) ([]byte, error) {

	if _, ok := value.(string); ok {
		tc, err := definitions.Get().GetTransactionTypeCodeByTransactionTypeName(value.(string))
		if err != nil {
			tc, err = definitions.Get().GetLedgerEntryTypeCodeByLedgerEntryTypeName(value.(string))
			if err != nil {
				return nil, err
			}
		}
		value = tc
	}

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, uint16(value.(int32)))

	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (u *UInt16) FromParser(p *serdes.BinaryParser) ([]byte, error) {
	return nil, nil
}

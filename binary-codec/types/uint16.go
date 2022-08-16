package types

import (
	"bytes"
	"encoding/binary"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

type UInt16 struct{}

func (u *UInt16) SerializeJson(value any) ([]byte, error) {

	if _, ok := value.(string); ok {
		tc, err := definitions.Get().GetTransactionTypeCodeByTransactionTypeName(value.(string))
		if err != nil {
			tc, err = definitions.Get().GetLedgerEntryTypeCodeByLedgerEntryTypeName(value.(string))
			if err != nil {
				return nil, err
			} else {
				value = tc
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

package types

import (
	"bytes"
	"encoding/binary"

	"github.com/CreatureDev/xrpl-go/binary-codec/definitions"
	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
)

// UInt16 represents a 16-bit unsigned integer.
type UInt16 struct{}

// FromJson converts a JSON value into a serialized byte slice representing a 16-bit unsigned integer.
// If the input value is a string, it's assumed to be a transaction type or ledger entry type name, and the
// method will attempt to convert it into a corresponding type code. If the conversion fails, an error is returned.
func (u *UInt16) FromJson(value any) ([]byte, error) {
	switch v := value.(type) {
	case uint16:
		value = v
	case uint:
		value = uint16(v)
	case int:
		value = uint16(v)
	case transactions.TxType:
		tc, err := definitions.Get().GetTransactionTypeCodeByTransactionTypeName(string(v))
		if err != nil {
			return nil, err
		}
		value = uint16(tc)
	case ledger.LedgerEntryType:
		tc, err := definitions.Get().GetLedgerEntryTypeCodeByLedgerEntryTypeName(string(v))
		if err != nil {
			return nil, err
		}
		value = uint16(tc)
	}
	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into a JSON integer value. This method assumes the parser contains data representing
// a 16-bit unsigned integer. If the parsing fails, an error is returned.
func (u *UInt16) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {
	b, err := p.ReadBytes(2)
	if err != nil {
		return nil, err
	}
	return int(binary.BigEndian.Uint16(b)), nil
}

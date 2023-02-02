package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type CheckCreate struct {
	BaseTx
	Destination    Address
	SendMax        CurrencyAmount
	DestinationTag uint    `json:",omitempty"`
	Expiration     uint    `json:",omitempty"`
	InvoiceID      Hash256 `json:",omitempty"`
}

func (*CheckCreate) TxType() TxType {
	return CheckCreateTx
}

func UnmarshalCheckCreateTx(data json.RawMessage) (Tx, error) {
	var ret CheckCreate
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

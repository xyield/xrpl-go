package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Payment struct {
	BaseTx
	Amount         CurrencyAmount
	Destination    Address
	DestinationTag uint           `json:",omitempty"`
	InvoiceID      uint           `json:",omitempty"`
	Paths          [][]PathStep   `json:",omitempty"`
	SendMax        CurrencyAmount `json:",omitempty"`
	DeliverMin     CurrencyAmount `json:",omitempty"`
}

func (*Payment) TxType() TxType {
	return PaymentTx
}

func UnmarshalPaymentTx(data json.RawMessage) (Tx, error) {
	var ret Payment
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

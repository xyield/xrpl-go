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
	type pHelper struct {
		BaseTx
		Amount         json.RawMessage
		Destination    Address
		DestinationTag uint            `json:",omitempty"`
		InvoiceID      uint            `json:",omitempty"`
		Paths          [][]PathStep    `json:",omitempty"`
		SendMax        json.RawMessage `json:",omitempty"`
		DeliverMin     json.RawMessage `json:",omitempty"`
	}
	var h pHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	ret = Payment{
		BaseTx:         h.BaseTx,
		Destination:    h.Destination,
		DestinationTag: h.DestinationTag,
		InvoiceID:      h.InvoiceID,
		Paths:          h.Paths,
	}
	var amount, max, min CurrencyAmount
	var err error
	amount, err = UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return nil, err
	}
	max, err = UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return nil, err
	}
	min, err = UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return nil, err
	}
	ret.Amount = amount
	ret.DeliverMin = min
	ret.SendMax = max

	return &ret, nil
}

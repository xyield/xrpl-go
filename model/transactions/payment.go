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

func (p *Payment) UnmarshalJSON(data []byte) error {
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
		return err
	}
	*p = Payment{
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
		return err
	}
	max, err = UnmarshalCurrencyAmount(h.SendMax)
	if err != nil {
		return err
	}
	min, err = UnmarshalCurrencyAmount(h.DeliverMin)
	if err != nil {
		return err
	}
	p.Amount = amount
	p.DeliverMin = min
	p.SendMax = max

	return nil
}

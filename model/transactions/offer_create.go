package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type OfferCreate struct {
	BaseTx
	Expiration    uint `json:",omitempty"`
	OfferSequence uint `json:",omitempty"`
	TakerGets     CurrencyAmount
	TakerPays     CurrencyAmount
}

func (*OfferCreate) TxType() TxType {
	return OfferCreateTx
}

func (o *OfferCreate) UnmarshalJSON(data []byte) error {
	type ocHelper struct {
		BaseTx
		Expiration    uint `json:",omitempty"`
		OfferSequence uint `json:",omitempty"`
		TakerGets     json.RawMessage
		TakerPays     json.RawMessage
	}
	var h ocHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*o = OfferCreate{
		BaseTx:        h.BaseTx,
		Expiration:    h.Expiration,
		OfferSequence: h.OfferSequence,
	}

	var gets, pays CurrencyAmount
	var err error
	gets, err = UnmarshalCurrencyAmount(h.TakerGets)
	if err != nil {
		return err
	}
	pays, err = UnmarshalCurrencyAmount(h.TakerPays)
	if err != nil {
		return err
	}
	o.TakerGets = gets
	o.TakerPays = pays

	return nil
}

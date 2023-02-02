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

func UnmarshalOfferCreateTx(data json.RawMessage) (Tx, error) {
	var ret OfferCreate
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

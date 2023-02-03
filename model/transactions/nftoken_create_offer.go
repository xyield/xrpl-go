package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenCreateOffer struct {
	BaseTx
	Owner       Address `json:",omitempty"`
	NFTokenID   NFTokenID
	Amount      CurrencyAmount
	Expiration  uint    `json:",omitempty"`
	Destination Address `json:",omitempty"`
}

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}

func UnmarshalNFTokenCreateOfferTx(data json.RawMessage) (Tx, error) {
	var ret NFTokenCreateOffer
	type ncoHelper struct {
		BaseTx
		Owner       Address `json:",omitempty"`
		NFTokenID   NFTokenID
		Amount      json.RawMessage
		Expiration  uint    `json:",omitempty"`
		Destination Address `json:",omitempty"`
	}
	var h ncoHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	ret = NFTokenCreateOffer{
		BaseTx:      h.BaseTx,
		Owner:       h.Owner,
		NFTokenID:   h.NFTokenID,
		Expiration:  h.Expiration,
		Destination: h.Destination,
	}

	amount, err := UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return nil, err
	}
	ret.Amount = amount
	return &ret, nil
}

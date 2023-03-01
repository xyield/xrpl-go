package transactions

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenCreateOffer struct {
	BaseTx
	Owner       types.Address `json:",omitempty"`
	NFTokenID   types.NFTokenID
	Amount      types.CurrencyAmount
	Expiration  uint          `json:",omitempty"`
	Destination types.Address `json:",omitempty"`
}

func (*NFTokenCreateOffer) TxType() TxType {
	return NFTokenCreateOfferTx
}

func (n *NFTokenCreateOffer) UnmarshalJSON(data []byte) error {
	type ncoHelper struct {
		BaseTx
		Owner       types.Address `json:",omitempty"`
		NFTokenID   types.NFTokenID
		Amount      json.RawMessage
		Expiration  uint          `json:",omitempty"`
		Destination types.Address `json:",omitempty"`
	}
	var h ncoHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*n = NFTokenCreateOffer{
		BaseTx:      h.BaseTx,
		Owner:       h.Owner,
		NFTokenID:   h.NFTokenID,
		Expiration:  h.Expiration,
		Destination: h.Destination,
	}

	amount, err := types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	n.Amount = amount
	return nil
}

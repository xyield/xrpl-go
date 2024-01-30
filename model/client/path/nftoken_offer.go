package path

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type NFTokenOffer struct {
	Amount            types.CurrencyAmount `json:"amount"`
	Flags             uint                 `json:"flags"`
	NFTokenOfferIndex string               `json:"nft_offer_index"`
	Owner             types.Address        `json:"owner"`
}

func (o *NFTokenOffer) UnmarshalJSON(data []byte) error {
	type ntoHelper struct {
		Amount            json.RawMessage `json:"amount"`
		Flags             uint            `json:"flags"`
		NFTokenOfferIndex string          `json:"nft_offer_index"`
		Owner             types.Address   `json:"owner"`
	}
	var h ntoHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*o = NFTokenOffer{
		Flags:             h.Flags,
		NFTokenOfferIndex: h.NFTokenOfferIndex,
		Owner:             h.Owner,
	}
	var c types.CurrencyAmount
	c, err = types.UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	o.Amount = c

	return nil
}

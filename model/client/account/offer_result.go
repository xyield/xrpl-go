package account

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type OfferResultFlags uint32

func (f OfferResultFlags) ToUint() uint32 {
	return uint32(f)
}

type OfferResult struct {
	Flags      OfferResultFlags     `json:"flags"`
	Sequence   uint32               `json:"seq"`
	TakerGets  types.CurrencyAmount `json:"taker_gets"`
	TakerPays  types.CurrencyAmount `json:"taker_pays"`
	Quality    string               `json:"quality"`
	Expiration uint                 `json:"expiration,omitempty"`
}

func (r *OfferResult) UnmarshalJSON(data []byte) error {
	type orHelper struct {
		Flags      OfferResultFlags `json:"flags"`
		Sequence   uint32           `json:"seq"`
		TakerGets  json.RawMessage  `json:"taker_gets"`
		TakerPays  json.RawMessage  `json:"taker_pays"`
		Quality    string           `json:"quality"`
		Expiration uint             `json:"expiration,omitempty"`
	}
	var h orHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = OfferResult{
		Flags:      h.Flags,
		Sequence:   h.Sequence,
		Quality:    h.Quality,
		Expiration: h.Expiration,
	}

	var gets, pays types.CurrencyAmount
	var err error
	gets, err = types.UnmarshalCurrencyAmount(h.TakerGets)
	if err != nil {
		return err
	}
	pays, err = types.UnmarshalCurrencyAmount(h.TakerPays)
	if err != nil {
		return err
	}

	r.TakerGets = gets
	r.TakerPays = pays
	return nil
}

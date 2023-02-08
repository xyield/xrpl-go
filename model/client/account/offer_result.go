package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type OfferResultFlags uint

type OfferResult struct {
	Flags      OfferResultFlags `json:"flags"`
	Sequence   uint             `json:"seq"`
	TakerGets  CurrencyAmount   `json:"taker_gets"`
	TakerPays  CurrencyAmount   `json:"taker_pays"`
	Quality    string           `json:"quality"`
	Expiration uint             `json:"expiration,omitempty"`
}

func (r *OfferResult) UnmarshalJSON(data []byte) error {
	type orHelper struct {
		Flags      OfferResultFlags `json:"flags"`
		Sequence   uint             `json:"seq"`
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

	r.TakerGets = gets
	r.TakerPays = pays
	return nil
}

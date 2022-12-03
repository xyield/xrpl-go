package account

import (
	. "github.com/xyield/xrpl-go/model/transactions"
)

type OfferFlags uint

type OfferResult struct {
	Flags      OfferFlags     `json:"flags"`
	Sequence   uint           `json:"seq"`
	TakerGets  CurrencyAmount `json:"taker_gets"`
	TakerPays  CurrencyAmount `json:"taker_pays"`
	Quality    string         `json:"quality"`
	Expiration uint           `json:"expiration,omitempty"`
}

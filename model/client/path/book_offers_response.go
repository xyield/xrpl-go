package path

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type BookOffersResponse struct {
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index,omitempty"`
	LedgerIndex        common.LedgerIndex `json:"ledger_index,omitempty"`
	LedgerHash         common.LedgerHash  `json:"ledger_hash,omitempty"`
	Offers             []BookOffer        `json:"offers"`
}

type BookOffer struct {
	ledger.Offer
	OwnerFunds      string               `json:"offer_funds"`
	TakerGetsFunded types.CurrencyAmount `json:"taker_gets_funded,omitempty"`
	TakerPaysFunded types.CurrencyAmount `json:"taker_pays_funded,omitempty"`
	Quality         string               `json:"quality"`
}

func (o *BookOffer) UnmarshalJSON(data []byte) error {
	type boHelper struct {
		ledger.Offer
		OwnerFunds      string          `json:"offer_funds"`
		TakerGetsFunded json.RawMessage `json:"taker_gets_funded,omitempty"`
		TakerPaysFunded json.RawMessage `json:"taker_pays_funded,omitempty"`
		Quality         string          `json:"quality"`
	}
	var h boHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*o = BookOffer{
		Offer:      h.Offer,
		OwnerFunds: h.OwnerFunds,
		Quality:    h.Quality,
	}
	var g, p types.CurrencyAmount
	g, err = types.UnmarshalCurrencyAmount(h.TakerGetsFunded)
	if err != nil {
		return err
	}
	o.TakerGetsFunded = g
	p, err = types.UnmarshalCurrencyAmount(h.TakerPaysFunded)
	if err != nil {
		return err
	}
	o.TakerPaysFunded = p

	return nil
}

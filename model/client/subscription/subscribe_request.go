package subscription

import (
	"fmt"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type SubscribeRequest struct {
	Streams          []string             `json:"streams,omitempty"`
	Accounts         []types.Address      `json:"accounts,omitempty"`
	AccountsProposed []types.Address      `json:"accounts_proposed,omitempty"`
	Books            []SubscribeOrderBook `json:"books,omitempty"`
	Url              string               `json:"url,omitempty"`
	UrlUsername      string               `json:"url_username,omitempty"`
	UrlPassword      string               `json:"url_password,omitempty"`
}

func (*SubscribeRequest) Method() string {
	return "subscribe"
}

type SubscribeOrderBook struct {
	TakerGets types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays types.IssuedCurrencyAmount `json:"taker_pays"`
	Taker     types.Address              `json:"taker"`
	Snapshot  bool                       `json:"snapshot,omitempty"`
	Both      bool                       `json:"both,omitempty"`
}

func (r *SubscribeRequest) Validate() error {
	for _, a := range r.Accounts {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("subscribe request accounts: %w", err)
		}
	}

	for _, a := range r.AccountsProposed {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("subscribe request accounts proposed: %w", err)
		}
	}

	return nil
}

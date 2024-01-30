package subscribe

import (
	"fmt"

	"github.com/xyield/xrpl-go/model/transactions/types"
)

type UnsubscribeRequest struct {
	Streams          []string               `json:"streams,omitempty"`
	Accounts         []types.Address        `json:"accounts,omitempty"`
	AccountsProposed []types.Address        `json:"accounts_proposed,omitempty"`
	Books            []UnsubscribeOrderBook `json:"books,omitempty"`
}

func (*UnsubscribeRequest) Method() string {
	return "unsubscribe"
}

type UnsubscribeOrderBook struct {
	TakerGets types.IssuedCurrencyAmount `json:"taker_gets"`
	TakerPays types.IssuedCurrencyAmount `json:"taker_pays"`
	Both      bool                       `json:"both,omitempty"`
}

func (r *UnsubscribeRequest) Validate() error {
	for _, a := range r.Accounts {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("unsubscribe request accounts: %w", err)
		}
	}

	for _, a := range r.AccountsProposed {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("unsubscribe request accounts proposed: %w", err)
		}
	}

	return nil
}

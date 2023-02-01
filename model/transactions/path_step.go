package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type PathStep struct {
	Account  Address `json:"account,omitempty"`
	Currency string  `json:"currency,omitempty"`
	Issuer   Address `json:"issuer,omitempty"`
}

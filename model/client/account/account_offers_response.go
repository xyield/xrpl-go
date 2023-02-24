package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountOffersResponse struct {
	Account            Address       `json:"account"`
	Offers             []OfferResult `json:"offers"`
	LedgerCurrentIndex LedgerIndex   `json:"ledger_current_index,omitempty"`
	LedgerIndex        LedgerIndex   `json:"ledger_index,omitempty"`
	LedgerHash         LedgerHash    `json:"ledger_hash,omitempty"`
	Marker             any           `json:"marker,omitempty"`
}

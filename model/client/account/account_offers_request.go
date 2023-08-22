package account

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountOffersRequest struct {
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*AccountOffersRequest) Method() string {
	return "account_offers"
}

func (r *AccountOffersRequest) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return fmt.Errorf("account offers request: %w", err)
	}

	if r.Limit != 0 && (r.Limit < 10 || r.Limit > 400) {
		return fmt.Errorf("account offers request: invalid limit, must be 10 <= limit <= 400")
	}

	return nil
}

func (r *AccountOffersRequest) UnmarshalJSON(data []byte) error {
	type aorHelper struct {
		Account     types.Address     `json:"account"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		Limit       int               `json:"limit,omitempty"`
		Marker      any               `json:"marker,omitempty"`
		Strict      bool              `json:"strict,omitempty"`
	}
	var h aorHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountOffersRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Marker:     h.Marker,
		Strict:     h.Strict,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

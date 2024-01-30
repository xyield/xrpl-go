package account

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountTransactionsRequest struct {
	Account        types.Address          `json:"account"`
	LedgerIndexMin int                    `json:"ledger_index_min,omitempty"`
	LedgerIndexMax int                    `json:"ledger_index_max,omitempty"`
	LedgerHash     common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex    common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Binary         bool                   `json:"binary,omitempty"`
	Forward        bool                   `json:"forward,omitempty"`
	Limit          int                    `json:"limit,omitempty"`
	Marker         any                    `json:"marker,omitempty"`
}

func (*AccountTransactionsRequest) Method() string {
	return "account_tx"
}

func (r *AccountTransactionsRequest) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return fmt.Errorf("account transactions request: %w", err)
	}

	if r.Limit != 0 && (r.Limit < 10 || r.Limit > 400) {
		return fmt.Errorf("account transactions request: invalid limit, must be 10 <= limit <= 400")
	}

	return nil
}

func (r *AccountTransactionsRequest) UnmarshalJSON(data []byte) error {
	type atrHelper struct {
		Account        types.Address     `json:"account"`
		LedgerIndexMin int               `json:"ledger_index_min,omitempty"`
		LedgerIndexMax int               `json:"ledger_index_max,omitempty"`
		LedgerHash     common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex    json.RawMessage   `json:"ledger_index,omitempty"`
		Binary         bool              `json:"binary,omitempty"`
		Forward        bool              `json:"forward,omitempty"`
		Limit          int               `json:"limit,omitempty"`
		Marker         any               `json:"marker,omitempty"`
	}
	var h atrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountTransactionsRequest{
		Account:        h.Account,
		LedgerIndexMin: h.LedgerIndexMin,
		LedgerIndexMax: h.LedgerIndexMax,
		LedgerHash:     h.LedgerHash,
		Binary:         h.Binary,
		Forward:        h.Forward,
		Limit:          h.Limit,
		Marker:         h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

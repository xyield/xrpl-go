package account

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountCurrenciesRequest struct {
	Account     types.Address          `json:"account"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Strict      bool                   `json:"strict,omitempty"`
}

func (*AccountCurrenciesRequest) Method() string {
	return "account_currencies"
}

func (r *AccountCurrenciesRequest) UnmarshalJSON(data []byte) error {
	type acrHelper struct {
		Account     types.Address     `json:"account"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		Strict      bool              `json:"strict,omitempty"`
	}
	var h acrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountCurrenciesRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Strict:     h.Strict,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

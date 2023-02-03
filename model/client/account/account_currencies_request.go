package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountCurrenciesRequest struct {
	Account     Address         `json:"account"`
	LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex LedgerSpecifier `json:"ledger_index,omitempty"`
	Strict      bool            `json:"strict,omitempty"`
}

func (r *AccountCurrenciesRequest) UnmarshalJSON(data []byte) error {
	type acrHelper struct {
		Account     Address         `json:"account"`
		LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
		LedgerIndex json.RawMessage `json:"ledger_index,omitempty"`
		Strict      bool            `json:"strict,omitempty"`
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

	i, err := UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

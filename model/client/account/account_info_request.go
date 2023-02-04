package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountInfoRequest struct {
	Account     Address         `json:"account"`
	LedgerIndex LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
	Queue       bool            `json:"queue,omitempty"`
	SignerList  bool            `json:"signer_list,omitempty"`
	Strict      bool            `json:"strict,omitempty"`
}

func (*AccountInfoRequest) Method() string {
	return "account_info"
}

func (r *AccountInfoRequest) UnmarshalJSON(data []byte) error {
	type airHelper struct {
		Account     Address         `json:"account"`
		LedgerIndex json.RawMessage `json:"ledger_index,omitempty"`
		LedgerHash  LedgerHash      `json:"ledger_hash,omitempty"`
		Queue       bool            `json:"queue,omitempty"`
		SignerList  bool            `json:"signer_list,omitempty"`
		Strict      bool            `json:"strict,omitempty"`
	}
	var h airHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountInfoRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Queue:      h.Queue,
		SignerList: h.SignerList,
		Strict:     h.Strict,
	}

	i, err := UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

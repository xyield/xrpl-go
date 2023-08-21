package account

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountNFTsRequest struct {
	Account     types.Address          `json:"account"`
	LedgerIndex common.LedgerSpecifier `json:"ledger_index,omitempty"`
	LedgerHash  common.LedgerHash      `json:"ledger_hash,omitempty"`
	Limit       int                    `json:"limit,omitempty"`
	Marker      any                    `json:"marker,omitempty"`
}

func (*AccountNFTsRequest) Method() string {
	return "account_nfts"
}

func (r *AccountNFTsRequest) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return err
	}

	if r.Limit != 0 && (r.Limit < 10 || r.Limit > 400) {
		return fmt.Errorf("invalid limit, must be 10 <= limit <= 400")
	}

	return nil
}

func (r *AccountNFTsRequest) UnmarshalJSON(data []byte) error {
	type anrHelper struct {
		Account     types.Address     `json:"account"`
		LedgerIndex json.RawMessage   `json:"ledger_index,omitempty"`
		LedgerHash  common.LedgerHash `json:"ledger_hash,omitempty"`
		Limit       int               `json:"limit,omitempty"`
		Marker      any               `json:"marker,omitempty"`
	}
	var h anrHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountNFTsRequest{
		Account:    h.Account,
		LedgerHash: h.LedgerHash,
		Limit:      h.Limit,
		Marker:     h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

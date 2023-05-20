package path

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositAuthorizedRequest struct {
	SourceAccount      types.Address          `json:"source_account"`
	DestinationAccount types.Address          `json:"destination_account"`
	LedgerHash         common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex        common.LedgerSpecifier `json:"ledger_index,omitempty"`
}

func (*DepositAuthorizedRequest) Method() string {
	return "deposit_authorized"
}

func (r *DepositAuthorizedRequest) UnmarshalJSON(data []byte) error {
	type darHelper struct {
		SourceAccount      types.Address     `json:"source_account"`
		DestinationAccount types.Address     `json:"destination_account"`
		LedgerHash         common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex        json.RawMessage   `json:"ledger_index,omitempty"`
	}
	var h darHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = DepositAuthorizedRequest{
		SourceAccount:      h.SourceAccount,
		DestinationAccount: h.DestinationAccount,
		LedgerHash:         h.LedgerHash,
	}
	var spec common.LedgerSpecifier
	spec, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = spec
	return nil
}

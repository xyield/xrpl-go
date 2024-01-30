package path

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
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

func (r *DepositAuthorizedRequest) Validate() error {
	if err := r.SourceAccount.Validate(); err != nil {
		return fmt.Errorf("deposit authorized source: %w", err)
	}
	if err := r.DestinationAccount.Validate(); err != nil {
		return fmt.Errorf("deposit authorized destination: %w", err)
	}

	return nil
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

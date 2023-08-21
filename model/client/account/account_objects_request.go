package account

import (
	"encoding/json"
	"fmt"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountObjectType string

const (
	CheckObject          AccountObjectType = "check"
	DepositPreauthObject AccountObjectType = "deposit_preauth"
	EscrowObject         AccountObjectType = "escrow"
	NFTOfferObject       AccountObjectType = "nft_offer"
	OfferObject          AccountObjectType = "offer"
	PaymentChannelObject AccountObjectType = "payment_channel"
	SignerListObject     AccountObjectType = "signer_list"
	StateObject          AccountObjectType = "state"
	TicketObject         AccountObjectType = "ticket"
)

type AccountObjectsRequest struct {
	Account              types.Address          `json:"account"`
	Type                 AccountObjectType      `json:"type,omitempty"`
	DeletionBlockersOnly bool                   `json:"deletion_blockers_only,omitempty"`
	LedgerHash           common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex          common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Limit                int                    `json:"limit,omitempty"`
	Marker               any                    `json:"marker,omitempty"`
}

func (*AccountObjectsRequest) Method() string {
	return "account_objects"
}

func (r *AccountObjectsRequest) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return err
	}

	if r.Limit != 0 && (r.Limit < 10 || r.Limit > 400) {
		return fmt.Errorf("invalid limit, must be 10 <= limit <= 400")
	}

	return nil
}

func (r *AccountObjectsRequest) UnmarshalJSON(data []byte) error {
	type aorHelper struct {
		Account              types.Address     `json:"account"`
		Type                 AccountObjectType `json:"type,omitempty"`
		DeletionBlockersOnly bool              `json:"deletion_blockers_only,omitempty"`
		LedgerHash           common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex          json.RawMessage   `json:"ledger_index,omitempty"`
		Limit                int               `json:"limit,omitempty"`
		Marker               any               `json:"marker,omitempty"`
	}
	var h aorHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = AccountObjectsRequest{
		Account:              h.Account,
		Type:                 h.Type,
		DeletionBlockersOnly: h.DeletionBlockersOnly,
		LedgerHash:           h.LedgerHash,
		Limit:                h.Limit,
		Marker:               h.Marker,
	}

	i, err := common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

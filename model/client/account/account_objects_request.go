package account

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountObjectType string

const (
	CheckObject          AccountObjectType = "check"
	DepositPreauthObject                   = "deposit_preauth"
	EscrowObject                           = "escrow"
	NFTOfferObject                         = "nft_offer"
	OfferObject                            = "offer"
	PaymentChannelObject                   = "payment_channel"
	SignerListObject                       = "signer_list"
	StateObject                            = "state"
	TicketObject                           = "ticket"
)

type AccountObjectsRequest struct {
	Account              Address           `json:"account"`
	Type                 AccountObjectType `json:"type,omitempty"`
	DeletionBlockersOnly bool              `json:"deletion_blockers_only,omitempty"`
	LedgerHash           LedgerHash        `json:"ledger_hash,omitempty"`
	LedgerIndex          LedgerSpecifier   `json:"ledger_index,omitempty"`
	Limit                int               `json:"limit,omitempty"`
	Marker               interface{}       `json:"marker,omitempty"`
}

func (*AccountObjectsRequest) Method() string {
	return "account_objects"
}

func (r *AccountObjectsRequest) UnmarshalJSON(data []byte) error {
	type aorHelper struct {
		Account              Address           `json:"account"`
		Type                 AccountObjectType `json:"type,omitempty"`
		DeletionBlockersOnly bool              `json:"deletion_blockers_only,omitempty"`
		LedgerHash           LedgerHash        `json:"ledger_hash,omitempty"`
		LedgerIndex          json.RawMessage   `json:"ledger_index,omitempty"`
		Limit                int               `json:"limit,omitempty"`
		Marker               interface{}       `json:"marker,omitempty"`
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

	i, err := UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.LedgerIndex = i
	return nil
}

package account

import (
	. "github.com/xyield/xrpl-go/model/client/common"
	. "github.com/xyield/xrpl-go/model/transactions"
)

type AccountObjectType string

const (
	CheckObject          AccountObjectType = "check"
	DepositPreauthObject                   = "deposit_preauth"
	EscrowObject                           = "escrow"
	OfferObject                            = "offer"
	PaymentChannelObject                   = "payment_channel"
	SignerListObject                       = "signer_list"
	TicketObject                           = "ticket"
	StateObject                            = "state"
)

type AccountObjectsRequest struct {
	Account              Address           `json:"account"`
	Type                 AccountObjectType `json:"type,omitempty"`
	DeletionBlockersOnly bool              `json:"deletion_blockers_only,omitempty"`
	LedgerHash           LedgerHash        `json:"ledger_hash,omitempty"`
	LedgerIndex          LedgerIndex       `json:"ledger_index,omitempty"`
	Limit                int               `json:"limit,omitempty"`
	Marker               interface{}       `json:"marker,omitempty'`
}

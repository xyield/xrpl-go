package ledger

import (
	"encoding/json"
	"fmt"
)

type LedgerEntryType string

const (
	AccountRootEntry    LedgerEntryType = "AccountRoot"
	AmendmentsEntry                     = "Amendments"
	CheckEntry                          = "Check"
	DepositPreauthEntry                 = "DepositPreauth"
	DirectoryNodeEntry                  = "DirectoryNode"
	EscrowEntry                         = "Escrow"
	FeeSettingsEntry                    = "FeeSettings"
	LedgerHashesEntry                   = "LedgerHashes"
	NegativeUNLEntry                    = "NegativeUNL"
	NFTokenOfferEntry                   = "NFTokenOffer"
	NFTokenPageEntry                    = "NFTokenPage"
	OfferEntry                          = "Offer"
	PayChannelEntry                     = "PayChannel"
	RippleStateEntry                    = "RippleState"
	SignerListEntry                     = "SignerList"
	TicketEntry                         = "Ticket"
)

type LedgerObject interface {
	EntryType() LedgerEntryType
}

func UnmarshalLedgerObject(data []byte) (LedgerObject, error) {
	type helper struct {
		LedgerEntryType
	}
	var h helper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}

	switch h.LedgerEntryType {
	case AccountRootEntry:
		var a AccountRoot
		err := json.Unmarshal(data, &a)
		return &a, err
	case AmendmentsEntry:
		var a Amendments
		err := json.Unmarshal(data, &a)
		return &a, err
	case CheckEntry:
		var c Check
		err := json.Unmarshal(data, &c)
		return &c, err
	case DepositPreauthEntry:
		var d DepositPreauth
		err := json.Unmarshal(data, &d)
		return &d, err
	case DirectoryNodeEntry:
		var d DirectoryNode
		err := json.Unmarshal(data, &d)
		return &d, err
	case EscrowEntry:
		var e Escrow
		err := json.Unmarshal(data, &e)
		return &e, err
	case FeeSettingsEntry:
		var f FeeSettings
		err := json.Unmarshal(data, &f)
		return &f, err
	case LedgerHashesEntry:
		var l LedgerHashes
		err := json.Unmarshal(data, &l)
		return &l, err
	case NegativeUNLEntry:
		var n NegativeUNL
		err := json.Unmarshal(data, &n)
		return &n, err
	case NFTokenOfferEntry:
		var n NFTokenOffer
		err := json.Unmarshal(data, &n)
		return &n, err
	case NFTokenPageEntry:
		var n NFTokenPage
		err := json.Unmarshal(data, &n)
		return &n, err
	case OfferEntry:
		var o Offer
		err := json.Unmarshal(data, &o)
		return &o, err
	case PayChannelEntry:
		var p PayChannel
		err := json.Unmarshal(data, &p)
		return &p, err
	case RippleStateEntry:
		var r RippleState
		err := json.Unmarshal(data, &r)
		return &r, err
	case SignerListEntry:
		var s SignerList
		err := json.Unmarshal(data, &s)
		return &s, err
	case TicketEntry:
		var t Ticket
		err := json.Unmarshal(data, &t)
		return &t, err
	}

	return nil, fmt.Errorf("Unsupported ledger object of type %s", h.LedgerEntryType)
}

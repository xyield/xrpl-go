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
	if data == nil || len(data) == 0 {
		return nil, nil
	}
	type helper struct {
		LedgerEntryType
	}
	var h helper
	if err := json.Unmarshal(data, &h); err != nil {
		return nil, err
	}
	var o LedgerObject
	switch h.LedgerEntryType {
	case AccountRootEntry:
		o = &AccountRoot{}
	case AmendmentsEntry:
		o = &Amendments{}
	case CheckEntry:
		o = &Check{}
	case DepositPreauthEntry:
		o = &DepositPreauth{}
	case DirectoryNodeEntry:
		o = &DirectoryNode{}
	case EscrowEntry:
		o = &Escrow{}
	case FeeSettingsEntry:
		o = &FeeSettings{}
	case LedgerHashesEntry:
		o = &LedgerHashes{}
	case NegativeUNLEntry:
		o = &NegativeUNL{}
	case NFTokenOfferEntry:
		o = &NFTokenOffer{}
	case NFTokenPageEntry:
		o = &NFTokenPage{}
	case OfferEntry:
		o = &Offer{}
	case PayChannelEntry:
		o = &PayChannel{}
	case RippleStateEntry:
		o = &RippleState{}
	case SignerListEntry:
		o = &SignerList{}
	case TicketEntry:
		o = &Ticket{}
	default:
		return nil, fmt.Errorf("Unsupported ledger object of type %s", h.LedgerEntryType)
	}
	if err := json.Unmarshal(data, o); err != nil {
		return nil, err
	}
	return o, nil

}

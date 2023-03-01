package ledger

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type LedgerEntryRequest struct {
	Binary         bool                   `json:"binary,omitempty"`
	LedgerHash     common.LedgerHash      `json:"ledger_hash,omitempty"`
	LedgerIndex    common.LedgerSpecifier `json:"ledger_index,omitempty"`
	Index          string                 `json:"index,omitempty"`
	AccountRoot    types.Address          `json:"account_root,omitempty"`
	Directory      EntryRequestOrString   `json:"directory,omitempty"`
	Offer          EntryRequestOrString   `json:"offer,omitempty"`
	RippleState    EntryRequestOrString   `json:"ripple_state,omitempty"`
	Check          string                 `json:"check,omitempty"`
	Escrow         EntryRequestOrString   `json:"escrow,omitempty"`
	PaymentChannel string                 `json:"payment_channel,omitempty"`
	DepositPreauth EntryRequestOrString   `json:"deposit_preauth,omitempty"`
	Ticket         EntryRequestOrString   `json:"ticket,omitempty"`
}

type EntryRequestOrString interface {
	LedgerEntryRequestField()
}

type EntryString string

func (EntryString) LedgerEntryRequestField() {}

type DirectoryEntryReq struct {
	SubIndex uint   `json:"sub_index,omitempty"`
	DirRoot  string `json:"dir_root,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

func (*DirectoryEntryReq) LedgerEntryRequestField() {}

type OfferEntryReq struct {
	Account types.Address `json:"account"`
	Seq     uint          `json:"seq"`
}

func (*OfferEntryReq) LedgerEntryRequestField() {}

type RippleStateEntryReq struct {
	Accounts []types.Address `json:"accounts"`
	Currency string          `json:"currency"`
}

func (*RippleStateEntryReq) LedgerEntryRequestField() {}

type EscrowEntryReq struct {
	Owner types.Address `json:"owner"`
	Seq   uint          `json:"seq"`
}

func (*EscrowEntryReq) LedgerEntryRequestField() {}

type DepositPreauthEntryReq struct {
	Owner      types.Address `json:"owner"`
	Authorized types.Address `json:"authorized"`
}

func (*DepositPreauthEntryReq) LedgerEntryRequestField() {}

type TicketEntryReq struct {
	Account   types.Address `json:"account"`
	TicketSeq int           `json:"ticket_seq"`
}

func (*TicketEntryReq) LedgerEntryRequestField() {}

func parseEntryRequestField(data []byte, target EntryRequestOrString) (EntryRequestOrString, error) {
	if data == nil || len(data) == 0 {
		return nil, nil
	}
	if data[0] == '"' {
		var s EntryString
		err := json.Unmarshal(data, &s)
		return s, err
	}
	err := json.Unmarshal(data, target)
	return target, err
}

func (r *LedgerEntryRequest) UnmarshalJSON(data []byte) error {
	type lerHelper struct {
		Binary         bool              `json:"binary,omitempty"`
		LedgerHash     common.LedgerHash `json:"ledger_hash,omitempty"`
		LedgerIndex    json.RawMessage   `json:"ledger_index,omitempty"`
		Index          string            `json:"index,omitempty"`
		AccountRoot    types.Address     `json:"account_root,omitempty"`
		Directory      json.RawMessage   `json:"directory,omitempty"`
		Offer          json.RawMessage   `json:"offer,omitempty"`
		RippleState    json.RawMessage   `json:"ripple_state,omitempty"`
		Check          string            `json:"check,omitempty"`
		Escrow         json.RawMessage   `json:"escrow,omitempty"`
		PaymentChannel string            `json:"payment_channel,omitempty"`
		DepositPreauth json.RawMessage   `json:"deposit_preauth,omitempty"`
		Ticket         json.RawMessage   `json:"ticket,omitempty"`
	}
	var h lerHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = LedgerEntryRequest{
		Binary:         h.Binary,
		LedgerHash:     h.LedgerHash,
		Index:          h.Index,
		AccountRoot:    h.AccountRoot,
		Check:          h.Check,
		PaymentChannel: h.PaymentChannel,
	}
	r.LedgerIndex, err = common.UnmarshalLedgerSpecifier(h.LedgerIndex)
	if err != nil {
		return err
	}
	r.Directory, err = parseEntryRequestField(h.Directory, &DirectoryEntryReq{})
	if err != nil {
		return err
	}
	r.Offer, err = parseEntryRequestField(h.Offer, &OfferEntryReq{})
	if err != nil {
		return err
	}
	r.RippleState, err = parseEntryRequestField(h.RippleState, &RippleStateEntryReq{})
	if err != nil {
		return err
	}
	r.Escrow, err = parseEntryRequestField(h.Escrow, &EscrowEntryReq{})
	if err != nil {
		return err
	}
	r.DepositPreauth, err = parseEntryRequestField(h.DepositPreauth, &DepositPreauthEntryReq{})
	if err != nil {
		return err
	}
	r.Ticket, err = parseEntryRequestField(h.Ticket, &TicketEntryReq{})

	return err
}

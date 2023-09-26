package ledger

import (
	"encoding/json"
	"fmt"

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
	RippleState    *RippleStateEntryReq   `json:"ripple_state,omitempty"`
	Check          string                 `json:"check,omitempty"`
	Escrow         EntryRequestOrString   `json:"escrow,omitempty"`
	PaymentChannel string                 `json:"payment_channel,omitempty"`
	DepositPreauth EntryRequestOrString   `json:"deposit_preauth,omitempty"`
	Ticket         EntryRequestOrString   `json:"ticket,omitempty"`
}

func (*LedgerEntryRequest) Method() string {
	return "ledger_entry"
}

func (r *LedgerEntryRequest) Validate() error {
	setCount := 0

	if r.Index != "" {
		setCount++
	}

	if r.AccountRoot != "" {
		setCount++
		if err := r.AccountRoot.Validate(); err != nil {
			return fmt.Errorf("ledger entry account root: %w", err)
		}
	}

	if r.Directory != nil {
		setCount++
		if err := r.Directory.Validate(); err != nil {
			return fmt.Errorf("ledger entry directory: %w", err)
		}
	}

	if r.Offer != nil {
		setCount++
		if err := r.Offer.Validate(); err != nil {
			return fmt.Errorf("ledger entry offer: %w", err)
		}
	}

	if r.RippleState != nil {
		setCount++
		if err := r.Offer.Validate(); err != nil {
			return fmt.Errorf("ledger entry ripple state: %w", err)
		}
	}

	if r.Check != "" {
		setCount++
	}

	if r.Escrow != nil {
		setCount++
		if err := r.Escrow.Validate(); err != nil {
			return fmt.Errorf("ledger entry escrow: %w", err)
		}
	}

	if r.PaymentChannel != "" {
		setCount++
	}

	if r.DepositPreauth != nil {
		setCount++
		if err := r.DepositPreauth.Validate(); err != nil {
			return fmt.Errorf("ledger entry deposit preauth: %w", err)
		}
	}

	if r.Ticket != nil {
		setCount++
		if err := r.Ticket.Validate(); err != nil {
			return fmt.Errorf("ledger entry ticket: %w", err)
		}
	}

	if setCount != 1 {
		return fmt.Errorf("ledger entry: exactly one ledger entry object may be requested, found %d", setCount)
	}

	return nil
}

type EntryRequestOrString interface {
	LedgerEntryRequestField()
	Validate() error
}

type EntryString string

func (EntryString) Validate() error {
	return nil
}

func (EntryString) LedgerEntryRequestField() {}

type DirectoryEntryReq struct {
	SubIndex uint   `json:"sub_index,omitempty"`
	DirRoot  string `json:"dir_root,omitempty"`
	Owner    string `json:"owner,omitempty"`
}

func (*DirectoryEntryReq) Validate() error {
	return nil
}

func (*DirectoryEntryReq) LedgerEntryRequestField() {}

type OfferEntryReq struct {
	Account types.Address `json:"account"`
	Seq     uint          `json:"seq"`
}

func (*OfferEntryReq) LedgerEntryRequestField() {}

func (r *OfferEntryReq) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return fmt.Errorf("offer entry account: %w", err)
	}
	return nil
}

type RippleStateEntryReq struct {
	Accounts []types.Address `json:"accounts"`
	Currency string          `json:"currency"`
}

func (*RippleStateEntryReq) LedgerEntryRequestField() {}

func (r *RippleStateEntryReq) Validate() error {
	if len(r.Accounts) != 2 {
		return fmt.Errorf("ripple state entry requires two accounts")
	}
	for i, a := range r.Accounts {
		if err := a.Validate(); err != nil {
			return fmt.Errorf("ripple state entry account %d: %w", i+1, err)
		}
	}
	return nil
}

type EscrowEntryReq struct {
	Owner types.Address `json:"owner"`
	Seq   uint          `json:"seq"`
}

func (*EscrowEntryReq) LedgerEntryRequestField() {}

func (r *EscrowEntryReq) Validate() error {
	if err := r.Owner.Validate(); err != nil {
		return fmt.Errorf("escrow entry owner: %w", err)
	}
	return nil
}

type DepositPreauthEntryReq struct {
	Owner      types.Address `json:"owner"`
	Authorized types.Address `json:"authorized"`
}

func (*DepositPreauthEntryReq) LedgerEntryRequestField() {}

func (r *DepositPreauthEntryReq) Validate() error {
	if err := r.Owner.Validate(); err != nil {
		return fmt.Errorf("deposit preauth entry owner: %w", err)
	}
	if err := r.Authorized.Validate(); err != nil {
		return fmt.Errorf("deposit preauth entry authorized: %w", err)
	}
	return nil
}

type TicketEntryReq struct {
	Account   types.Address `json:"account"`
	TicketSeq int           `json:"ticket_seq"`
}

func (*TicketEntryReq) LedgerEntryRequestField() {}

func (r *TicketEntryReq) Validate() error {
	if err := r.Account.Validate(); err != nil {
		return fmt.Errorf("ticket entry account: %w", err)
	}
	return nil
}

func parseEntryRequestField(data []byte, target EntryRequestOrString) (EntryRequestOrString, error) {
	if len(data) == 0 {
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
		Binary         bool                 `json:"binary,omitempty"`
		LedgerHash     common.LedgerHash    `json:"ledger_hash,omitempty"`
		LedgerIndex    json.RawMessage      `json:"ledger_index,omitempty"`
		Index          string               `json:"index,omitempty"`
		AccountRoot    types.Address        `json:"account_root,omitempty"`
		Directory      json.RawMessage      `json:"directory,omitempty"`
		Offer          json.RawMessage      `json:"offer,omitempty"`
		RippleState    *RippleStateEntryReq `json:"ripple_state,omitempty"`
		Check          string               `json:"check,omitempty"`
		Escrow         json.RawMessage      `json:"escrow,omitempty"`
		PaymentChannel string               `json:"payment_channel,omitempty"`
		DepositPreauth json.RawMessage      `json:"deposit_preauth,omitempty"`
		Ticket         json.RawMessage      `json:"ticket,omitempty"`
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
		RippleState:    h.RippleState,
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

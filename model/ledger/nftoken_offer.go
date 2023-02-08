package ledger

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type NFTokenOffer struct {
	Amount            CurrencyAmount
	Destination       Address `json:",omitempty"`
	Expiration        uint    `json:",omitempty"`
	Flags             uint
	LedgerEntryType   LedgerEntryType
	NFTokenID         Hash256
	NFTokenOfferNode  string `json:",omitempty"`
	Owner             Address
	OwnerNode         string `json:",omitempty"`
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
}

func (*NFTokenOffer) EntryType() LedgerEntryType {
	return NFTokenOfferEntry

}

func (n *NFTokenOffer) UnmarshalJSON(data []byte) error {
	type nftHelper struct {
		Amount            json.RawMessage
		Destination       Address
		Expiration        uint
		Flags             uint
		LedgerEntryType   LedgerEntryType
		NFTokenID         Hash256
		NFTokenOfferNode  string
		Owner             Address
		OwnerNode         string
		PreviousTxnID     Hash256
		PreviousTxnLgrSeq uint
	}
	var h nftHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*n = NFTokenOffer{
		Destination:       h.Destination,
		Expiration:        h.Expiration,
		Flags:             h.Flags,
		LedgerEntryType:   h.LedgerEntryType,
		NFTokenID:         h.NFTokenID,
		NFTokenOfferNode:  h.NFTokenOfferNode,
		Owner:             h.Owner,
		OwnerNode:         h.OwnerNode,
		PreviousTxnID:     h.PreviousTxnID,
		PreviousTxnLgrSeq: h.PreviousTxnLgrSeq,
	}
	amnt, err := UnmarshalCurrencyAmount(h.Amount)
	if err != nil {
		return err
	}
	n.Amount = amnt
	return nil
}

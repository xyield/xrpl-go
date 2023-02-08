package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

// TODO flags

type RippleState struct {
	Balance           IssuedCurrencyAmount
	Flags             uint
	HighLimit         IssuedCurrencyAmount
	HighNode          string
	HighQualityIn     uint `json:",omitempty"`
	HighQualityOut    uint `json:",omitempty"`
	LedgerEntryType   LedgerEntryType
	LowLimit          IssuedCurrencyAmount
	LowNode           string
	LowQualityIn      uint `json:",omitempty"`
	LowQualityOut     uint `json:",omitempty"`
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
}

func (*RippleState) EntryType() LedgerEntryType {
	return RippleStateEntry
}

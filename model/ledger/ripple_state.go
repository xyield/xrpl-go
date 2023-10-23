package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

// TODO flags

type RippleState struct {
	Balance           types.IssuedCurrencyAmount
	Flags             uint
	HighLimit         types.IssuedCurrencyAmount
	HighNode          string
	HighQualityIn     uint `json:",omitempty"`
	HighQualityOut    uint `json:",omitempty"`
	LedgerEntryType   LedgerEntryType
	LowLimit          types.IssuedCurrencyAmount
	LowNode           string
	LowQualityIn      uint `json:",omitempty"`
	LowQualityOut     uint `json:",omitempty"`
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
}

func (*RippleState) EntryType() LedgerEntryType {
	return RippleStateEntry
}

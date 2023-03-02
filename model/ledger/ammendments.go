package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type Amendments struct {
	Amendments      []types.Hash256 `json:",omitempty"`
	Flags           uint
	LedgerEntryType LedgerEntryType
	Majorities      []Majority `json:",omitempty"`
}

func (*Amendments) EntryType() LedgerEntryType {
	return AmendmentsEntry
}

type Majority struct {
	Amendment types.Hash256
	CloseTime uint
}

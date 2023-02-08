package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Amendments struct {
	Amendments      []Hash256 `json:",omitempty"`
	Flags           uint
	LedgerEntryType string
	Majorities      []Majority `json:",omitempty"`
}

func (*Amendments) EntryType() LedgerEntryType {
	return AmendmentsEntry
}

type Majority struct {
	Amendment Hash256
	CloseTime uint
}

package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Amendments struct {
	Amendments      []Hash256
	Flags           uint
	LedgerEntryType string
	Majorities      []Majority
}

func (*Amendments) EntryType() LedgerEntryType {
	return AmendmentsEntry
}

type Majority struct {
	Amendment Hash256
	CloseTime uint
}

package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type Amendments struct {
	Amendments      []types.Hash256 `json:",omitempty"`
	Flags           uint
	LedgerEntryType LedgerEntryType
	Majorities      []MajorityEntry `json:",omitempty"`
}

func (*Amendments) EntryType() LedgerEntryType {
	return AmendmentsEntry
}

type MajorityEntry struct {
	Majority Majority
}

type Majority struct {
	Amendment types.Hash256
	CloseTime uint
}

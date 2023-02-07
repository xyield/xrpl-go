package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Ammendments struct {
	Ammendments     []Hash256
	Flags           uint
	LedgerEntryType string
	Majorities      []Majority
}

type Majority struct {
	Amendment Hash256
	CloseTime uint
}

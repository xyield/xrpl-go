package common

import (
	"strconv"

	. "github.com/xyield/xrpl-go/model/transactions"
)

type LedgerSpecifier interface {
	Ledger() string
}

type LedgerIndex uint

func (l LedgerIndex) Ledger() string {
	return strconv.FormatUint(uint64(l), 10)
}

type LedgerTitle string

const (
	CURRENT   LedgerTitle = "current"
	VALIDATED             = "validated"
	CLOSED                = "closed"
)

func (l LedgerTitle) Ledger() string {
	return string(l)
}

type LedgerHash Hash256

func (l LedgerHash) Ledger() string {
	return string(l)
}

package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositPreauth struct {
	Account           Address
	Authorize         Address
	Flags             uint
	LedgerEntryType   string
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
}

func (*DepositPreauth) EntryType() LedgerEntryType {
	return DepositPreauthEntry
}

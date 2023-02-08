package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositPreauthObj struct {
	Account           Address
	Authorize         Address
	Flags             uint
	LedgerEntryType   string
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
}

func (*DepositPreauthObj) EntryType() LedgerEntryType {
	return DepositPreauthObjEntry
}

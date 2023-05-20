package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type DepositPreauthObj struct {
	Account           types.Address
	Authorize         types.Address
	Flags             uint
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
}

func (*DepositPreauthObj) EntryType() LedgerEntryType {
	return DepositPreauthObjEntry
}

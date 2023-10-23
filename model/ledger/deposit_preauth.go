package ledger

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type DepositPreauthObj struct {
	Account           types.Address
	Authorize         types.Address
	Flags             uint32
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
}

func (*DepositPreauthObj) EntryType() LedgerEntryType {
	return DepositPreauthObjEntry
}

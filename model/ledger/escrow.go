package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Escrow struct {
	Account           Address
	Amount            string
	CancelAfter       uint
	Condition         string
	Destination       Address
	DestinationNode   string
	DestinationTag    uint
	FinishAfter       uint
	Flags             uint
	LedgerEntryType   string
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	SourceTag         uint
}

func (*Escrow) EntryType() LedgerEntryType {
	return EscrowEntry
}

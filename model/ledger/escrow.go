package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type Escrow struct {
	Account           types.Address
	Amount            string
	CancelAfter       uint   `json:",omitempty"`
	Condition         string `json:",omitempty"`
	Destination       types.Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint   `json:",omitempty"`
	FinishAfter       uint   `json:",omitempty"`
	Flags             uint
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	SourceTag         uint `json:",omitempty"`
}

func (*Escrow) EntryType() LedgerEntryType {
	return EscrowEntry
}

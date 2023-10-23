package ledger

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type Escrow struct {
	Account           types.Address
	Amount            types.XRPCurrencyAmount
	CancelAfter       uint   `json:",omitempty"`
	Condition         string `json:",omitempty"`
	Destination       types.Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint   `json:",omitempty"`
	FinishAfter       uint   `json:",omitempty"`
	Flags             uint32
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	SourceTag         uint `json:",omitempty"`
}

func (*Escrow) EntryType() LedgerEntryType {
	return EscrowEntry
}

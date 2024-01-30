package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type Ticket struct {
	Account           types.Address
	Flags             uint
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	TicketSequence    uint
}

func (*Ticket) EntryType() LedgerEntryType {
	return TicketEntry
}

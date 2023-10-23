package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type Ticket struct {
	Account           types.Address
	Flags             uint32
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	TicketSequence    uint32
}

func (*Ticket) EntryType() LedgerEntryType {
	return TicketEntry
}

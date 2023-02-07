package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type Ticket struct {
	Account           Address
	Flags             uint
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	TicketSequence    uint
}

func (*Ticket) EntryType() LedgerEntryType {
	return TicketEntry
}

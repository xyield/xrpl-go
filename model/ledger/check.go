package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

// TODO verify format of SendMax
type Check struct {
	Account           Address
	Destination       Address
	DestinationNode   string
	DestinationTag    uint
	Expiration        uint
	Flags             uint
	InvoiceID         Hash256
	LedgerEntryType   string
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	SendMax           string
	Sequence          uint
	SourceTag         uint
}

package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

// TODO verify format of SendMax
type Check struct {
	Account           Address
	Destination       Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint   `json:",omitempty"`
	Expiration        uint   `json:",omitempty"`
	Flags             uint
	InvoiceID         Hash256 `json:",omitempty"`
	LedgerEntryType   string
	OwnerNode         string
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	SendMax           string
	Sequence          uint
	SourceTag         uint `json:",omitempty"`
}

func (*Check) EntryType() LedgerEntryType {
	return CheckEntry
}

package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

// TODO verify format of SendMax
type Check struct {
	Account           types.Address
	Destination       types.Address
	DestinationNode   string `json:",omitempty"`
	DestinationTag    uint   `json:",omitempty"`
	Expiration        uint   `json:",omitempty"`
	Flags             uint32
	InvoiceID         types.Hash256 `json:",omitempty"`
	LedgerEntryType   LedgerEntryType
	OwnerNode         string
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint
	SendMax           string
	Sequence          uint
	SourceTag         uint `json:",omitempty"`
}

func (*Check) EntryType() LedgerEntryType {
	return CheckEntry
}

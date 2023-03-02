package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountRoot struct {
	Account           types.Address
	AccountTxnID      types.Hash256           `json:",omitempty"`
	Balance           types.XRPCurrencyAmount `json:",omitempty"`
	BurnedNFTokens    uint32                  `json:",omitempty"`
	Domain            string                  `json:",omitempty"`
	EmailHash         types.Hash128           `json:",omitempty"`
	Flags             uint64
	LedgerEntryType   LedgerEntryType
	MessageKey        string        `json:",omitempty"`
	MintedNFTokens    uint32        `json:",omitempty"`
	NFTokenMinter     types.Address `json:",omitempty"`
	OwnerCount        uint64
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint64
	RegularKey        types.Address `json:",omitempty"`
	Sequence          uint64
	TicketCount       uint32        `json:",omitempty"`
	TickSize          uint8         `json:",omitempty"`
	TransferRate      uint32        `json:",omitempty"`
	Index             types.Hash256 `json:"index,omitempty"`
}

func (*AccountRoot) EntryType() LedgerEntryType {
	return AccountRootEntry
}

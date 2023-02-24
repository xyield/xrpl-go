package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type AccountRoot struct {
	Account           Address
	AccountTxnID      Hash256                 `json:",omitempty"`
	Balance           types.XRPCurrencyAmount `json:",omitempty"`
	BurnedNFTokens    uint32                  `json:",omitempty"`
	Domain            string                  `json:",omitempty"`
	EmailHash         Hash128                 `json:",omitempty"`
	Flags             uint64
	LedgerEntryType   LedgerEntryType
	MessageKey        string  `json:",omitempty"`
	MintedNFTokens    uint32  `json:",omitempty"`
	NFTokenMinter     Address `json:",omitempty"`
	OwnerCount        uint64
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint64
	RegularKey        Address `json:",omitempty"`
	Sequence          uint64
	TicketCount       uint32  `json:",omitempty"`
	TickSize          uint8   `json:",omitempty"`
	TransferRate      uint32  `json:",omitempty"`
	Index             Hash256 `json:"index,omitempty"`
}

func (*AccountRoot) EntryType() LedgerEntryType {
	return AccountRootEntry
}

package ledger

import (
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type AccountRoot struct {
	Account           types.Address
	AccountTxnID      types.Hash256           `json:",omitempty"`
	Balance           types.XRPCurrencyAmount `json:",omitempty"`
	BurnedNFTokens    uint32                  `json:",omitempty"`
	Domain            string                  `json:",omitempty"`
	EmailHash         types.Hash128           `json:",omitempty"`
	Flags             *types.Flag
	LedgerEntryType   LedgerEntryType
	MessageKey        string        `json:",omitempty"`
	MintedNFTokens    uint32        `json:",omitempty"`
	NFTokenMinter     types.Address `json:",omitempty"`
	OwnerCount        uint32
	PreviousTxnID     types.Hash256
	PreviousTxnLgrSeq uint32
	RegularKey        types.Address `json:",omitempty"`
	Sequence          uint32
	TicketCount       uint32 `json:",omitempty"`
	TickSize          uint8  `json:",omitempty"`
	TransferRate      uint32 `json:",omitempty"`
	// TODO determine if this is a required field
	//Index             types.Hash256 `json:"index,omitempty"`
}

func (*AccountRoot) EntryType() LedgerEntryType {
	return AccountRootEntry
}

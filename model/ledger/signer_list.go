package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type SignerListFlags uint32

const (
	LsfOneOwnerCount SignerListFlags = 0x00010000
)

type SignerList struct {
	LedgerEntryType   LedgerEntryType
	Flags             SignerListFlags
	PreviousTxnID     string
	PreviousTxnLgrSeq uint64
	OwnerNode         string
	SignerEntries     []SignerEntryWrapper
	SignerListID      uint64
	SignerQuorum      uint64
}

type SignerEntryWrapper struct {
	SignerEntry SignerEntry
}

type SignerEntry struct {
	Account       types.Address
	SignerWeight  uint64
	WalletLocator types.Hash256 `json:",omitempty"`
}

func (*SignerList) EntryType() LedgerEntryType {
	return SignerListEntry
}

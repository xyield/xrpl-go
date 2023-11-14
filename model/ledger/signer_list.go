package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type SignerListFlags uint32

const (
	LsfOneOwnerCount SignerListFlags = 0x00010000
)

func (f SignerListFlags) ToUint() uint32 {
	return uint32(f)
}

type SignerList struct {
	LedgerEntryType   LedgerEntryType
	Flags             SignerListFlags
	PreviousTxnID     string
	PreviousTxnLgrSeq uint32
	OwnerNode         string
	SignerEntries     []SignerEntryWrapper
	SignerListID      uint32
	SignerQuorum      uint32
}

type SignerEntryWrapper struct {
	SignerEntry SignerEntry
}

type SignerEntry struct {
	Account       types.Address
	SignerWeight  uint16
	WalletLocator types.Hash256 `json:",omitempty"`
}

func (*SignerList) EntryType() LedgerEntryType {
	return SignerListEntry
}

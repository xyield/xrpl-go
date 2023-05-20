package ledger

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
	SignerEntries     []SignerEntry
	SignerListId      uint64
	SignerQuorum      uint64
}

func (*SignerList) EntryType() LedgerEntryType {
	return SignerListEntry
}

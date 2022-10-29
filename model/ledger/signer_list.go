package ledger

type SignerListFlags uint32

const (
	LsfOneOwnerCount SignerListFlags = 0x00010000
)

type SignerList struct {
	LedgerEntryType   string          `json:"LedgerEntryType"`
	Flags             SignerListFlags `json:"Flags"`
	PreviousTxnID     string          `json:"PreviousTxnID"`
	PreviousTxnLgrSeq uint64          `json:"PreviousTxnLgrSeq"`
	OwnerNode         string          `json:"OwnerNode"`
	SignerEntries     []SignerEntry   `json:"SignerEntries"`
	SignerListId      uint64          `json:"SignerListID"`
	SignerQuorum      uint64          `json:"SignerQuorum"`
}

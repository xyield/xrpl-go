package ledger

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type LedgerHashes struct {
	FirstLedgerSequence uint32
	Flags               uint32
	Hashes              []types.Hash256
	LastLedgerSequence  uint32
	LedgerEntryType     LedgerEntryType
}

func (*LedgerHashes) EntryType() LedgerEntryType {
	return LedgerHashesEntry
}

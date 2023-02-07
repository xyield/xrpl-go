package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type LedgerHashes struct {
	FirstLedgerSequence uint
	Flags               uint
	Hashes              []Hash256
	LastLedgerSequence  uint
	LedgerEntryType     string
}

func (*LedgerHashes) EntryType() LedgerEntryType {
	return LedgerHashesEntry
}

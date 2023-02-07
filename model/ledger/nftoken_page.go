package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type NFTokenPage struct {
	LedgerEntryType   LedgerEntryType
	NextPageMin       Hash256
	PreviousPageMin   Hash256
	PreviousTxnID     Hash256
	PreviousTxnLgrSeq uint
	NFTokens          []NFToken
}

func (*NFTokenPage) EntryType() LedgerEntryType {
	return NFTokenPageEntry
}

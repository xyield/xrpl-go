package ledger

import . "github.com/xyield/xrpl-go/model/transactions/types"

type NFTokenPage struct {
	LedgerEntryType   LedgerEntryType
	NextPageMin       Hash256 `json:",omitempty"`
	PreviousPageMin   Hash256
	PreviousTxnID     Hash256   `json:",omitempty"`
	PreviousTxnLgrSeq uint      `json:",omitempty"`
	NFTokens          []NFToken `json:",omitempty"`
}

func (*NFTokenPage) EntryType() LedgerEntryType {
	return NFTokenPageEntry
}

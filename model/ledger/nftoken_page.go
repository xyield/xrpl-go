package ledger

import "github.com/xyield/xrpl-go/model/transactions/types"

type NFTokenPage struct {
	LedgerEntryType   LedgerEntryType
	NextPageMin       types.Hash256 `json:",omitempty"`
	PreviousPageMin   types.Hash256 `json:",omitempty"`
	PreviousTxnID     types.Hash256 `json:",omitempty"`
	PreviousTxnLgrSeq uint          `json:",omitempty"`
	NFTokens          []types.NFToken
}

func (*NFTokenPage) EntryType() LedgerEntryType {
	return NFTokenPageEntry
}

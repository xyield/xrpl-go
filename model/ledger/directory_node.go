package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type DirectoryNode struct {
	Flags             uint
	Indexes           []Hash256
	IndexNext         uint64
	IndexPrevious     uint64
	LedgerEntryType   string
	Owner             Address
	RootIndex         Hash256
	TakerGetsCurrency string
	TakerGetsIssuer   string
	TakerPaysCurrency string
	TakerPaysIssuer   string
}

func (*DirectoryNode) EntryType() LedgerEntryType {
	return DirectoryNodeEntry
}

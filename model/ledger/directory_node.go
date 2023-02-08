package ledger

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type DirectoryNode struct {
	Flags             uint
	Indexes           []Hash256
	IndexNext         uint64 `json:",omitempty"`
	IndexPrevious     uint64 `json:",omitempty"`
	LedgerEntryType   string
	Owner             Address `json:",omitempty"`
	RootIndex         Hash256
	TakerGetsCurrency string `json:",omitempty"`
	TakerGetsIssuer   string `json:",omitempty"`
	TakerPaysCurrency string `json:",omitempty"`
	TakerPaysIssuer   string `json:",omitempty"`
}

func (*DirectoryNode) EntryType() LedgerEntryType {
	return DirectoryNodeEntry
}

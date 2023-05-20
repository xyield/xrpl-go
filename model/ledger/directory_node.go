package ledger

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type DirectoryNode struct {
	Flags             uint
	Indexes           []types.Hash256
	IndexNext         string `json:",omitempty"`
	IndexPrevious     string `json:",omitempty"`
	LedgerEntryType   LedgerEntryType
	Owner             types.Address `json:",omitempty"`
	RootIndex         types.Hash256
	TakerGetsCurrency string `json:",omitempty"`
	TakerGetsIssuer   string `json:",omitempty"`
	TakerPaysCurrency string `json:",omitempty"`
	TakerPaysIssuer   string `json:",omitempty"`
}

func (*DirectoryNode) EntryType() LedgerEntryType {
	return DirectoryNodeEntry
}

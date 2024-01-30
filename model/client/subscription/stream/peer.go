package stream

import "github.com/CreatureDev/xrpl-go/model/client/common"

type PeerStatusStream struct {
	Type           StreamType         `json:"type"`
	Action         string             `json:"action"`
	Date           uint64             `json:"date"`
	LedgerHash     common.LedgerHash  `json:"ledger_hash"`
	LedgerIndex    common.LedgerIndex `json:"ledger_index"`
	LedgerIndexMax common.LedgerIndex `json:"ledger_index_max,omitempty"`
	LedgerIndexMin common.LedgerIndex `json:"ledger_index_min,omitempty"`
}

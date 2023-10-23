package account

import "github.com/CreatureDev/xrpl-go/model/transactions/types"

type QueueData struct {
	TxnCount           uint64                  `json:"txn_count"`
	AuthChangeQueued   bool                    `json:"auth_change_queued,omitempty"`
	LowestSequence     uint32                  `json:"lowest_sequence,omitempty"`
	HighestSequence    uint32                  `json:"highest_sequence,omitempty"`
	MaxSpendDropsTotal types.XRPCurrencyAmount `json:"max_spend_drops_total,omitempty"`
	Transactions       []QueueTransaction      `json:"transactions,omitempty"`
}

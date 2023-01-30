package account

type QueueData struct {
	TxnCount           uint64             `json:"txn_count"`
	AuthChangeQueued   bool               `json:"auth_change_queued"`
	LowestSequence     uint64             `json:"lowest_sequence"`
	HighestSequence    uint64             `json:"highest_sequence"`
	MaxSpendDropsTotal string             `json:"max_spend_drops_total"`
	Transactions       []QueueTransaction `json:"transactions"`
}

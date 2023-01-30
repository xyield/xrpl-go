package account

type QueueTransaction struct {
	AuthChange    bool   `json:"auth_chage"`
	Fee           string `json:"fee"`
	FeeLevel      string `json:"fee_level"`
	MaxSpendDrops string `json:"max_spend_drops"`
	Seq           int    `json:"seq"`
}

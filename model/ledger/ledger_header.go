package ledger

import "github.com/xyield/xrpl-go/model/transactions/types"

type LedgerHeader struct {
	LedgerIndex         string                  `json:"ledger_index"`
	LedgerHash          string                  `json:"ledger_hash"`
	AccountHash         string                  `json:"account_hash"`
	CloseTime           uint                    `json:"close_time"`
	Closed              bool                    `json:"closed"`
	ParentHash          string                  `json:"parent_hash"`
	TotalCoins          types.XRPCurrencyAmount `json:"total_coins"`
	TransactionHash     string                  `json:"transaction_hash"`
	CloseTimeResolution uint                    `json:"close_time_resolution"`
}

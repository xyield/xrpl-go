package server

import (
	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type FeeResponse struct {
	CurrentLedgerSize  string             `json:"current_ledger_size"`
	CurrentQueueSize   string             `json:"current_queue_size"`
	Drops              FeeDrops           `json:"drops"`
	ExpectedLedgerSize string             `json:"expected_ledger_size"`
	LedgerCurrentIndex common.LedgerIndex `json:"ledger_current_index"`
	Levels             FeeLevels          `json:"levels"`
	MaxQueueSize       string             `json:"max_queue_size"`
}

type FeeDrops struct {
	BaseFee       types.XRPCurrencyAmount `json:"base_fee"`
	MedianFee     types.XRPCurrencyAmount `json:"median_fee"`
	MinimumFee    types.XRPCurrencyAmount `json:"minimum_fee"`
	OpenLedgerFee types.XRPCurrencyAmount `json:"open_ledger_fee"`
}

type FeeLevels struct {
	MedianLevel     types.XRPCurrencyAmount `json:"median_level"`
	MinimumLevel    types.XRPCurrencyAmount `json:"minimum_level"`
	OpenLedgerLevel types.XRPCurrencyAmount `json:"open_ledger_level"`
	ReferenceLevel  types.XRPCurrencyAmount `json:"reference_level"`
}

package server

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestFeeResponse(t *testing.T) {
	s := FeeResponse{
		CurrentLedgerSize: "14",
		CurrentQueueSize:  "0",
		Drops: FeeDrops{
			BaseFee:       10,
			MedianFee:     11000,
			MinimumFee:    10,
			OpenLedgerFee: 10,
		},
		ExpectedLedgerSize: "24",
		LedgerCurrentIndex: 26575101,
		Levels: FeeLevels{
			MedianLevel:     281600,
			MinimumLevel:    256,
			OpenLedgerLevel: 256,
			ReferenceLevel:  256,
		},
		MaxQueueSize: "480",
	}

	j := `{
	"current_ledger_size": "14",
	"current_queue_size": "0",
	"drops": {
		"base_fee": "10",
		"median_fee": "11000",
		"minimum_fee": "10",
		"open_ledger_fee": "10"
	},
	"expected_ledger_size": "24",
	"ledger_current_index": 26575101,
	"levels": {
		"median_level": "281600",
		"minimum_level": "256",
		"open_ledger_level": "256",
		"reference_level": "256"
	},
	"max_queue_size": "480"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

package clio

import (
	"testing"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestLedgerRequest(t *testing.T) {
	s := LedgerRequest{
		LedgerIndex: common.VALIDATED,
	}
	j := `{
	"ledger_index": "validated",
	"full": false,
	"accounts": false,
	"transactions": false,
	"expand": false,
	"owner_funds": false,
	"binary": false,
	"queue": false,
	"diff": false
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerResponse(t *testing.T) {
	// TODO  test AccountState
	s := LedgerResponse{
		Ledger: ClioLedger{
			AccountHash:         "10EFE192F59B3DE2A2BE5BCE2CA5DC83D066105696FCFC24C055359AAEBD6941",
			CloseFlags:          0,
			CloseTime:           711134782,
			CloseTimeHuman:      "2022-Jul-14 17:26:22.000000000 UTC",
			CloseTimeResolution: 10,
			Closed:              true,
			LedgerHash:          "D3878EF6C92B84678AE2FBADC40961A161A128EA54AE59C2775CE076C2AE7A85",
			LedgerIndex:         "19977716",
			ParentCloseTime:     711134781,
			ParentHash:          "D6DE54039FE5A22D86CD522F1A9B7794E487B74D9B6B8CBDE23F240F434B6749",
			TotalCoins:          99987079398940307,
			TransactionHash:     "0000000000000000000000000000000000000000000000000000000000000000",
		},
		LedgerHash:  "D3878EF6C92B84678AE2FBADC40961A161A128EA54AE59C2775CE076C2AE7A85",
		LedgerIndex: 19977716,
		Validated:   true,
	}
	j := `{
	"ledger": {
		"account_hash": "10EFE192F59B3DE2A2BE5BCE2CA5DC83D066105696FCFC24C055359AAEBD6941",
		"close_flags": 0,
		"close_time": 711134782,
		"close_time_human": "2022-Jul-14 17:26:22.000000000 UTC",
		"close_time_resolution": 10,
		"closed": true,
		"ledger_hash": "D3878EF6C92B84678AE2FBADC40961A161A128EA54AE59C2775CE076C2AE7A85",
		"ledger_index": "19977716",
		"parent_close_time": 711134781,
		"parent_hash": "D6DE54039FE5A22D86CD522F1A9B7794E487B74D9B6B8CBDE23F240F434B6749",
		"total_coins": "99987079398940307",
		"transaction_hash": "0000000000000000000000000000000000000000000000000000000000000000"
	},
	"ledger_hash": "D3878EF6C92B84678AE2FBADC40961A161A128EA54AE59C2775CE076C2AE7A85",
	"ledger_index": 19977716,
	"validated": true
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

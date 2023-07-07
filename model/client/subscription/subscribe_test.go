package subscribe

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestSubscribeRequest(t *testing.T) {
	s := SubscribeRequest{
		Streams:          []string{"abc", "def"},
		Accounts:         []types.Address{"ghi", "jkl"},
		AccountsProposed: []types.Address{"bcd", "efg"},
		Books: []SubscribeOrderBook{
			{
				TakerGets: types.IssuedCurrencyAmount{
					Currency: "EUR",
				},
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "USD",
				},
				Taker: "jkl",
			},
		},
	}

	j := `{
	"streams": [
		"abc",
		"def"
	],
	"accounts": [
		"ghi",
		"jkl"
	],
	"accounts_proposed": [
		"bcd",
		"efg"
	],
	"books": [
		{
			"taker_gets": {
				"currency": "EUR"
			},
			"taker_pays": {
				"currency": "USD"
			},
			"taker": "jkl"
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestSubscribeResponse(t *testing.T) {
	s := SubscribeResponse{
		LoadBase:         10,
		LoadFactor:       10,
		Random:           "abc",
		ServerStatus:     "def",
		FeeBase:          123,
		FeeRef:           456,
		LedgerHash:       "ghi",
		LedgerIndex:      123,
		LedgerTime:       567,
		ReserveBase:      56,
		ReserveInc:       78,
		ValidatedLedgers: "123-456",
	}

	j := `{
	"load_base": 10,
	"load_factor": 10,
	"random": "abc",
	"server_status": "def",
	"fee_base": 123,
	"fee_ref": 456,
	"ledger_hash": "ghi",
	"ledger_index": 123,
	"ledger_time": 567,
	"reserve_base": 56,
	"reserve_inc": 78,
	"validated_ledgers": "123-456"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

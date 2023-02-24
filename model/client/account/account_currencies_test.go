package account

import (
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestAccountCurrenciesRequest(t *testing.T) {
	s := AccountCurrenciesRequest{
		Account:     "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		Strict:      true,
		LedgerIndex: LedgerIndex(1234),
	}

	j := `{
	"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"ledger_index": 1234,
	"strict": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountCurrenciesResponse(t *testing.T) {
	s := AccountCurrenciesResponse{
		LedgerHash:  "abc",
		LedgerIndex: 123,
		ReceiveCurrencies: []string{
			"USD",
			"JPY",
		},
		SendCurrencies: []string{
			"USD",
			"CAD",
		},
		Validated: true,
	}
	j := `{
	"ledger_hash": "abc",
	"ledger_index": 123,
	"receive_currencies": [
		"USD",
		"JPY"
	],
	"send_currencies": [
		"USD",
		"CAD"
	],
	"validated": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

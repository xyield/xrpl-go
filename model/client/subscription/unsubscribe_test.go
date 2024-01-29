package subscription

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestUnsubscribeRequest(t *testing.T) {
	s := UnsubscribeRequest{
		Streams:          []string{"ledger", "server", "transactions", "transactions_proposed"},
		Accounts:         []types.Address{"rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"},
		AccountsProposed: []types.Address{"rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"},
		Books: []UnsubscribeOrderBook{
			{
				TakerGets: types.IssuedCurrencyAmount{
					Issuer:   "rUQTpMqAF5jhykj4FExVeXakrZpiKF6cQV",
					Currency: "USD",
				},
				TakerPays: types.IssuedCurrencyAmount{
					Currency: "XRP",
				},
				Both: true,
			},
		},
	}

	j := `{
	"streams": [
		"ledger",
		"server",
		"transactions",
		"transactions_proposed"
	],
	"accounts": [
		"rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
	],
	"accounts_proposed": [
		"rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
	],
	"books": [
		{
			"taker_gets": {
				"issuer": "rUQTpMqAF5jhykj4FExVeXakrZpiKF6cQV",
				"currency": "USD"
			},
			"taker_pays": {
				"currency": "XRP"
			},
			"both": true
		}
	]
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

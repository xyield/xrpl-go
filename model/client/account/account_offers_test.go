package account

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestAccountOffersRequest(t *testing.T) {
	s := AccountOffersRequest{
		Account:     "abc",
		LedgerIndex: common.LedgerIndex(10),
		Marker:      "123",
	}
	j := `{
	"account": "abc",
	"ledger_index": 10,
	"marker": "123"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

func TestAccountOffersResponse(t *testing.T) {
	s := AccountOffersResponse{
		Account: "abc",
		Offers: []OfferResult{
			{Flags: 0,
				Sequence: 1,
				TakerGets: types.IssuedCurrencyAmount{
					Issuer:   "def",
					Currency: "USD",
					Value:    "100",
				},
				TakerPays:  types.XRPCurrencyAmount(1),
				Quality:    "1",
				Expiration: 50000000,
			},
		},
		LedgerCurrentIndex: 54321,
		LedgerIndex:        54320,
		LedgerHash:         "def",
	}
	j := `{
	"account": "abc",
	"offers": [
		{
			"flags": 0,
			"seq": 1,
			"taker_gets": {
				"issuer": "def",
				"currency": "USD",
				"value": "100"
			},
			"taker_pays": "1",
			"quality": "1",
			"expiration": 50000000
		}
	],
	"ledger_current_index": 54321,
	"ledger_index": 54320,
	"ledger_hash": "def"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

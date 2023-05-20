package account

import (
	"testing"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestAccountLinesRequest(t *testing.T) {
	s := AccountLinesRequest{
		Account:     "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		Peer:        "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
		LedgerHash:  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
		LedgerIndex: common.LedgerIndex(256),
		Limit:       10,
		Marker:      map[string]interface{}{"abc": "def"},
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"ledger_hash": "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
	"ledger_index": 256,
	"peer": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
	"limit": 10,
	"marker": {
		"abc": "def"
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountLinesResponse(t *testing.T) {
	s := AccountLinesReponse{
		Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		Lines: []TrustLine{
			{
				Account:    "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				Balance:    "123",
				Currency:   "USD",
				Limit:      "456",
				LimitPeer:  "10",
				QualityIn:  1,
				QualityOut: 2,
			},
		},
		LedgerCurrentIndex: 123,
		LedgerIndex:        345,
		LedgerHash:         "abc",
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"lines": [
		{
			"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
			"balance": "123",
			"currency": "USD",
			"limit": "456",
			"limit_peer": "10",
			"quality_in": 1,
			"quality_out": 2
		}
	],
	"ledger_current_index": 123,
	"ledger_index": 345,
	"ledger_hash": "abc"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

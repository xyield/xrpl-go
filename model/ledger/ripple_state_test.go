package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestRippleState(t *testing.T) {
	var s LedgerObject = &RippleState{
		Balance: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer:   "rrrrrrrrrrrrrrrrrrrrBZbvji",
			Value:    "-10",
		},
		Flags: 393216,
		HighLimit: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer:   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			Value:    "110",
		},
		HighNode:        "0000000000000000",
		LedgerEntryType: RippleStateEntry,
		LowLimit: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer:   "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
			Value:    "0",
		},
		LowNode:           "0000000000000000",
		PreviousTxnID:     "E3FE6EA3D48F0C2B639448020EA4F03D4F4F8FFDB243A852A0F59177921B4879",
		PreviousTxnLgrSeq: 14090896,
	}

	j := `{
	"Balance": {
		"issuer": "rrrrrrrrrrrrrrrrrrrrBZbvji",
		"currency": "USD",
		"value": "-10"
	},
	"Flags": 393216,
	"HighLimit": {
		"issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"currency": "USD",
		"value": "110"
	},
	"HighNode": "0000000000000000",
	"LedgerEntryType": "RippleState",
	"LowLimit": {
		"issuer": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
		"currency": "USD",
		"value": "0"
	},
	"LowNode": "0000000000000000",
	"PreviousTxnID": "E3FE6EA3D48F0C2B639448020EA4F03D4F4F8FFDB243A852A0F59177921B4879",
	"PreviousTxnLgrSeq": 14090896
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

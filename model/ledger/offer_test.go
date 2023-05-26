package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestOffer(t *testing.T) {
	var s LedgerObject = &Offer{
		Account:           "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
		BookDirectory:     "ACC27DE91DBA86FC509069EAF4BC511D73128B780F2E54BF5E07A369E2446000",
		BookNode:          "0000000000000000",
		Flags:             131072,
		LedgerEntryType:   OfferEntry,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
		PreviousTxnLgrSeq: 14524914,
		Sequence:          866,
		TakerGets: types.IssuedCurrencyAmount{
			Issuer:   "r9Dr5xwkeLegBeXq6ujinjSBLQzQ1zQGjH",
			Currency: "XAG",
			Value:    "37",
		},
		TakerPays: types.XRPCurrencyAmount(79550000000),
	}

	j := `{
	"Account": "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
	"BookDirectory": "ACC27DE91DBA86FC509069EAF4BC511D73128B780F2E54BF5E07A369E2446000",
	"BookNode": "0000000000000000",
	"Flags": 131072,
	"LedgerEntryType": "Offer",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	"PreviousTxnLgrSeq": 14524914,
	"Sequence": 866,
	"TakerPays": "79550000000",
	"TakerGets": {
		"issuer": "r9Dr5xwkeLegBeXq6ujinjSBLQzQ1zQGjH",
		"currency": "XAG",
		"value": "37"
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

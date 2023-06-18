package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestPayChannel(t *testing.T) {
	var s LedgerObject = &PayChannel{
		Account:           "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
		Destination:       "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		Amount:            types.XRPCurrencyAmount(4325800),
		Balance:           types.XRPCurrencyAmount(2323423),
		PublicKey:         "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A",
		SettleDelay:       3600,
		Expiration:        536027313,
		CancelAfter:       536891313,
		SourceTag:         0,
		DestinationTag:    1002341,
		DestinationNode:   "0000000000000000",
		Flags:             0,
		LedgerEntryType:   PayChannelEntry,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
		PreviousTxnLgrSeq: 14524914,
	}

	j := `{
	"Account": "rBqb89MRQJnMPq8wTwEbtz4kvxrEDfcYvt",
	"Amount": "4325800",
	"Balance": "2323423",
	"CancelAfter": 536891313,
	"Destination": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
	"DestinationTag": 1002341,
	"DestinationNode": "0000000000000000",
	"Expiration": 536027313,
	"Flags": 0,
	"LedgerEntryType": "PayChannel",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "F0AB71E777B2DA54B86231E19B82554EF1F8211F92ECA473121C655BFC5329BF",
	"PreviousTxnLgrSeq": 14524914,
	"PublicKey": "32D2471DB72B27E3310F355BB33E339BF26F8392D5A93D3BC0FC3B566612DA0F0A",
	"SettleDelay": 3600
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

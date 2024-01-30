package ledger

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/test"
)

func TestCheck(t *testing.T) {
	var s LedgerObject = &Check{
		Account:           "rUn84CUYbNjRoTQ6mSW7BVJPSVJNLb1QLo",
		Destination:       "rfkE1aSy9G8Upk4JssnwBxhEv5p4mn2KTy",
		DestinationNode:   "0000000000000000",
		DestinationTag:    1,
		Expiration:        570113521,
		Flags:             0,
		InvoiceID:         "46060241FABCF692D4D934BA2A6C4427CD4279083E38C77CBE642243E43BE291",
		LedgerEntryType:   CheckEntry,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "5463C6E08862A1FAE5EDAC12D70ADB16546A1F674930521295BC082494B62924",
		PreviousTxnLgrSeq: 6,
		SendMax:           "100000000",
		Sequence:          2,
	}

	j := `{
	"Account": "rUn84CUYbNjRoTQ6mSW7BVJPSVJNLb1QLo",
	"Destination": "rfkE1aSy9G8Upk4JssnwBxhEv5p4mn2KTy",
	"DestinationNode": "0000000000000000",
	"DestinationTag": 1,
	"Expiration": 570113521,
	"Flags": 0,
	"InvoiceID": "46060241FABCF692D4D934BA2A6C4427CD4279083E38C77CBE642243E43BE291",
	"LedgerEntryType": "Check",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "5463C6E08862A1FAE5EDAC12D70ADB16546A1F674930521295BC082494B62924",
	"PreviousTxnLgrSeq": 6,
	"SendMax": "100000000",
	"Sequence": 2
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

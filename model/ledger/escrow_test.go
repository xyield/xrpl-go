package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestEscrow(t *testing.T) {
	var s LedgerObject
	s = &Escrow{
		Account:           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		Amount:            types.XRPCurrencyAmount(10000),
		CancelAfter:       545440232,
		Condition:         "A0258020A82A88B2DF843A54F58772E4A3861866ECDB4157645DD9AE528C1D3AEEDABAB6810120",
		Destination:       "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		DestinationNode:   "0000000000000000",
		DestinationTag:    23480,
		FinishAfter:       545354132,
		Flags:             0,
		LedgerEntryType:   EscrowEntry,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "C44F2EB84196B9AD820313DBEBA6316A15C9A2D35787579ED172B87A30131DA7",
		PreviousTxnLgrSeq: 28991004,
		SourceTag:         11747,
	}

	j := `{
	"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
	"Amount": "10000",
	"CancelAfter": 545440232,
	"Condition": "A0258020A82A88B2DF843A54F58772E4A3861866ECDB4157645DD9AE528C1D3AEEDABAB6810120",
	"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"DestinationNode": "0000000000000000",
	"DestinationTag": 23480,
	"FinishAfter": 545354132,
	"Flags": 0,
	"LedgerEntryType": "Escrow",
	"OwnerNode": "0000000000000000",
	"PreviousTxnID": "C44F2EB84196B9AD820313DBEBA6316A15C9A2D35787579ED172B87A30131DA7",
	"PreviousTxnLgrSeq": 28991004,
	"SourceTag": 11747
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

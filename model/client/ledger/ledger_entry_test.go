package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestLedgerEntryRequest(t *testing.T) {
	s := LedgerEntryRequest{
		LedgerIndex: common.VALIDATED,
		Directory: &DirectoryEntryReq{
			Owner: "abc",
		},
	}
	j := `{
	"ledger_index": "validated",
	"directory": {
		"owner": "abc"
	}
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestLedgerEntryResponse(t *testing.T) {
	s := LedgerEntryResponse{
		Index:       "13F1A95D7AAB7108D5CE7EEAF504B2894B8C674E6D68499076441C4837282BF8",
		LedgerIndex: 61966146,
		Node: &ledger.AccountRoot{
			Account:           "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			AccountTxnID:      "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
			Balance:           types.XRPCurrencyAmount(424021949),
			Domain:            "6D64756F31332E636F6D",
			EmailHash:         "98B4375E1D753E5B91627516F6D70977",
			Flags:             types.SetFlag(9568256),
			LedgerEntryType:   ledger.AccountRootEntry,
			MessageKey:        "0000000000000000000000070000000300",
			OwnerCount:        12,
			PreviousTxnID:     "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
			PreviousTxnLgrSeq: 61965653,
			RegularKey:        "rD9iJmieYHn8jTtPjwwkW2Wm9sVDvPXLoJ",
			Sequence:          385,
			TransferRate:      4294967295,
		},
	}
	j := `{
	"index": "13F1A95D7AAB7108D5CE7EEAF504B2894B8C674E6D68499076441C4837282BF8",
	"ledger_index": 61966146,
	"node": {
		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"AccountTxnID": "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
		"Balance": "424021949",
		"Domain": "6D64756F31332E636F6D",
		"EmailHash": "98B4375E1D753E5B91627516F6D70977",
		"Flags": 9568256,
		"LedgerEntryType": "AccountRoot",
		"MessageKey": "0000000000000000000000070000000300",
		"OwnerCount": 12,
		"PreviousTxnID": "4E0AA11CBDD1760DE95B68DF2ABBE75C9698CEB548BEA9789053FCB3EBD444FB",
		"PreviousTxnLgrSeq": 61965653,
		"RegularKey": "rD9iJmieYHn8jTtPjwwkW2Wm9sVDvPXLoJ",
		"Sequence": 385,
		"TransferRate": 4294967295
	}
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

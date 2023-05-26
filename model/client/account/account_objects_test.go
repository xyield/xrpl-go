package account

import (
	"testing"

	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/model/ledger"
	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestAccountObjectsRequest(t *testing.T) {
	s := AccountObjectsRequest{
		Account:     "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
		Type:        SignerListObject,
		LedgerIndex: common.LedgerIndex(123),
	}

	j := `{
	"account": "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
	"type": "signer_list",
	"ledger_index": 123
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountObjectsResponse(t *testing.T) {
	s := AccountObjectsResponse{
		Account: "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
		AccountObjects: []ledger.LedgerObject{
			&ledger.SignerList{
				LedgerEntryType:   ledger.SignerListEntry,
				Flags:             0,
				PreviousTxnID:     "abc",
				PreviousTxnLgrSeq: 123,
				OwnerNode:         "bob",
				SignerEntries:     []ledger.SignerEntry{},
				SignerListID:      213,
				SignerQuorum:      0,
			},
			&ledger.NFTokenOffer{
				Amount:            types.XRPCurrencyAmount(1),
				Destination:       "abc",
				Expiration:        1,
				Flags:             0,
				LedgerEntryType:   ledger.NFTokenOfferEntry,
				NFTokenID:         "qwe",
				Owner:             "asd",
				OwnerNode:         "zxc",
				PreviousTxnID:     "",
				PreviousTxnLgrSeq: 123,
			},
		},
		LedgerHash:         "rty",
		LedgerIndex:        123,
		LedgerCurrentIndex: 1234,
		Limit:              1,
		Validated:          true,
	}

	j := `{
	"account": "rsuHaTvJh1bDmDoxX9QcKP7HEBSBt4XsHx",
	"account_objects": [
		{
			"LedgerEntryType": "SignerList",
			"Flags": 0,
			"PreviousTxnID": "abc",
			"PreviousTxnLgrSeq": 123,
			"OwnerNode": "bob",
			"SignerEntries": [],
			"SignerListID": 213,
			"SignerQuorum": 0
		},
		{
			"Amount": "1",
			"Destination": "abc",
			"Expiration": 1,
			"Flags": 0,
			"LedgerEntryType": "NFTokenOffer",
			"NFTokenID": "qwe",
			"Owner": "asd",
			"OwnerNode": "zxc",
			"PreviousTxnID": "",
			"PreviousTxnLgrSeq": 123
		}
	],
	"ledger_hash": "rty",
	"ledger_index": 123,
	"ledger_current_index": 1234,
	"limit": 1,
	"validated": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

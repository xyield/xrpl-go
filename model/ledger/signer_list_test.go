package ledger

import (
	"testing"

	"github.com/xyield/xrpl-go/test"
)

func TestSignerList(t *testing.T) {
	var s LedgerObject = &SignerList{
		Flags:             0,
		LedgerEntryType:   SignerListEntry,
		OwnerNode:         "0000000000000000",
		PreviousTxnID:     "5904C0DC72C58A83AEFED2FFC5386356AA83FCA6A88C89D00646E51E687CDBE4",
		PreviousTxnLgrSeq: 16061435,
		SignerEntries: []SignerEntryWrapper{
			{
				SignerEntry: SignerEntry{
					Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
					SignerWeight: 2,
				},
			},
			{
				SignerEntry: SignerEntry{
					Account:      "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
					SignerWeight: 1,
				},
			},
			{
				SignerEntry: SignerEntry{
					Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
					SignerWeight: 1,
				},
			},
		},
		SignerListID: 0,
		SignerQuorum: 3,
	}

	j := `{
	"LedgerEntryType": "SignerList",
	"Flags": 0,
	"PreviousTxnID": "5904C0DC72C58A83AEFED2FFC5386356AA83FCA6A88C89D00646E51E687CDBE4",
	"PreviousTxnLgrSeq": 16061435,
	"OwnerNode": "0000000000000000",
	"SignerEntries": [
		{
			"SignerEntry": {
				"Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"SignerWeight": 2
			}
		},
		{
			"SignerEntry": {
				"Account": "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
				"SignerWeight": 1
			}
		},
		{
			"SignerEntry": {
				"Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
				"SignerWeight": 1
			}
		}
	],
	"SignerListID": 0,
	"SignerQuorum": 3
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

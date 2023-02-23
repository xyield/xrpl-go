package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/ledger"
	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestSignerListSetTx(t *testing.T) {
	s := SignerListSet{
		BaseTx: BaseTx{
			Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			TransactionType: SignerListSetTx,
			Fee:             types.XRPCurrencyAmount(12),
		},
		SignerQuorum: 3,
		SignerEntries: []ledger.SignerWrapper{
			{
				SignerEntry: ledger.SignerEntry{
					Account:      "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
					SignerWeight: 2,
				},
			},
			{
				SignerEntry: ledger.SignerEntry{
					Account:      "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
					SignerWeight: 1,
				},
			},
			{
				SignerEntry: ledger.SignerEntry{
					Account:      "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
					SignerWeight: 1,
				},
			},
		},
	}

	j := `{
	"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
	"TransactionType": "SignerListSet",
	"Fee": "12",
	"SignerQuorum": 3,
	"SignerEntries": [
		{
			"SignerEntry": {
				"Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
				"SignerWeight": 2
			}
		},
		{
			"SignerEntry": {
				"Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
				"SignerWeight": 1
			}
		},
		{
			"SignerEntry": {
				"Account": "raKEEVSGnKSD9Zyvxu4z6Pqpm4ABH8FS6n",
				"SignerWeight": 1
			}
		}
	]
}`
	if err := test.SerializeAndDeserialize(s, j); err != nil {
		t.Error(err)
	}

	tx, err := UnmarshalTx(json.RawMessage(j))
	if err != nil {
		t.Errorf("UnmarshalTx error: %s", err.Error())
	}
	if !reflect.DeepEqual(tx, &s) {
		t.Error("UnmarshalTx result differs from expected")
	}
}

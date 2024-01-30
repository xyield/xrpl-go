package transactions

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestTransactionEntryRequest(t *testing.T) {
	s := TransactionEntryRequest{
		LedgerHash:  "abc",
		LedgerIndex: common.CLOSED,
		TxHash:      "def",
	}
	j := `{
	"ledger_hash": "abc",
	"ledger_index": "closed",
	"tx_hash": "def"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

func TestTransactionEntryResponse(t *testing.T) {
	s := TransactionEntryResponse{
		LedgerIndex: 56865245,
		LedgerHash:  "793E56131D8D4ABFB27FA383BFC44F2978B046E023FF46C588D7E0C874C2472A",
		Metadata: transactions.TxObjMeta{
			AffectedNodes: []transactions.AffectedNode{
				{
					CreatedNode: &transactions.CreatedNode{
						LedgerEntryType: ledger.NegativeUNLEntry,
						LedgerIndex:     "123",
						NewFields: &ledger.NegativeUNL{
							Flags: 16,
						},
					},
				},
			},
			TransactionIndex:  10,
			TransactionResult: "abc",
		},
		Tx: &transactions.AccountDelete{
			BaseTx: transactions.BaseTx{
				Account:         "abc",
				TransactionType: transactions.AccountDeleteTx,
				Fee:             types.XRPCurrencyAmount(1),
				Sequence:        10,
				SigningPubKey:   "we2",
				TxnSignature:    "1a",
			},
			Destination: "def",
		},
	}
	j := `{
	"ledger_index": 56865245,
	"ledger_hash": "793E56131D8D4ABFB27FA383BFC44F2978B046E023FF46C588D7E0C874C2472A",
	"metadata": {
		"AffectedNodes": [
			{
				"CreatedNode": {
					"LedgerEntryType": "NegativeUNL",
					"LedgerIndex": "123",
					"NewFields": {
						"Flags": 16,
						"LedgerEntryType": ""
					}
				}
			}
		],
		"TransactionIndex": 10,
		"TransactionResult": "abc"
	},
	"tx_json": {
		"Account": "abc",
		"TransactionType": "AccountDelete",
		"Fee": "1",
		"Sequence": 10,
		"SigningPubKey": "we2",
		"TxnSignature": "1a",
		"Destination": "def"
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

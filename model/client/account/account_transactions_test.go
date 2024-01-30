package account

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/ledger"
	tx "github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestAccountTransactionsRequest(t *testing.T) {
	s := AccountTransactionsRequest{
		Account:        "tre",
		LedgerIndexMin: 1,
		LedgerIndexMax: 10,
		LedgerHash:     "123",
		LedgerIndex:    common.CLOSED,
		Binary:         true,
		Forward:        false,
		Limit:          0,
	}
	j := `{
	"account": "tre",
	"ledger_index_min": 1,
	"ledger_index_max": 10,
	"ledger_hash": "123",
	"ledger_index": "closed",
	"binary": true
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountTransactionsResponse(t *testing.T) {
	s := AccountTransactionsResponse{
		Account:        "tre",
		LedgerIndexMin: 1,
		LedgerIndexMax: 10,
		Limit:          2,
		Transactions: []AccountTransaction{
			{
				LedgerIndex: 1,
				Meta:        tx.TxBinMeta("txbinmeta"),
				Tx: &tx.AccountDelete{
					BaseTx: tx.BaseTx{
						Account:         "abc",
						TransactionType: "AccountDelete",
						Fee:             types.XRPCurrencyAmount(1),
						Sequence:        10,
						SigningPubKey:   "we2",
						TxnSignature:    "1a",
					},
					Destination: "def",
				},
				TxBlob:    "wow",
				Validated: true,
			},
			{
				LedgerIndex: 2,
				Meta: tx.TxObjMeta{
					AffectedNodes: []tx.AffectedNode{
						{
							CreatedNode: &tx.CreatedNode{
								LedgerEntryType: ledger.NegativeUNLEntry,
								LedgerIndex:     "123",
								NewFields: &ledger.NegativeUNL{
									Flags: 16,
								},
							},
						},
					},
					PartialDeliveredAmount: types.XRPCurrencyAmount(256),
					TransactionIndex:       10,
					TransactionResult:      "abc",
					DeliveredAmount: types.IssuedCurrencyAmount{
						Issuer:   "abc",
						Currency: "USD",
						Value:    "100",
					},
				},
				Tx: &tx.AccountDelete{
					BaseTx: tx.BaseTx{
						Account:         "abc",
						TransactionType: "AccountDelete",
						Fee:             types.XRPCurrencyAmount(1),
						Sequence:        10,
						SigningPubKey:   "we2",
						TxnSignature:    "1a",
					},
					Destination: "def",
				},
				TxBlob:    "wow",
				Validated: true,
			},
		},
		Validated: true,
	}

	j := `{
	"account": "tre",
	"ledger_index_min": 1,
	"ledger_index_max": 10,
	"limit": 2,
	"transactions": [
		{
			"ledger_index": 1,
			"meta": "txbinmeta",
			"tx": {
				"Account": "abc",
				"TransactionType": "AccountDelete",
				"Fee": "1",
				"Sequence": 10,
				"SigningPubKey": "we2",
				"TxnSignature": "1a",
				"Destination": "def"
			},
			"tx_blob": "wow",
			"validated": true
		},
		{
			"ledger_index": 2,
			"meta": {
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
				"DeliveredAmount": "256",
				"TransactionIndex": 10,
				"TransactionResult": "abc",
				"delivered_amount": {
					"issuer": "abc",
					"currency": "USD",
					"value": "100"
				}
			},
			"tx": {
				"Account": "abc",
				"TransactionType": "AccountDelete",
				"Fee": "1",
				"Sequence": 10,
				"SigningPubKey": "we2",
				"TxnSignature": "1a",
				"Destination": "def"
			},
			"tx_blob": "wow",
			"validated": true
		}
	],
	"validated": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

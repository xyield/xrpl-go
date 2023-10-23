package transactions

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/ledger"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

func TestTxResponse(t *testing.T) {
	s := TxResponse{
		Tx: &transactions.OfferCreate{
			BaseTx: transactions.BaseTx{
				TransactionType:    transactions.OfferCreateTx,
				Account:            "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
				Fee:                types.XRPCurrencyAmount(12),
				Flags:              types.SetFlag(0),
				LastLedgerSequence: 56865248,
				Sequence:           5037710,
				SigningPubKey:      "03B51A3EDF70E4098DA7FB053A01C5A6A0A163A30ED1445F14F87C7C3295FCB3BE",
				TxnSignature:       "3045022100A5023A0E64923616FCDB6D664F569644C7C9D1895772F986CD6B981B515B02A00220530C973E9A8395BC6FE2484948D2751F6B030FC7FB8575D1BFB406368AD554D9",
			},
			TakerGets: types.XRPCurrencyAmount(15000000000),
			TakerPays: types.IssuedCurrencyAmount{
				Currency: "CNY",
				Issuer:   "rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y",
				Value:    "20160.75",
			},
			OfferSequence: 5037708,
		},

		Hash:        "C53ECF838647FA5A4C780377025FEC7999AB4182590510CA461444B207AB74A9",
		Date:        648248020,
		LedgerIndex: 56865245,
		Meta: transactions.TxObjMeta{
			AffectedNodes: []transactions.AffectedNode{
				{
					ModifiedNode: &transactions.ModifiedNode{
						FinalFields: &ledger.DirectoryNode{
							Flags:             types.SetFlag(0),
							RootIndex:         "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
							TakerGetsCurrency: "0000000000000000000000000000000000000000",
							TakerGetsIssuer:   "0000000000000000000000000000000000000000",
							TakerPaysCurrency: "000000000000000000000000434E590000000000",
							TakerPaysIssuer:   "CED6E99370D5C00EF4EBF72567DA99F5661BFB3A",
						},
						LedgerEntryType: ledger.DirectoryNodeEntry,
						LedgerIndex:     "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
					},
				},
				{
					ModifiedNode: &transactions.ModifiedNode{
						FinalFields: &ledger.AccountRoot{
							Account:    "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
							Balance:    types.XRPCurrencyAmount(10404767991),
							Flags:      types.SetFlag(0),
							OwnerCount: 3,
							Sequence:   5037711,
						},
						LedgerEntryType: ledger.AccountRootEntry,
						LedgerIndex:     "1DECD9844E95FFBA273F1B94BA0BF2564DDF69F2804497A6D7837B52050174A2",
						PreviousFields: &ledger.AccountRoot{
							Balance:  types.XRPCurrencyAmount(10404768003),
							Sequence: 5037710,
						},
						PreviousTxnID:     "4DC47B246B5EB9CCE92ABA8C482479E3BF1F946CABBEF74CA4DE36521D5F9008",
						PreviousTxnLgrSeq: 56865244,
					},
				},
				{
					DeletedNode: &transactions.DeletedNode{
						FinalFields: &ledger.Offer{
							Account:           "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
							BookDirectory:     "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
							BookNode:          "0000000000000000",
							Flags:             0,
							OwnerNode:         "0000000000000000",
							PreviousTxnID:     "8F5FF57B404827F12BDA7561876A13C3E3B3095CBF75334DBFB5F227391A660C",
							PreviousTxnLgrSeq: 56865244,
							Sequence:          5037708,
							TakerGets:         types.XRPCurrencyAmount(15000000000),
							TakerPays: types.IssuedCurrencyAmount{
								Currency: "CNY",
								Issuer:   "rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y",
								Value:    "20160.75",
							},
						},
						LedgerEntryType: ledger.OfferEntry,
						LedgerIndex:     "26AAE6CA8D29E28A47C92ADF22D5D96A0216F0551E16936856DDC8CB1AAEE93B",
					},
				},
				{
					ModifiedNode: &transactions.ModifiedNode{
						FinalFields: &ledger.DirectoryNode{
							Flags:         types.SetFlag(0),
							IndexNext:     "0000000000000000",
							IndexPrevious: "0000000000000000",
							Owner:         "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
							RootIndex:     "47FAF5D102D8CE655574F440CDB97AC67C5A11068BB3759E87C2B9745EE94548",
						},
						LedgerEntryType: ledger.DirectoryNodeEntry,
						LedgerIndex:     "47FAF5D102D8CE655574F440CDB97AC67C5A11068BB3759E87C2B9745EE94548",
					},
				},
				{
					CreatedNode: &transactions.CreatedNode{
						LedgerEntryType: ledger.OfferEntry,
						LedgerIndex:     "8BAEE3C7DE04A568E96007420FA11ABD0BC9AE44D35932BB5640E9C3FB46BC9B",
						NewFields: &ledger.Offer{
							Account:       "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
							BookDirectory: "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
							Sequence:      5037710,
							TakerGets:     types.XRPCurrencyAmount(15000000000),
							TakerPays: types.IssuedCurrencyAmount{
								Currency: "CNY",
								Issuer:   "rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y",
								Value:    "20160.75",
							},
						},
					},
				},
			},
			TransactionIndex:  0,
			TransactionResult: "tesSUCCESS",
		},
		Validated: true,
	}

	j := `{
	"Account": "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
	"Fee": "12",
	"Flags": 0,
	"LastLedgerSequence": 56865248,
	"OfferSequence": 5037708,
	"Sequence": 5037710,
	"SigningPubKey": "03B51A3EDF70E4098DA7FB053A01C5A6A0A163A30ED1445F14F87C7C3295FCB3BE",
	"TakerGets": "15000000000",
	"TakerPays": {
		"currency": "CNY",
		"issuer": "rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y",
		"value": "20160.75"
	},
	"TransactionType": "OfferCreate",
	"TxnSignature": "3045022100A5023A0E64923616FCDB6D664F569644C7C9D1895772F986CD6B981B515B02A00220530C973E9A8395BC6FE2484948D2751F6B030FC7FB8575D1BFB406368AD554D9",
	"date": 648248020,
	"hash": "C53ECF838647FA5A4C780377025FEC7999AB4182590510CA461444B207AB74A9",
	"ledger_index": 56865245,
	"meta": {
		"AffectedNodes": [
			{
				"ModifiedNode": {
					"FinalFields": {
						"Flags": 0,
						"RootIndex": "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
						"TakerGetsCurrency": "0000000000000000000000000000000000000000",
						"TakerGetsIssuer": "0000000000000000000000000000000000000000",
						"TakerPaysCurrency": "000000000000000000000000434E590000000000",
						"TakerPaysIssuer": "CED6E99370D5C00EF4EBF72567DA99F5661BFB3A"
					},
					"LedgerEntryType": "DirectoryNode",
					"LedgerIndex": "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400"
				}
			},
			{
				"ModifiedNode": {
					"FinalFields": {
						"Account": "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
						"Balance": "10404767991",
						"Flags": 0,
						"OwnerCount": 3,
						"Sequence": 5037711
					},
					"LedgerEntryType": "AccountRoot",
					"LedgerIndex": "1DECD9844E95FFBA273F1B94BA0BF2564DDF69F2804497A6D7837B52050174A2",
					"PreviousFields": {
						"Balance": "10404768003",
						"Sequence": 5037710
					},
					"PreviousTxnID": "4DC47B246B5EB9CCE92ABA8C482479E3BF1F946CABBEF74CA4DE36521D5F9008",
					"PreviousTxnLgrSeq": 56865244
				}
			},
			{
				"DeletedNode": {
					"FinalFields": {
						"Account": "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
						"BookDirectory": "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
						"BookNode": "0000000000000000",
						"Flags": 0,
						"OwnerNode": "0000000000000000",
						"PreviousTxnID": "8F5FF57B404827F12BDA7561876A13C3E3B3095CBF75334DBFB5F227391A660C",
						"PreviousTxnLgrSeq": 56865244,
						"Sequence": 5037708,
						"TakerGets": "15000000000",
						"TakerPays": {
							"currency": "CNY",
							"issuer": "rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y",
							"value": "20160.75"
						}
					},
					"LedgerEntryType": "Offer",
					"LedgerIndex": "26AAE6CA8D29E28A47C92ADF22D5D96A0216F0551E16936856DDC8CB1AAEE93B"
				}
			},
			{
				"ModifiedNode": {
					"FinalFields": {
						"Flags": 0,
						"IndexNext": "0000000000000000",
						"IndexPrevious": "0000000000000000",
						"Owner": "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
						"RootIndex": "47FAF5D102D8CE655574F440CDB97AC67C5A11068BB3759E87C2B9745EE94548"
					},
					"LedgerEntryType": "DirectoryNode",
					"LedgerIndex": "47FAF5D102D8CE655574F440CDB97AC67C5A11068BB3759E87C2B9745EE94548"
				}
			},
			{
				"CreatedNode": {
					"LedgerEntryType": "Offer",
					"LedgerIndex": "8BAEE3C7DE04A568E96007420FA11ABD0BC9AE44D35932BB5640E9C3FB46BC9B",
					"NewFields": {
						"Account": "rhhh49pFH96roGyuC4E5P4CHaNjS1k8gzM",
						"BookDirectory": "02BAAC1E67C1CE0E96F0FA2E8061020536CEDD043FEB0FF54F04C66806CF7400",
						"Sequence": 5037710,
						"TakerGets": "15000000000",
						"TakerPays": {
							"currency": "CNY",
							"issuer": "rKiCet8SdvWxPXnAgYarFUXMh1zCPz432Y",
							"value": "20160.75"
						}
					}
				}
			}
		],
		"TransactionIndex": 0,
		"TransactionResult": "tesSUCCESS"
	},
	"validated": true
}`
	// Due to structure of meta nodes (inclusion of empty fields that server omits) do not test Marshal, only Unmarshal
	if err := Deserialize(s, j); err != nil {
		t.Error(err)
	}
}

func Deserialize(s TxResponse, d string) error {
	var decode TxResponse
	err := json.Unmarshal([]byte(d), &decode)
	if err != nil {
		return err
	}
	if !reflect.DeepEqual(s, decode) {
		if !reflect.DeepEqual(s.Tx, decode.Tx) {
			fmt.Println("Bad tx")
		}
		fmt.Printf("%+v\n", s.Tx.(*transactions.OfferCreate))
		fmt.Printf("%+v\n", decode.Tx.(*transactions.OfferCreate))
		return fmt.Errorf("json decoding does not match expected struct")
	}
	return nil

}

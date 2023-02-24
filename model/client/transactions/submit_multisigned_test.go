package transactions

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions"
	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestSubmitMultisignedRequest(t *testing.T) {
	s := SubmitMultisignedRequest{
		Tx: &transactions.Payment{
			Amount: types.IssuedCurrencyAmount{
				Currency: "USD",
				Issuer:   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				Value:    "1",
			},
			Destination: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			BaseTx: transactions.BaseTx{
				Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				TransactionType: "Payment",
				Fee:             types.XRPCurrencyAmount(10000),
				Sequence:        360,
				Flags:           2147483648,
				SigningPubKey:   "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
				TxnSignature:    "304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F8580",
				Signers: []transactions.Signer{
					{
						SignerData: transactions.SignerData{
							Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SigningPubKey: "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF",
							TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
						},
					},
					{
						SignerData: transactions.SignerData{
							Account:       "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SigningPubKey: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
							TxnSignature:  "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
						},
					},
				},
			},
		},
		FailHard: true,
	}

	j := `{
	"tx_json": {
		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"TransactionType": "Payment",
		"Fee": "10000",
		"Sequence": 360,
		"Flags": 2147483648,
		"Signers": [
			{
				"Signer": {
					"Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
					"TxnSignature": "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
					"SigningPubKey": "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF"
				}
			},
			{
				"Signer": {
					"Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
					"TxnSignature": "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
					"SigningPubKey": "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B"
				}
			}
		],
		"SigningPubKey": "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
		"TxnSignature": "304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F8580",
		"Amount": {
			"issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			"currency": "USD",
			"value": "1"
		},
		"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX"
	},
	"fail_hard": true
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

func TestSubmitMultisignedResponse(t *testing.T) {
	s := SubmitMultisignedResponse{
		EngineResult:        "tesSUCCESS",
		EngineResultCode:    0,
		EngineResultMessage: "The transaction was applied. Only final in a validated ledger.",
		TxBlob:              "1200002280000000240000016861D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA9684000000000002710732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB7446304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F858081144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754",
		Tx: &transactions.Payment{
			Amount: types.IssuedCurrencyAmount{
				Currency: "USD",
				Issuer:   "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				Value:    "1",
			},
			Destination: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			BaseTx: transactions.BaseTx{
				Account:         "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				TransactionType: "Payment",
				Fee:             types.XRPCurrencyAmount(10000),
				Sequence:        360,
				Flags:           2147483648,
				SigningPubKey:   "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
				TxnSignature:    "304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F8580",
				Signers: []transactions.Signer{
					{
						SignerData: transactions.SignerData{
							Account:       "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
							SigningPubKey: "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF",
							TxnSignature:  "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
						},
					},
					{
						SignerData: transactions.SignerData{
							Account:       "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
							SigningPubKey: "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B",
							TxnSignature:  "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
						},
					},
				},
			},
		},
	}
	j := `{
	"engine_result": "tesSUCCESS",
	"engine_result_code": 0,
	"engine_result_message": "The transaction was applied. Only final in a validated ledger.",
	"tx_blob": "1200002280000000240000016861D4838D7EA4C6800000000000000000000000000055534400000000004B4E9C06F24296074F7BC48F92A97916C6DC5EA9684000000000002710732103AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB7446304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F858081144B4E9C06F24296074F7BC48F92A97916C6DC5EA983143E9D4A2B8AA0780F682D136F7A56D6724EF53754",
	"tx_json": {
		"Account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		"TransactionType": "Payment",
		"Fee": "10000",
		"Sequence": 360,
		"Flags": 2147483648,
		"Signers": [
			{
				"Signer": {
					"Account": "rsA2LpzuawewSBQXkiju3YQTMzW13pAAdW",
					"TxnSignature": "30450221009C195DBBF7967E223D8626CA19CF02073667F2B22E206727BFE848FF42BEAC8A022048C323B0BED19A988BDBEFA974B6DE8AA9DCAE250AA82BBD1221787032A864E5",
					"SigningPubKey": "02B3EC4E5DD96029A647CFA20DA07FE1F85296505552CCAC114087E66B46BD77DF"
				}
			},
			{
				"Signer": {
					"Account": "rUpy3eEg8rqjqfUoLeBnZkscbKbFsKXC3v",
					"TxnSignature": "30440220680BBD745004E9CFB6B13A137F505FB92298AD309071D16C7B982825188FD1AE022004200B1F7E4A6A84BB0E4FC09E1E3BA2B66EBD32F0E6D121A34BA3B04AD99BC1",
					"SigningPubKey": "028FFB276505F9AC3F57E8D5242B386A597EF6C40A7999F37F1948636FD484E25B"
				}
			}
		],
		"SigningPubKey": "03AB40A0490F9B7ED8DF29D246BF2D6269820A0EE7742ACDD457BEA7C7D0931EDB",
		"TxnSignature": "304402200E5C2DD81FDF0BE9AB2A8D797885ED49E804DBF28E806604D878756410CA98B102203349581946B0DDA06B36B35DBC20EDA27552C1F167BCF5C6ECFF49C6A46F8580",
		"Amount": {
			"issuer": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			"currency": "USD",
			"value": "1"
		},
		"Destination": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX"
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

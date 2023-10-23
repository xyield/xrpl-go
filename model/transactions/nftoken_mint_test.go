package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestNFTokenMintTx(t *testing.T) {
	s := NFTokenMint{
		BaseTx: BaseTx{
			Account:         "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			TransactionType: NFTokenMintTx,
			Fee:             types.XRPCurrencyAmount(10),
			Flags:           types.SetFlag(8),
			Memos: []MemoWrapper{
				{
					Memo: Memo{
						MemoData: "72656E74",
						MemoType: "687474703A2F2F6578616D706C652E636F6D2F6D656D6F2F67656E65726963",
					},
				},
			},
		},
		NFTokenTaxon: 0,
		TransferFee:  314,
		URI:          "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469",
	}

	j := `{
	"Account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
	"TransactionType": "NFTokenMint",
	"Fee": "10",
	"Flags": 8,
	"Memos": [
		{
			"Memo": {
				"MemoData": "72656E74",
				"MemoType": "687474703A2F2F6578616D706C652E636F6D2F6D656D6F2F67656E65726963"
			}
		}
	],
	"NFTokenTaxon": 0,
	"TransferFee": 314,
	"URI": "697066733A2F2F62616679626569676479727A74357366703775646D37687537367568377932366E6634646675796C71616266336F636C67747179353566627A6469"
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
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

package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestNFTokenBurnTransaction(t *testing.T) {
	s := NFTokenBurn{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: NFTokenBurnTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		NFTokenID: "B1C2D3",
	}

	j := `{
	"Account": "abcdef",
	"TransactionType": "NFTokenBurn",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"NFTokenID": "B1C2D3"
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

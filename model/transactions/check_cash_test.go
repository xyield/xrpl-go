package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestCheckCashTransaction(t *testing.T) {
	s := CheckCash{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: CheckCashTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		CheckID: "A1B2C3D4A6D1",
		Amount: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer:   "abcd",
			Value:    "1234",
		},
		DeliverMin: types.IssuedCurrencyAmount{
			Currency: "USD",
			Issuer:   "abcd",
			Value:    "1200",
		},
	}
	j := `{
	"Account": "abcdef",
	"TransactionType": "CheckCash",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"CheckID": "A1B2C3D4A6D1",
	"Amount": {
		"issuer": "abcd",
		"currency": "USD",
		"value": "1234"
	},
	"DeliverMin": {
		"issuer": "abcd",
		"currency": "USD",
		"value": "1200"
	}
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

package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestPaymentTx(t *testing.T) {
	s := Payment{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentTx,
			Fee:             types.XRPCurrencyAmount(1000),
			Flags:           types.SetFlag(262144),
		},
		Amount: types.IssuedCurrencyAmount{
			Issuer:   "def",
			Currency: "USD",
			Value:    "1",
		},
		Destination: "hij",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "Payment",
	"Fee": "1000",
	"Flags": 262144,
	"Amount": {
		"issuer": "def",
		"currency": "USD",
		"value": "1"
	},
	"Destination": "hij"
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

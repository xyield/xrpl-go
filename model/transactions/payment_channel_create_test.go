package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestPaymentChannelCreate(t *testing.T) {
	s := PaymentChannelCreate{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentChannelCreateTx,
		},
		Amount:      types.XRPCurrencyAmount(1000),
		Destination: "def",
		SettleDelay: 10,
		PublicKey:   "abcd",
	}

	j := `{
	"Account": "abc",
	"TransactionType": "PaymentChannelCreate",
	"Amount": "1000",
	"Destination": "def",
	"SettleDelay": 10,
	"PublicKey": "abcd"
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

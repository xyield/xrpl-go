package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestPaymentChannelFund(t *testing.T) {
	s := PaymentChannelFund{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: PaymentChannelFundTx,
		},
		Channel: "ABACAD",
		Amount:  types.XRPCurrencyAmount(2000),
	}

	j := `{
	"Account": "abc",
	"TransactionType": "PaymentChannelFund",
	"Channel": "ABACAD",
	"Amount": "2000"
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

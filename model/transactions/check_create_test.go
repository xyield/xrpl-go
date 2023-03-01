package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestCheckCreateTransaction(t *testing.T) {
	s := CheckCreate{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: CheckCreateTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Destination: "plmokn",
		SendMax: types.IssuedCurrencyAmount{
			Issuer:   "ijnuhb",
			Currency: "JPY",
			Value:    "1234",
		},
		DestinationTag: 10,
		Expiration:     98765,
		InvoiceID:      "A1232452DBC",
	}
	j := `{
	"Account": "abcdef",
	"TransactionType": "CheckCreate",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Destination": "plmokn",
	"SendMax": {
		"issuer": "ijnuhb",
		"currency": "JPY",
		"value": "1234"
	},
	"DestinationTag": 10,
	"Expiration": 98765,
	"InvoiceID": "A1232452DBC"
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

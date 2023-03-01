package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestEscrowCreateTransaction(t *testing.T) {
	s := EscrowCreate{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: EscrowCreateTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Amount:      types.XRPCurrencyAmount(5000),
		Destination: "defghi",
		CancelAfter: 9000000,
	}

	j := `{
	"Account": "abcdef",
	"TransactionType": "EscrowCreate",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Amount": "5000",
	"Destination": "defghi",
	"CancelAfter": 9000000
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

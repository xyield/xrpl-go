package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestEscrowFinishTransaction(t *testing.T) {
	s := EscrowFinish{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: EscrowFinishTx,
			Fee:             types.XRPCurrencyAmount(1),
			Sequence:        1234,
			SigningPubKey:   "ghijk",
			TxnSignature:    "A1B2C3D4E5F6",
		},
		Owner:         "abcdef",
		OfferSequence: 1232,
	}

	j := `{
	"Account": "abcdef",
	"TransactionType": "EscrowFinish",
	"Fee": "1",
	"Sequence": 1234,
	"SigningPubKey": "ghijk",
	"TxnSignature": "A1B2C3D4E5F6",
	"Owner": "abcdef",
	"OfferSequence": 1232
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

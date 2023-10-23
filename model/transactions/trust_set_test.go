package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestTrustSetTx(t *testing.T) {
	s := TrustSet{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    TrustSetTx,
			Fee:                types.XRPCurrencyAmount(12),
			Flags:              types.SetFlag(262144),
			Sequence:           12,
			LastLedgerSequence: 8007750,
		},
		LimitAmount: types.IssuedCurrencyAmount{
			Issuer:   "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
			Currency: "USD",
			Value:    "100",
		},
	}

	j := `{
	"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"TransactionType": "TrustSet",
	"Fee": "12",
	"Sequence": 12,
	"Flags": 262144,
	"LastLedgerSequence": 8007750,
	"LimitAmount": {
		"issuer": "rsP3mgGb2tcYUrxiLFiHJiQXhsziegtwBc",
		"currency": "USD",
		"value": "100"
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

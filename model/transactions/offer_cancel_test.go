package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestOfferCancelTx(t *testing.T) {
	s := OfferCancel{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    OfferCancelTx,
			Fee:                types.XRPCurrencyAmount(12),
			Sequence:           7,
			LastLedgerSequence: 7108629,
		},
		OfferSequence: 6,
	}
	j := `{
	"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"TransactionType": "OfferCancel",
	"Fee": "12",
	"Sequence": 7,
	"LastLedgerSequence": 7108629,
	"OfferSequence": 6
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

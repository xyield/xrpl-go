package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestOfferCreateTx(t *testing.T) {
	s := OfferCreate{
		BaseTx: BaseTx{
			Account:            "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			TransactionType:    OfferCreateTx,
			Fee:                types.XRPCurrencyAmount(12),
			Sequence:           8,
			LastLedgerSequence: 7108682,
		},
		TakerGets: types.XRPCurrencyAmount(6000000),
		TakerPays: types.IssuedCurrencyAmount{
			Issuer:   "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
			Currency: "GKO",
			Value:    "2",
		},
	}

	j := `{
	"Account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
	"TransactionType": "OfferCreate",
	"Fee": "12",
	"Sequence": 8,
	"LastLedgerSequence": 7108682,
	"TakerGets": "6000000",
	"TakerPays": {
		"issuer": "ruazs5h1qEsqpke88pcqnaseXdm6od2xc",
		"currency": "GKO",
		"value": "2"
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

package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestNFTokenCancelOfferTx(t *testing.T) {
	s := NFTokenCancelOffer{
		BaseTx: BaseTx{
			Account:         "abcdef",
			TransactionType: NFTokenCancelOfferTx,
		},
		NFTokenOffers: []types.Hash256{
			"ABC",
			"DEF",
		},
	}
	j := `{
	"Account": "abcdef",
	"TransactionType": "NFTokenCancelOffer",
	"NFTokenOffers": [
		"ABC",
		"DEF"
	]
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

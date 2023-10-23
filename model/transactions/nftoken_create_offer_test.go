package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestNFTokenCreateOfferTx(t *testing.T) {
	s := NFTokenCreateOffer{
		BaseTx: BaseTx{
			Account:         "rs8jBmmfpwgmrSPgwMsh7CvKRmRt1JTVSX",
			TransactionType: NFTokenCreateOfferTx,
			Flags:           types.SetFlag(1),
		},
		NFTokenID: "000100001E962F495F07A990F4ED55ACCFEEF365DBAA76B6A048C0A200000007",
		Amount:    types.XRPCurrencyAmount(1000000),
	}

	j := `{
	"Account": "rs8jBmmfpwgmrSPgwMsh7CvKRmRt1JTVSX",
	"TransactionType": "NFTokenCreateOffer",
	"Flags": 1,
	"NFTokenID": "000100001E962F495F07A990F4ED55ACCFEEF365DBAA76B6A048C0A200000007",
	"Amount": "1000000"
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

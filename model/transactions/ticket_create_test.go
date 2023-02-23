package transactions

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestTicketCreateTx(t *testing.T) {
	s := TicketCreate{
		BaseTx: BaseTx{
			Account:         "abc",
			TransactionType: TicketCreateTx,
			Fee:             types.XRPCurrencyAmount(10),
			Sequence:        50,
		},
		TicketCount: 5,
	}

	j := `{
	"Account": "abc",
	"TransactionType": "TicketCreate",
	"Fee": "10",
	"Sequence": 50,
	"TicketCount": 5
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

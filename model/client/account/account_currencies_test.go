package account

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
)

func TestAccountCurrenciesRequest(t *testing.T) {
	accountCurrenciesStruct := AccountCurrenciesRequest{
		Account:     "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		Strict:      true,
		LedgerIndex: LedgerIndex(1234),
	}

	accountCurrenciesJson := `{
	"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"ledger_index": 1234,
	"strict": true
}`

	d, _ := json.MarshalIndent(accountCurrenciesStruct, "", "\t")
	if string(d) != accountCurrenciesJson {
		t.Error("json encoding does not match expected")
	}

	var conv AccountCurrenciesRequest
	if err := json.Unmarshal(d, &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountCurrenciesStruct) {
		t.Error("json decoding does not match expected struct")
	}
	if err := json.Unmarshal([]byte(accountCurrenciesJson), &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountCurrenciesStruct) {
		t.Error("json decoding does not match expected struct")
	}
}

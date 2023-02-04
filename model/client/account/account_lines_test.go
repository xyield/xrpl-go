package account

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
)

func TestAccountLinesRequest(t *testing.T) {
	accountLinesStruct := AccountLinesRequest{
		Account:     "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		Peer:        "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
		LedgerHash:  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
		LedgerIndex: LedgerIndex(256),
		Limit:       10,
		Marker:      map[string]interface{}{"abc": "def"},
	}

	accountLinesJson := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"ledger_hash": "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
	"ledger_index": 256,
	"peer": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
	"limit": 10,
	"marker": {
		"abc": "def"
	}
}`

	d, _ := json.MarshalIndent(accountLinesStruct, "", "\t")
	if string(d) != accountLinesJson {
		t.Error("json encoding does not match expected")
	}

	var conv AccountLinesRequest
	if err := json.Unmarshal(d, &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountLinesStruct) {
		t.Error("json decoding does not match expected struct")
	}
	if err := json.Unmarshal([]byte(accountLinesJson), &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountLinesStruct) {
		t.Error("json decoding does not match expected struct")
	}
}

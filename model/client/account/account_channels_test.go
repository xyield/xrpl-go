package account

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
)

func TestAccountChannelRequest(t *testing.T) {
	accountChannelStruct := AccountChannelsRequest{
		Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
		LedgerIndex:        VALIDATED,
	}

	accountChannelJson := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"destination_account": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
	"ledger_index": "validated"
}`
	d, _ := json.MarshalIndent(accountChannelStruct, "", "\t")
	if string(d) != accountChannelJson {
		t.Error("json encoding does not match expected")
	}

	var conv AccountChannelsRequest
	if err := json.Unmarshal(d, &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountChannelStruct) {
		t.Error("json decoding does not match expected struct")
	}
	if err := json.Unmarshal([]byte(accountChannelJson), &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountChannelStruct) {
		t.Error("json decoding does not match expected struct")
	}
}

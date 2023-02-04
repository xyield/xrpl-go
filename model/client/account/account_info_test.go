package account

import (
	"encoding/json"
	"reflect"
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
)

func TestAccountInfoRequest(t *testing.T) {
	accountInfoStruct := AccountInfoRequest{
		Account:     "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		LedgerIndex: CLOSED,
		Queue:       true,
		SignerList:  false,
		Strict:      true,
	}

	// SignerList assigned to default, omitted due to omitempty
	accountInfoJson := `{
	"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
	"ledger_index": "closed",
	"queue": true,
	"strict": true
}`
	d, _ := json.MarshalIndent(accountInfoStruct, "", "\t")
	if string(d) != accountInfoJson {
		t.Error("json encoding does not match expected")
	}

	var conv AccountInfoRequest
	if err := json.Unmarshal(d, &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountInfoStruct) {
		t.Error("json decoding does not match expected struct")
	}
	if err := json.Unmarshal([]byte(accountInfoJson), &conv); err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(conv, accountInfoStruct) {
		t.Error("json decoding does not match expected struct")
	}
}

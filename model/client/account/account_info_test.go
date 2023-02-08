package account

import (
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestAccountInfoRequest(t *testing.T) {
	s := AccountInfoRequest{
		Account:     "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
		LedgerIndex: CLOSED,
		Queue:       true,
		SignerList:  false,
		Strict:      true,
	}

	// SignerList assigned to default, omitted due to omitempty
	j := `{
	"account": "rG1QQv2nh2gr7RCZ1P8YYcBUKCCN633jCn",
	"ledger_index": "closed",
	"queue": true,
	"strict": true
}`
	if err := test.SerializeAndDeserialize(s, j); err != nil {
		t.Error(err)
	}
}

func TestAccountInfoResponse(t *testing.T) {

}

package account

import (
	"testing"

	. "github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestAccountChannelRequest(t *testing.T) {
	s := AccountChannelsRequest{
		Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
		LedgerIndex:        VALIDATED,
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"destination_account": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
	"ledger_index": "validated"
}`
	if err := test.SerializeAndDeserialize(s, j); err != nil {
		t.Error(err)
	}

}

func TestAccountChannelsResponse(t *testing.T) {

}

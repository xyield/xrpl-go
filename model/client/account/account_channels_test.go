package account

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xyield/xrpl-go/model/client/common"
	"github.com/xyield/xrpl-go/test"
)

func TestAccountChannelRequest(t *testing.T) {
	s := AccountChannelsRequest{
		Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
		LedgerIndex:        common.VALIDATED,
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"destination_account": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
	"ledger_index": "validated"
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}

}

func TestAccountChannelsResponse(t *testing.T) {
	s := AccountChannelsResponse{
		Account: "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
		Channels: []ChannelResult{
			{
				Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				Amount:             "100",
				Balance:            "200",
				ChannelID:          "500",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
			},
		},
		LedgerIndex: 123,
		LedgerHash:  "abc",
		Validated:   true,
		Limit:       1,
	}

	j := `{
	"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
	"channels": [
		{
			"account": "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
			"amount": "100",
			"balance": "200",
			"channel_id": "500",
			"destination_account": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3"
		}
	],
	"ledger_index": 123,
	"ledger_hash": "abc",
	"validated": true,
	"limit": 1
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestValidate(t *testing.T) {
	s := AccountChannelsRequest{
		Account: "",
	}

	err := s.Validate()

	require.EqualError(t, err, "no account ID specified")
}

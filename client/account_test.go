package client

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	rpcutils "github.com/xyield/xrpl-go/client/jsonrpc/utils"
	"github.com/xyield/xrpl-go/model/client/account"
	"github.com/xyield/xrpl-go/model/client/common"
)

func TestGetAccountChannels(t *testing.T) {

	tt := []struct {
		description string
		input       account.AccountChannelsRequest
		response    string
		output      account.AccountChannelsResponse
		expectedErr error
	}{
		{
			description: "validate failed on account type",
			input: account.AccountChannelsRequest{
				Account:            "",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
				LedgerIndex:        common.VALIDATED,
			},
			response:    "",
			output:      account.AccountChannelsResponse{},
			expectedErr: errors.New("no account ID specified"),
		},
		{
			description: "SendRequest returns error",
			input: account.AccountChannelsRequest{
				Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
				LedgerIndex:        common.VALIDATED,
			},
			response: `{
				"result": {
					"error": "ledgerIndexMalformed",
					"request": {
						"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
						"command": "account_info",
						"ledger_index": "-",
						"strict": true
					},
					"status": "error"
				}
			}`,
			output:      account.AccountChannelsResponse{},
			expectedErr: errors.New("ledgerIndexMalformed"),
		},
		{
			description: "GetResult returns an error",
			input: account.AccountChannelsRequest{
				Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
				LedgerIndex:        common.VALIDATED,
			},
			response: `{
				"result": {
					"account": 134,
					"ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
					"ledger_index": 71766343
				},
				"warning": "none",
				"warnings":
				[{
					"id": 1,
					"message": "message"
				}]
			}`,
			output:      account.AccountChannelsResponse{},
			expectedErr: errors.New("1 error(s) decoding:\n\n* 'account' expected type 'types.Address', got unconvertible type 'float64', value: '134'"),
		},
		{
			description: "successful response",
			input: account.AccountChannelsRequest{
				Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
				LedgerIndex:        common.VALIDATED,
			},
			response: `{
				"result": {
					"account": 134,
					"ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
					"ledger_index": 71766343
				},
				"warning": "none",
				"warnings":
				[{
					"id": 1,
					"message": "message"
				}]
			}`,
			output: account.AccountChannelsResponse{
				Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				LedgerIndex: 71766343,
				LedgerHash:  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			},
			expectedErr: errors.New("1 error(s) decoding:\n\n* 'account' expected type 'types.Address', got unconvertible type 'float64', value: '134'"),
		},
	}

	for _, tc := range tt {

		t.Run(tc.description, func(t *testing.T) {

			mc := &rpcutils.MockClient{}
			mc.DoFunc = rpcutils.MockResponse(tc.response, 200, mc)

			accounts := &accountImpl{
				Client: mc,
			}
			result, err := accounts.GetAccountChannels(&tc.input)

			if tc.expectedErr != nil {
				assert.EqualError(t, err, tc.expectedErr.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.output, result)
			}
		})
	}
}

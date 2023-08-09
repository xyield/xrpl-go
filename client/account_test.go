package client

// import (
// 	"testing"

// 	"github.com/mitchellh/mapstructure"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/require"
// 	"github.com/xyield/xrpl-go/model/client/account"
// 	"github.com/xyield/xrpl-go/model/client/common"
// )

// type mockClient struct {
// 	mock.Mock
// }

// type mockClientXrplResponse struct {
// 	Result map[string]any
// }

// func (m *mockClientXrplResponse) GetResult(v any) {
// 	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v})
// 	_ = dec.Decode(m.Result)
// }

// func (m *mockClient) SendRequest(req XRPLRequest) (XRPLResponse, error) {
// 	args := m.Called(req)
// 	return args.Get(0).(XRPLResponse), args.Error(1)
// }

// func TestGetAccountChannels(t *testing.T) {

// 	tt := []struct {
// 		description string
// 		input       account.AccountChannelsRequest
// 		response    string
// 		output      account.AccountChannelsResponse
// 		expectedErr error
// 	}{
// 		// {
// 		// 	description: "validate failed on account type",
// 		// 	input: account.AccountChannelsRequest{
// 		// 		Account:            "",
// 		// 		DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
// 		// 		LedgerIndex:        common.VALIDATED,
// 		// 	},
// 		// 	response:    "",
// 		// 	output:      account.AccountChannelsResponse{},
// 		// 	expectedErr: errors.New("no account ID specified"),
// 		// },
// 		// {
// 		// 	description: "SendRequest returns error",
// 		// 	input: account.AccountChannelsRequest{
// 		// 		Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
// 		// 		DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
// 		// 		LedgerIndex:        common.VALIDATED,
// 		// 	},
// 		// 	response: `{
// 		// 		"result": {
// 		// 			"error": "ledgerIndexMalformed",
// 		// 			"request": {
// 		// 				"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
// 		// 				"command": "account_info",
// 		// 				"ledger_index": "-",
// 		// 				"strict": true
// 		// 			},
// 		// 			"status": "error"
// 		// 		}
// 		// 	}`,
// 		// 	output:      account.AccountChannelsResponse{},
// 		// 	expectedErr: errors.New("ledgerIndexMalformed"),
// 		// },
// 		// {
// 		// 	description: "GetResult returns an error",
// 		// 	input: account.AccountChannelsRequest{
// 		// 		Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
// 		// 		DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
// 		// 		LedgerIndex:        common.VALIDATED,
// 		// 	},
// 		// 	response: `{
// 		// 		"result": {
// 		// 			"account": 134,
// 		// 			"ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
// 		// 			"ledger_index": 71766343
// 		// 		},
// 		// 		"warning": "none",
// 		// 		"warnings":
// 		// 		[{
// 		// 			"id": 1,
// 		// 			"message": "message"
// 		// 		}]
// 		// 	}`,
// 		// 	output:      account.AccountChannelsResponse{},
// 		// 	expectedErr: errors.New("1 error(s) decoding:\n\n* 'account' expected type 'types.Address', got unconvertible type 'float64', value: '134'"),
// 		// },
// 		{
// 			description: "successful response",
// 			input: account.AccountChannelsRequest{
// 				Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
// 				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
// 				LedgerIndex:        common.VALIDATED,
// 			},
// 			response: `{
// 				"result": {
// 					"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 					"ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
// 					"ledger_index": 71766343
// 				},
// 				"warning": "none",
// 				"warnings":
// 				[{
// 					"id": 1,
// 					"message": "message"
// 				}]
// 			}`,
// 			output: account.AccountChannelsResponse{
// 				Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 				LedgerIndex: 71766343,
// 				LedgerHash:  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
// 			},
// 			expectedErr: nil,
// 		},
// 	}

// 	for _, tc := range tt {

// 		t.Run(tc.description, func(t *testing.T) {

// 			cl := new(mockClient)
// 			a := &accountImpl{Client: cl}

// 			cl.On("SendRequest",
// 				&account.AccountChannelsRequest{Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 					DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
// 					LedgerIndex:        common.VALIDATED}).Return(&mockClientXrplResponse{
// 				Result: map[string]any{
// 					"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 					"channels": []any{
// 						map[string]any{
// 							"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 							"amount":              "1000",
// 							"balance":             "0",
// 							"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
// 							"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
// 							"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
// 							"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
// 							"settle_delay":        60,
// 						},
// 					},
// 					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
// 					"ledger_index": 71766314,
// 					"validated":    true,
// 				},
// 			}, nil)

// 			res, _, _ := a.GetAccountChannels(&account.AccountChannelsRequest{Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 				DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
// 				LedgerIndex:        common.VALIDATED})

// 			require.Equal(t, &account.AccountChannelsResponse{
// 				Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 				Channels: []account.ChannelResult{
// 					{
// 						Account:            "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
// 						Amount:             "1000",
// 						Balance:            "0",
// 						ChannelID:          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
// 						DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
// 						PublicKey:          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
// 						PublicKeyHex:       "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
// 						SettleDelay:        60,
// 					},
// 				},
// 				LedgerIndex: 71766314,
// 				LedgerHash:  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
// 				Validated:   true,
// 			}, res)

// 			// mc := &rpcutils.MockClient{}
// 			// mc.DoFunc = rpcutils.MockResponse(tc.response, 200, mc)

// 			// accounts := &accountImpl{
// 			// 	Client: mc,
// 			// }
// 			// result, xrplResponse, err := accounts.GetAccountChannels(&tc.input)

// 			// if tc.expectedErr != nil {
// 			// 	assert.EqualError(t, err, tc.expectedErr.Error())
// 			// } else {
// 			// 	assert.NoError(t, err)
// 			// 	assert.Equal(t, &tc.output, result)
// 			// 	assert.Equal(t, tc.xrplResponse, xrplResponse)
// 			// }
// 		})
// 	}
// }

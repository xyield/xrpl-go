package client

import (
	"errors"
	"testing"

	"github.com/mitchellh/mapstructure"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/xyield/xrpl-go/model/client/account"
	"github.com/xyield/xrpl-go/model/client/common"
)

type mockClient struct {
	mock.Mock
}

type mockClientXrplResponse struct {
	Result map[string]any
}

func (m *mockClientXrplResponse) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v})
	if err != nil {
		return err
	}
	_ = dec.Decode(m.Result)
	return nil
}

func (m *mockClient) SendRequest(req XRPLRequest) (XRPLResponse, error) {
	args := m.Called(req)
	return args.Get(0).(XRPLResponse), args.Error(1)
}

func TestGetAccountChannels(t *testing.T) {

	tt := []struct {
		description       string
		input             account.AccountChannelsRequest
		sendRequestResult mockClientXrplResponse
		output            account.AccountChannelsResponse
		expectedErr       error
	}{
		{
			description: "validate failed on account type",
			input: account.AccountChannelsRequest{
				Account:            "",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
				LedgerIndex:        common.VALIDATED,
			},
			output:      account.AccountChannelsResponse{},
			expectedErr: errors.New("no account ID specified"),
		},
		{
			description: "GetResult returns an error",
			input: account.AccountChannelsRequest{
				Account:            "rLHmBn4fT92w4F6ViyYbjoizLTo83tHTHu",
				DestinationAccount: "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
			},
			sendRequestResult: mockClientXrplResponse{
				Result: map[string]any{
					"account":             123,
					"destination_account": "rnZvsWuLem5Ha46AZs61jLWR9R5esinkG3",
				},
			},
			output:      account.AccountChannelsResponse{},
			expectedErr: errors.New("1 error(s) decoding:\n\n* 'account' expected type 'types.Address', got unconvertible type 'float64', value: '134'"),
		},
		{
			description: "successful response",
			input: account.AccountChannelsRequest{
				Account:            "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				LedgerIndex:        common.VALIDATED,
			},
			sendRequestResult: mockClientXrplResponse{
				Result: map[string]any{
					"account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
					"channels": []any{
						map[string]any{
							"account":             "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
							"amount":              "1000",
							"balance":             "0",
							"channel_id":          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
							"destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
							"public_key":          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
							"public_key_hex":      "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
							"settle_delay":        60,
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": 71766314,
					"validated":    true,
				},
			},
			output: account.AccountChannelsResponse{
				Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				LedgerIndex: 71766314,
				LedgerHash:  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
				Channels: []account.ChannelResult{
					{
						Account:            "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
						Amount:             "1000",
						Balance:            "0",
						ChannelID:          "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
						DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
						PublicKey:          "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
						PublicKeyHex:       "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
						SettleDelay:        60,
					},
				},
				Validated: true,
			},
			expectedErr: nil,
		},
	}

	for _, tc := range tt {

		t.Run(tc.description, func(t *testing.T) {

			cl := new(mockClient)
			a := &accountImpl{Client: cl}

			cl.On("SendRequest", &tc.input).Return(&tc.sendRequestResult, nil)

			res, _, err := a.GetAccountChannels(&tc.input)

			if err != nil {
				require.Equal(t, tc.expectedErr, err)
			} else {
				require.Equal(t, &tc.output, res)
			}

		})
	}
}

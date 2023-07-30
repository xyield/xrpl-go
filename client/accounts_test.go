package client

import (
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

func (m *mockClientXrplResponse) GetResult(v any) {
	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v})
	_ = dec.Decode(m.Result)
}

func (m *mockClientXrplResponse) CheckError() error {
	return nil
}

func (m *mockClient) SendRequest(req common.XRPLRequest) (XRPLResponse, error) {
	args := m.Called(req)
	return args.Get(0).(XRPLResponse), args.Error(1)
}

func TestGetAccountChannels(t *testing.T) {
	cl := new(mockClient)
	a := &accountImpl{client: cl}

	cl.On("SendRequest",
		&account.AccountChannelsRequest{Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
			LedgerIndex:        common.VALIDATED}).Return(&mockClientXrplResponse{
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
	}, nil)

	res, _ := a.GetAccountChannels(&account.AccountChannelsRequest{Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
		DestinationAccount: "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
		LedgerIndex:        common.VALIDATED})
	require.Equal(t, &account.AccountChannelsResponse{
		Account: "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
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
		LedgerIndex: 71766314,
		LedgerHash:  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
		Validated:   true,
	}, res)
}

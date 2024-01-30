package websocket

import (
	"encoding/json"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
	"github.com/CreatureDev/xrpl-go/client"
	"github.com/CreatureDev/xrpl-go/model/client/account"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestSendRequest(t *testing.T) {
	tt := []struct {
		description    string
		req            client.XRPLRequest
		res            client.XRPLResponse
		expectedErr    error
		serverMessages []map[string]any
	}{
		{
			description: "successful request",
			req: &account.AccountChannelsRequest{
				Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			},
			res: &WebSocketClientXrplResponse{
				ID: 1,
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
							"settle_delay":        json.Number("60"),
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": json.Number("71766314"),
					"validated":    true,
				},
			},
			expectedErr: nil,
			serverMessages: []map[string]any{
				{
					"id": 1,
					"result": map[string]any{
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
			},
		},
		{
			description: "Invalid ID",
			req: &account.AccountChannelsRequest{
				Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			},
			res: &WebSocketClientXrplResponse{
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
							"settle_delay":        json.Number("60"),
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": json.Number("71766314"),
					"validated":    true,
				},
			},
			expectedErr: ErrIncorrectId,
			serverMessages: []map[string]any{
				{
					"id": 2,
					"result": map[string]any{
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
			},
		},
		{
			description: "Error response",
			req: &account.AccountChannelsRequest{
				Account: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
			},
			res: &WebSocketClientXrplResponse{
				ID: 1,
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
							"settle_delay":        json.Number("60"),
						},
					},
					"ledger_hash":  "1EDBBA3C793863366DF5B31C2174B6B5E6DF6DB89A7212B86838489148E2A581",
					"ledger_index": json.Number("71766314"),
					"validated":    true,
				},
			},
			expectedErr: &ErrorWebsocketClientXrplResponse{
				Type: "invalidParams",
				Request: map[string]any{
					"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				},
			},
			serverMessages: []map[string]any{
				{
					"id":    1,
					"error": "invalidParams",
					"value": map[string]any{
						"account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					},
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.description, func(t *testing.T) {
			ws := &test.MockWebSocketServer{Msgs: tc.serverMessages}
			s := ws.TestWebSocketServer(func(c *websocket.Conn) {
				for _, m := range tc.serverMessages {
					err := c.WriteJSON(m)
					if err != nil {
						println("error writing message")
					}
				}
			})
			defer s.Close()
			url, _ := test.ConvertHttpToWS(s.URL)
			cl := &WebsocketClient{cfg: &WebsocketConfig{URL: url}}

			res, err := cl.SendRequest(tc.req)

			if tc.expectedErr != nil {
				require.EqualError(t, err, tc.expectedErr.Error())
			} else {
				require.NoError(t, err)
				require.EqualValues(t, tc.res, res)
			}
		})
	}
}

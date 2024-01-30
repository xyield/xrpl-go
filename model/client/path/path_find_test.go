package path

import (
	"testing"

	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
	"github.com/CreatureDev/xrpl-go/test"
)

func TestPathFindRequest(t *testing.T) {
	s := PathFindRequest{
		Subcommand:         CREATE,
		SourceAccount:      "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		DestinationAmount: types.IssuedCurrencyAmount{
			Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			Currency: "USD",
			Value:    "0.001",
		},
	}

	j := `{
	"subcommand": "create",
	"source_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"destination_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"destination_amount": {
		"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		"currency": "USD",
		"value": "0.001"
	}
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestPathFindResponse(t *testing.T) {
	s := PathFindResponse{
		Alternatives: []PathAlternative{
			{
				PathsComputed: [][]transactions.PathStep{
					{
						{
							Currency: "USD",
							Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
					{
						{
							Currency: "USD",
							Issuer:   "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
						{
							Account: "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
					},
					{
						{
							Currency: "USD",
							Issuer:   "r9vbV3EHvXWjSkeQ6CAcYVPGeq7TuiXY2X",
						},
						{
							Account: "r9vbV3EHvXWjSkeQ6CAcYVPGeq7TuiXY2X",
						},
					},
					{
						{
							Currency: "USD",
							Issuer:   "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun",
						},
						{
							Account: "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun",
						},
					},
				},
				SourceAmount: types.XRPCurrencyAmount(251686),
			},
			{
				PathsComputed: [][]transactions.PathStep{
					{
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Currency: "USD",
							Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
					{
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Currency: "USD",
							Issuer:   "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
						{
							Account: "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
					},
					{
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Currency: "USD",
							Issuer:   "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun",
						},
						{
							Account: "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun",
						},
					},
				},
				SourceAmount: types.IssuedCurrencyAmount{
					Issuer:   "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					Currency: "BTC",
					Value:    "0.000001541291269274307",
				},
			},
			{
				PathsComputed: [][]transactions.PathStep{
					{
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Currency: "USD",
							Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
				},
				SourceAmount: types.IssuedCurrencyAmount{
					Currency: "CHF",
					Issuer:   "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					Value:    "0.0009211546262510451",
				},
			},
			{
				PathsComputed: [][]transactions.PathStep{
					{
						{
							Account: "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA",
						},
						{
							Currency: "USD",
							Issuer:   "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
						{
							Account: "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q",
						},
					},
					{
						{
							Account: "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA",
						},
						{
							Currency: "USD",
							Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
				},
				SourceAmount: types.IssuedCurrencyAmount{
					Currency: "CNY",
					Issuer:   "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
					Value:    "0.006293562",
				},
			},
		},

		DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		DestinationAmount: types.IssuedCurrencyAmount{
			Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			Currency: "USD",
			Value:    "0.001",
		},
		SourceAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		FullReply:     false,
	}
	j := `{
	"alternatives": [
		{
			"paths_computed": [
				[
					{
						"currency": "USD",
						"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				],
				[
					{
						"currency": "USD",
						"issuer": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
					},
					{
						"account": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
					}
				],
				[
					{
						"currency": "USD",
						"issuer": "r9vbV3EHvXWjSkeQ6CAcYVPGeq7TuiXY2X"
					},
					{
						"account": "r9vbV3EHvXWjSkeQ6CAcYVPGeq7TuiXY2X"
					}
				],
				[
					{
						"currency": "USD",
						"issuer": "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun"
					},
					{
						"account": "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun"
					}
				]
			],
			"source_amount": "251686"
		},
		{
			"paths_computed": [
				[
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"currency": "USD",
						"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				],
				[
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"currency": "USD",
						"issuer": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
					},
					{
						"account": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
					}
				],
				[
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"currency": "USD",
						"issuer": "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun"
					},
					{
						"account": "rLEsXccBGNR3UPuPu2hUXPjziKC3qKSBun"
					}
				]
			],
			"source_amount": {
				"issuer": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"currency": "BTC",
				"value": "0.000001541291269274307"
			}
		},
		{
			"paths_computed": [
				[
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"currency": "USD",
						"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				]
			],
			"source_amount": {
				"issuer": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"currency": "CHF",
				"value": "0.0009211546262510451"
			}
		},
		{
			"paths_computed": [
				[
					{
						"account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"
					},
					{
						"currency": "USD",
						"issuer": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
					},
					{
						"account": "rMwjYedjc7qqtKYVLiAccJSmCwih4LnE2q"
					}
				],
				[
					{
						"account": "razqQKzJRdB4UxFPWf5NEpEG3WMkmwgcXA"
					},
					{
						"currency": "USD",
						"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				]
			],
			"source_amount": {
				"issuer": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
				"currency": "CNY",
				"value": "0.006293562"
			}
		}
	],
	"destination_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"destination_amount": {
		"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		"currency": "USD",
		"value": "0.001"
	},
	"source_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"full_reply": false
}`

	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

package path

import (
	"testing"

	"github.com/xyield/xrpl-go/model/transactions"
	"github.com/xyield/xrpl-go/model/transactions/types"
	"github.com/xyield/xrpl-go/test"
)

func TestRipplePathFindRequest(t *testing.T) {
	s := RipplePathFindRequest{
		SourceAccount:      "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		DestinationAmount: types.IssuedCurrencyAmount{
			Issuer:   "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
			Currency: "USD",
			Value:    "0.001",
		},
		SourceCurrencies: []types.IssuedCurrencyAmount{
			{
				Currency: "XRP",
			},
			{
				Currency: "USD",
			},
		},
	}

	j := `{
	"source_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"destination_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"destination_amount": {
		"issuer": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
		"currency": "USD",
		"value": "0.001"
	},
	"source_currencies": [
		{
			"currency": "XRP"
		},
		{
			"currency": "USD"
		}
	]
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

func TestRipplePathFindResponse(t *testing.T) {
	s := RipplePathFindResponse{
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
							Issuer:   "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1",
						},
						{
							Account: "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
					{
						{
							Currency: "USD",
							Issuer:   "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1",
						},
						{
							Account: "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1",
						},
						{
							Account: "rLpq4LgabRfm1xEX5dpWfJovYBH6g7z99q",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
					{
						{
							Currency: "USD",
							Issuer:   "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1",
						},
						{
							Account: "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1",
						},
						{
							Account: "rPuBoajMjFoDjweJBrtZEBwUMkyruxpwwV",
						},
						{
							Account: "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B",
						},
					},
				},
				SourceAmount: types.XRPCurrencyAmount(256987),
			},
		},
		DestinationAccount: "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
		DestinationCurrencies: []string{
			"015841551A748AD2C1F76FF6ECB0CCCD00000000",
			"JOE",
			"DYM",
			"EUR",
			"CNY",
			"MXN",
			"BTC",
			"USD",
			"XRP",
		},
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
						"issuer": "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
					},
					{
						"account": "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				],
				[
					{
						"currency": "USD",
						"issuer": "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
					},
					{
						"account": "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
					},
					{
						"account": "rLpq4LgabRfm1xEX5dpWfJovYBH6g7z99q"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				],
				[
					{
						"currency": "USD",
						"issuer": "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
					},
					{
						"account": "rrpNnNLKrartuEqfJGpqyDwPj1AFPg9vn1"
					},
					{
						"account": "rPuBoajMjFoDjweJBrtZEBwUMkyruxpwwV"
					},
					{
						"account": "rvYAfWj5gh67oV6fW32ZzP3Aw4Eubs59B"
					}
				]
			],
			"source_amount": "256987"
		}
	],
	"destination_account": "r9cZA1mLK5R5Am25ArfXFmqgNwjZgnfk59",
	"destination_currencies": [
		"015841551A748AD2C1F76FF6ECB0CCCD00000000",
		"JOE",
		"DYM",
		"EUR",
		"CNY",
		"MXN",
		"BTC",
		"USD",
		"XRP"
	]
}`
	if err := test.SerializeAndDeserialize(t, s, j); err != nil {
		t.Error(err)
	}
}

package jsonrpc

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

// type testParamStruct struct {
// 	Account string `json:"account"`
// 	Info    string `json:"info"`
// }

func TestCheckForError(t *testing.T) {

	t.Run("Error Response", func(t *testing.T) {

		json := `{
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
		}`

		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 200, // error response still returns a 200
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)

		assert.Nil(t, bodyBytes)

		expErrpr := &JsonRpcClientError{ErrorString: "ledgerIndexMalformed"}

		assert.Equal(t, expErrpr, err)
	})

	t.Run("Error Response with error code", func(t *testing.T) {

		json := "Null Method" // is this right? https://xrpl.org/error-formatting.html#universal-errors

		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 400,
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)

		assert.Nil(t, bodyBytes)

		expErrpr := &JsonRpcClientError{ErrorString: "Null Method"}

		assert.Equal(t, expErrpr, err)
	})

	t.Run("No error Response", func(t *testing.T) {

		json := `{
			"result": {
			  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			  "channels": [
				{
				  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				  "amount": "1000",
				  "balance": "0",
				  "channel_id": "C7F634794B79DB40E87179A9D1BF05D05797AE7E92DF8E93FD6656E8C4BE3AE7",
				  "destination_account": "ra5nK24KXen9AHvsdFTKHSANinZseWnPcX",
				  "public_key": "aBR7mdD75Ycs8DRhMgQ4EMUEmBArF8SEh1hfjrT2V9DQTLNbJVqw",
				  "public_key_hex": "03CFD18E689434F032A4E84C63E2A3A6472D684EAF4FD52CA67742F3E24BAE81B2",
				  "settle_delay": 60
				}
			  ],
			  "ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			  "ledger_index": 71766343,
			  "status": "success",
			  "validated": true
			}
		  }`

		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 200,
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)

		assert.Nil(t, err)
		assert.NotNil(t, bodyBytes)
	})
}

func TestCreateRequest(t *testing.T) {

	// t.Run("Create request", func(t *testing.T) {

	// 	params := testParamStruct{
	// 		Account: "account1",
	// 		Info:    "some account information",
	// 	}
	// 	expectedRequestBody := `{
	// 		"method": "account_channels",
	// 		"params": [
	// 			{
	// 				"account": "account1",
	// 				"info": "some account information"
	// 			}
	// 		]
	// 	}`

	// 	byteRequest, err := CreateRequest("account_channels", params)

	// 	assert.NoError(t, err)

	// 	a := AnyJson{}
	// 	_ = jsoniter.Unmarshal(byteRequest, &a)

	// 	// TODO: check byte array is same? check made into same struct? check json same?
	// 	assert.Equal(t, a, expectedRequestBody)
	// })

}

func TestSendRequest(t *testing.T) {

	// TODO: mock client

	// send correct headers and url
	// test the defer
	// mock out error response - return error
	// mock out success response - return the response
	// succesfully marshals into response struct passed in
}

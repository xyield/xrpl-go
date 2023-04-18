package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	jsonrpcclient "github.com/xyield/xrpl-go/xrpl-jsonrpc-client"
	rpcutils "github.com/xyield/xrpl-go/xrpl-jsonrpc-client/rpc-utils"
)

type testParamStruct struct {
	Account string `json:"account"`
	Info    string `json:"info"`
}

type testResponseStruct struct {
	Account     string `json:"account"`
	LedgerHash  string `json:"ledger_hash,omitempty"`
	LedgerIndex int    `json:"ledger_index,omitempty"`
}

// This method will be added to every response struct passed into SendRequest
func (r *testResponseStruct) UnmarshallJSON(data []byte) error {
	type Alias testResponseStruct
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = testResponseStruct(aux)
	return nil
}

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
		assert.NotNil(t, bodyBytes)
		expError := &JsonRpcClientError{ErrorString: "ledgerIndexMalformed"}
		assert.Equal(t, expError, err)
	})

	t.Run("Error Response with error code", func(t *testing.T) {

		json := "Null Method" // https://xrpl.org/error-formatting.html#universal-errors

		b := ioutil.NopCloser(bytes.NewReader([]byte(json)))
		res := &http.Response{
			StatusCode: 400,
			Body:       b,
		}

		bodyBytes, err := CheckForError(res)
		assert.NotNil(t, bodyBytes)
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

	t.Run("Create request", func(t *testing.T) {

		params := testParamStruct{
			Account: "account1",
			Info:    "some account information",
		}

		expetedBody := jsonRpcRequest{
			Method: "account_channels",
			Params: [1]interface{}{params},
		}
		expectedRequestBytes, _ := jsoniter.Marshal(expetedBody)

		byteRequest, err := CreateRequest("account_channels", params)

		assert.NoError(t, err)
		// assert bytes equal
		assert.Equal(t, expectedRequestBytes, byteRequest)
		// assert json equal
		assert.Equal(t, string(expectedRequestBytes), string(byteRequest))
	})
	t.Run("Create request - no parameters", func(t *testing.T) {

		expetedBody := jsonRpcRequest{
			Method: "account_channels",
		}
		expectedRequestBytes, _ := jsoniter.Marshal(expetedBody)

		byteRequest, err := CreateRequest("account_channels", nil)

		assert.NoError(t, err)
		// assert bytes equal
		assert.Equal(t, expectedRequestBytes, byteRequest)
		// assert json equal
		assert.Equal(t, string(expectedRequestBytes), string(byteRequest))
	})
}

func TestSendRequest(t *testing.T) {
	t.Run("Send request - sucessful response", func(t *testing.T) {

		response := `{
			"result": {
			  "account": "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			  "ledger_hash": "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			  "ledger_index": 71766343
				}
			}`

		mc := &rpcutils.MockClient{}
		mc.DoFunc = rpcutils.MockResponse(response, 200, mc)

		requestBodyBytes := createRequest()

		resStruct := &testResponseStruct{}

		cfg, err := jsonrpcclient.NewConfig("http://testnode/")
		assert.NoError(t, err)
		cfg.AddHttpClient(mc)

		err = SendRequest(requestBodyBytes, cfg, resStruct)

		assert.NoError(t, err)
		assert.Equal(t, "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn", resStruct.Account)
		assert.Equal(t, "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D", resStruct.LedgerHash)
		assert.Equal(t, 71766343, resStruct.LedgerIndex)

	})

	t.Run("Send request - error response", func(t *testing.T) {

		response := `{
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

		mc := &rpcutils.MockClient{}
		mc.DoFunc = rpcutils.MockResponse(response, 200, mc)

		requestBodyBytes := createRequest()

		resStruct := &testResponseStruct{}

		cfg, err := jsonrpcclient.NewConfig("http://testnode/")
		assert.NoError(t, err)
		cfg.AddHttpClient(mc)

		err = SendRequest(requestBodyBytes, cfg, resStruct)

		assert.Equal(t, "", resStruct.Account)
		assert.Equal(t, "", resStruct.LedgerHash)
		assert.EqualError(t, err, "ledgerIndexMalformed")
	})

	// test different struct also works?
	// test it sends correct headers and url
	// test the defer
}

func createRequest() []byte {
	params := testParamStruct{
		Account: "account1",
		Info:    "some account information",
	}
	requestBody := jsonRpcRequest{
		Method: "account_channels",
		Params: [1]interface{}{params},
	}
	requestBodyBytes, _ := jsoniter.Marshal(requestBody)

	return requestBodyBytes
}

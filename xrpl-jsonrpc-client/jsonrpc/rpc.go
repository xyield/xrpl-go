package jsonrpc

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	jsoniter "github.com/json-iterator/go"
	jsonrpcclient "github.com/xyield/xrpl-go/xrpl-jsonrpc-client"
)

type AnyJson map[string]interface{}

type JsonRpcClientError struct {
	ErrorString string
}

func (e *JsonRpcClientError) Error() string {
	return "No root path provided"
}

type jsonRpcRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params"`
}

type jsonRpcResponse struct {
	Result    AnyJson       `json:"result"`
	Warning   string        `json:"warning,omitempty"`
	Warnings  []interface{} `json:"warnings,omitempty"`
	Forwarded bool          `json:"forwarded,omitempty"`
}

// each method will have a successful response struct that will impl this - better way to do this than empty interface?
type RequestResult interface {
}

// each method has their own param struct to be passed into CreateRequest
// each param struct (for each request) will impl this interface - add method to impl?
type RequestParams interface {
}

func CreateRequest(method string, params RequestParams) ([]byte, error) {

	// TODO: Have to serilase the parameters for signed transactions - do this before this func?

	r := &jsonRpcRequest{
		Method: method,
		Params: [1]interface{}{params}, // array with 1 json array
	}

	// jsonBytes, err := json.Marshal(c) // TODO: add correct formatting for the marshal here? with correct spaces etc
	jsonBytes, err := json.MarshalIndent(r, "", "\t")

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

func CheckForError(res *http.Response) ([]byte, error) {

	b, err := ioutil.ReadAll(res.Body)
	if err != nil || b == nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, &JsonRpcClientError{ErrorString: string(b)}
	}

	var jr jsonRpcResponse

	err = jsoniter.Unmarshal(b, &jr)
	if err != nil {
		return nil, err
	}

	// result will have 'error' if error
	if _, ok := jr.Result["error"]; ok {
		return nil, &JsonRpcClientError{ErrorString: jr.Result["error"].(string)}
	}

	return b, nil
}

// pass config in or give this method a new interface reciever for each method
func SendRequest(body []byte, cfg *jsonrpcclient.Config, r RequestResult) error {

	// TODO: add CreateRequest in here too?

	req, err := http.NewRequest(http.MethodPost, cfg.Url, bytes.NewReader(body))
	if err != nil {
		return err
	}
	headers := map[string][]string{
		"Content-Type": {"application/json"},
	}
	req.Header = headers

	response, err := cfg.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	// allow client to reuse persistant connection
	defer response.Body.Close()

	// TODO: check if get 503 service unavailable (from rate limiting) and re-try if so

	responseBody, err := CheckForError(response)
	if err != nil {
		return err
	}

	// unmarshall successful response into expected struct
	err = jsoniter.Unmarshal(responseBody, &r)
	if err != nil {
		return err
	}

	return nil
}

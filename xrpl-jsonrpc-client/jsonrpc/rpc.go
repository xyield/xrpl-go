package jsonrpc

import (
	"bytes"
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
	return e.ErrorString
}

type jsonRpcRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params"`
}

type jsonRpcResponse struct {
	Result    AnyJson       `json:"result"`
	Warning   string        `json:"warning,omitempty"`
	Warnings  []interface{} `json:"warnings,omitempty"` // TODO: check full response body maps to this correctly
	Forwarded bool          `json:"forwarded,omitempty"`
}

// each method has their own request param struct to be passed into CreateRequest
type RequestParams interface {
}

// each method will have a successful response struct that will impl this
type ResponseType interface {
	UnmarshallJSON([]byte) error
}

// Params will have been serialised and added to request struct before passing to this method
// Have to serilase the parameters for signed transactions etc
func CreateRequest(method string, params RequestParams) ([]byte, error) {

	var body jsonRpcRequest

	if params == nil {
		body = jsonRpcRequest{
			Method: method,
		}
	} else {
		body = jsonRpcRequest{
			Method: method,
			Params: [1]interface{}{params}, // array with 1 json array - will be its own struct with json serialising tags
		}
	}

	jsonBytes, err := jsoniter.Marshal(body)
	// TODO: add correct formatting for the marshal here - do we need the indent formatter below?
	// jsonBytes, err := json.MarshalIndent(r, "", "\t")

	if err != nil {
		return nil, err
	}

	return jsonBytes, nil
}

// currently returning the whole response as well as the error - ok?
func CheckForError(res *http.Response) (jsonRpcResponse, error) {

	var jr jsonRpcResponse

	b, err := ioutil.ReadAll(res.Body)
	if err != nil || b == nil {
		return jr, err
	}

	if res.StatusCode != 200 {
		return jr, &JsonRpcClientError{ErrorString: string(b)}
	}

	err = jsoniter.Unmarshal(b, &jr)
	if err != nil {
		return jr, err
	}

	// result will have 'error' if error response
	if _, ok := jr.Result["error"]; ok {
		return jr, &JsonRpcClientError{ErrorString: jr.Result["error"].(string)}
	}

	return jr, nil
}

// TODO: pass config in or give this method a new interface reciever for each method interface?
func SendRequest(body []byte, cfg *jsonrpcclient.Config, responseStruct ResponseType) error {

	// add CreateRequest in here so only call 1 method?

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

	jsonRpcResponse, err := CheckForError(response)
	if err != nil {
		return err
	}

	// If no error unmarshall jsonRpcResponse.Result into the result struct
	b, err := jsoniter.Marshal(jsonRpcResponse.Result)
	if err != nil {
		return err
	}

	err = responseStruct.UnmarshallJSON(b)
	if err != nil {
		return err
	}

	// TODO: we are going to have to run the decode method here on the response body?

	return nil
}

/////////////////////// Doesn't work marshalling result into a string
// b, _ := jsoniter.Marshal(jsonRpcResponse.Result)
// fmt.Println(string(b))

// r, _ := GetResultString(response)

// resultBytes, err := jsoniter.Marshal(r.Result) // marshall the string version of result
// if err != nil {
// 	return &responseStruct, err
// }

// var jr jsonRpcResult

// b, err := ioutil.ReadAll(response.Body)
// if err != nil || b == nil {
// 	return &responseStruct, err
// }
// err = jsoniter.Unmarshal(b, &jr) // grab just stringified value of "result" so do not convert to AnyJson
// if err != nil {
// 	return &responseStruct, err
// }
///////////////

// err = jsoniter.Unmarshal(byteJson, &responseStruct) // error as can only unmarshall into a pointer
// fmt.Println("Operation: ", responseStruct)
// if err != nil {
// 	// return &responseStruct, err
// 	return err
// }

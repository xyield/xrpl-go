package jsonrpc

// import (
// 	"bytes"
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"net/http"
// 	"time"

// 	jsoniter "github.com/json-iterator/go"
// 	jsonrpcclient "github.com/xyield/xrpl-go/xrpl-jsonrpc-client"
// )

// type AnyJson map[string]interface{}

// type JsonRpcClientError struct {
// 	ErrorString string
// }

// func (e *JsonRpcClientError) Error() string {
// 	return e.ErrorString
// }

// type jsonRpcRequest struct {
// 	Method string         `json:"method"`
// 	Params [1]interface{} `json:"params"`
// }

// type ApiWarning struct {
// 	Id      int         `json:"id"`
// 	Message string      `json:"message"`
// 	Details interface{} `json:"details,omitempty"`
// }

// type jsonRpcResponse struct {
// 	Result    AnyJson      `json:"result"`
// 	Warning   string       `json:"warning,omitempty"`
// 	Warnings  []ApiWarning `json:"warnings,omitempty"`
// 	Forwarded bool         `json:"forwarded,omitempty"`
// }

// // each method has their own request param struct to be passed into CreateRequest
// type RequestParams interface {
// }

// // each method will have a successful response struct that will impl this
// type ResponseType interface {
// 	UnmarshallJSON([]byte) error
// }

// // CreateRequest formats the parameters and method name ready for sending request
// // Params will have been serialised if required and added to request struct before being passed to this method
// func CreateRequest(method string, params RequestParams) ([]byte, error) {

// 	var body jsonRpcRequest

// 	if params == nil {
// 		body = jsonRpcRequest{
// 			Method: method,
// 		}
// 	} else {
// 		body = jsonRpcRequest{
// 			Method: method,
// 			// each param object will have a struct with json serialising tags
// 			Params: [1]interface{}{params},
// 		}
// 	}

// 	jsonBytes, err := jsoniter.Marshal(body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to marshal JSON-RPC request for method %s with parameters %+v: %w", method, params, err)
// 	}

// 	return jsonBytes, nil
// }

// // CheckForError reads the http response and formats the error if it exists
// func CheckForError(res *http.Response) (jsonRpcResponse, error) {

// 	var jr jsonRpcResponse

// 	b, err := ioutil.ReadAll(res.Body)
// 	if err != nil || b == nil {
// 		return jr, err
// 	}

// 	// In case a different error code is returned
// 	if res.StatusCode != 200 {
// 		return jr, &JsonRpcClientError{ErrorString: string(b)}
// 	}

// 	err = jsoniter.Unmarshal(b, &jr)
// 	if err != nil {
// 		return jr, err
// 	}

// 	// result will have 'error' if error response
// 	if _, ok := jr.Result["error"]; ok {
// 		return jr, &JsonRpcClientError{ErrorString: jr.Result["error"].(string)}
// 	}

// 	return jr, nil
// }

// // SendRequest makes request to xrpl and returns error if exists or populates the response struct for each request if sucessful
// // TODO: pass in an object now with the method and params in, not as seperates
// func SendRequest(method string, params RequestParams, cfg *jsonrpcclient.Config, responseStruct ResponseType) (jsonRpcResponse, error) {

// 	var jr jsonRpcResponse

// 	body, err := CreateRequest(method, params)
// 	if err != nil {
// 		return jr, err
// 	}

// 	req, err := http.NewRequest(http.MethodPost, cfg.Url, bytes.NewReader(body))
// 	if err != nil {
// 		return jr, err
// 	}

// 	// add timeout context to prevent hanging
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
// 	defer cancel()
// 	req = req.WithContext(ctx)

// 	req.Header = cfg.Headers

// 	var response *http.Response

// 	response, err = cfg.HTTPClient.Do(req)
// 	if err != nil || response == nil {
// 		return jr, err
// 	}

// 	// allow client to reuse persistant connection
// 	defer response.Body.Close()

// 	// Check for service unavailable response and retry if so
// 	if response.StatusCode == 503 {

// 		maxRetries := 3
// 		backoffDuration := 1 * time.Second

// 		for i := 0; i < maxRetries; i++ {
// 			time.Sleep(backoffDuration)

// 			// Make request again after waiting
// 			response, err = cfg.HTTPClient.Do(req)
// 			if err != nil {
// 				return jr, err
// 			}

// 			if response.StatusCode != 503 {
// 				break
// 			}

// 			// Increase backoff duration for the next retry
// 			backoffDuration *= 2
// 		}

// 		if response.StatusCode == 503 {
// 			// Return service unavailable error here after retry 3 times
// 			return jr, &JsonRpcClientError{ErrorString: "Server is overloaded, rate limit exceeded"}
// 		}

// 	}

// 	jr, err = CheckForError(response)
// 	if err != nil {
// 		return jr, err
// 	}

// 	// If no error unmarshall jsonRpcResponse.Result into the result struct
// 	b, err := jsoniter.Marshal(jr.Result)
// 	if err != nil {
// 		return jr, err
// 	}

// 	err = responseStruct.UnmarshallJSON(b)
// 	if err != nil {
// 		return jr, err
// 	}

// 	return jr, err
// }

package jsonrpcclient

import (
	"bytes"
	"io"
	"net/http"
)

type mockClient struct {
	DoFunc       func(req *http.Request) (*http.Response, error)
	Spy          *http.Request
	RequestCount int
}

func (m *mockClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	// just in case you want default correct return value
	return &http.Response{}, nil
}

func mockResponse(resString string, statusCode int, m *mockClient) func(req *http.Request) (*http.Response, error) {
	return func(req *http.Request) (*http.Response, error) {
		m.Spy = req
		return &http.Response{
			StatusCode: statusCode,
			Body:       io.NopCloser(bytes.NewReader([]byte(resString))),
		}, nil
	}
}

// // Method to be able to set mock client as a Client object in each interface
// func (m *mockClient) sendRequest(reqParams common.XRPLRequest) (common.XRPLResponse, error) {
// 	req := http.Request{}
// 	res, err := m.Do(&req)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// map res to XRPLResponse and return this
// 	b, err := ioutil.ReadAll(res.Body)
// 	if err != nil || b == nil {
// 		return nil, err
// 	}

// 	var jr jsonrpcmodels.JsonRpcResponse
// 	err = jsoniter.Unmarshal(b, &jr)
// 	if err != nil {
// 		return jr, err
// 	}

// 	if _, ok := jr.Result["error"]; ok {
// 		return jr, errors.New(jr.Result["error"].(string))
// 	}
// 	return jr, nil
// }

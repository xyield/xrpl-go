package jsonrpcmodels

type JsonRpcRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params,omitempty"`
}

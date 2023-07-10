package jsonrpcclient

type AnyJson map[string]interface{}

type ApiWarning struct {
	Id      int         `json:"id"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

type jsonRpcResponse struct {
	Result    AnyJson      `json:"result"`
	Warning   string       `json:"warning,omitempty"`
	Warnings  []ApiWarning `json:"warnings,omitempty"`
	Forwarded bool         `json:"forwarded,omitempty"`
}

type jsonRpcRequest struct {
	Method string         `json:"method"`
	Params [1]interface{} `json:"params,omitempty"`
	// TODO: change back to just 1 size?
	// Params []interface{} `json:"params,omitempty"`
}

type JsonRpcClientError struct {
	ErrorString string
}

func (e *JsonRpcClientError) Error() string {
	return e.ErrorString
}

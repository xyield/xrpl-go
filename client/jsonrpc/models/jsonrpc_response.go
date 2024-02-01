package jsonrpcmodels

import (
	"encoding/json"
	"fmt"

	"github.com/CreatureDev/xrpl-go/client"
)

type JsonRpcResponse struct {
	Result    json.RawMessage              `json:"result"`
	Warning   string                       `json:"warning,omitempty"`
	Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                         `json:"forwarded,omitempty"`
}

type ApiWarning struct {
	Id      int         `json:"id"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (r JsonRpcResponse) GetResult(v any) error {
	if len(r.Result) == 0 {
		return nil
	}
	err := json.Unmarshal(r.Result, v)
	if err != nil {
		return err
	}
	return nil
}

func (r JsonRpcResponse) GetError() error {
	if len(r.Result) == 0 {
		return nil
	}
	type reqError struct {
		Error string `json:"error"`
	}
	var errResponse reqError
	err := json.Unmarshal(r.Result, &errResponse)
	if err != nil {
		return err
	}
	if errResponse.Error != "" {
		return fmt.Errorf(errResponse.Error)
	}
	return nil
}

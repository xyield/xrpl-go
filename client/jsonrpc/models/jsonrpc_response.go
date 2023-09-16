package jsonrpcmodels

import (
	"github.com/mitchellh/mapstructure"
	"github.com/xyield/xrpl-go/client"
)

type JsonRpcResponse struct {
	Result    AnyJson                      `json:"result"`
	Warning   string                       `json:"warning,omitempty"`
	Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                         `json:"forwarded,omitempty"`
}

type AnyJson map[string]interface{}

type ApiWarning struct {
	Id      int         `json:"id"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

func (r JsonRpcResponse) GetResult(v any) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v})
	if err != nil {
		return err
	}
	err = dec.Decode(r.Result)
	if err != nil {
		return err
	}
	return nil
}

func (r JsonRpcResponse) GetMarker() any {
	if _, ok := r.Result["marker"]; ok {
		return r.Result["marker"]
	}
	return nil
}

// // this will impl the XRPLPaginatedResponse
// type JsonRpcPaginatedResponse struct {
// 	results []JsonRpcResponse
// }

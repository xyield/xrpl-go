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
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json",
		Result: &v, DecodeHook: mapstructure.TextUnmarshallerHookFunc()})

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

type JsonRpcPaginationResponse struct {
	Pages []JsonRpcResponse
}

func (r JsonRpcPaginationResponse) GetXRPLPages() []client.XRPLResponse {

	res := make([]client.XRPLResponse, len(r.Pages))
	for i, page := range r.Pages {
		res[i] = page
	}

	return res
}

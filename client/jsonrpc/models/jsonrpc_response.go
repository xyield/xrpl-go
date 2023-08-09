package jsonrpcmodels

import "github.com/mitchellh/mapstructure"

type JsonRpcResponse struct {
	Result   AnyJson      `json:"result"`
	Warning  string       `json:"warning,omitempty"`
	Warnings []ApiWarning `json:"warnings,omitempty"` // TODO: update this to use the XRPLResponseWarning
	// Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool `json:"forwarded,omitempty"`
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

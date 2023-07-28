package websocket

import (
	"github.com/mitchellh/mapstructure"
	"github.com/xyield/xrpl-go/client"
)

var _ client.XRPLResponse = (*WebSocketClientXrplResponse)(nil)

type WebSocketClientXrplResponse struct {
	ID     int            `json:"id"`
	Status string         `json:"status"`
	Type   string         `json:"type"`
	Result map[string]any `json:"result"`
}

func (r *WebSocketClientXrplResponse) GetResult(v any) {
	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &v})
	_ = dec.Decode(r.Result)
}

package websocket

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/client"
)

var _ client.XRPLResponse = (*WebSocketClientXrplResponse)(nil)

type ErrorWebsocketClientXrplResponse struct {
	Type    string
	Request map[string]any
}

func (e *ErrorWebsocketClientXrplResponse) Error() string {
	return e.Type
}

type WebSocketClientXrplResponse struct {
	ID        int                          `json:"id"`
	Status    string                       `json:"status"`
	Type      string                       `json:"type"`
	Error     string                       `json:"error,omitempty"`
	Result    json.RawMessage              `json:"result,omitempty"`
	Value     map[string]any               `json:"value,omitempty"`
	Warning   string                       `json:"warning,omitempty"`
	Warnings  []client.XRPLResponseWarning `json:"warnings,omitempty"`
	Forwarded bool                         `json:"forwarded,omitempty"`
}

func (r *WebSocketClientXrplResponse) GetResult(v any) error {
	return json.Unmarshal(r.Result, v)
}

func (r *WebSocketClientXrplResponse) GetError() error {
	if r.Error != "" {
		return &ErrorWebsocketClientXrplResponse{
			Type:    r.Error,
			Request: r.Value,
		}
	}
	return nil
}

package websocket

import (
	"encoding/json"

	"github.com/mitchellh/mapstructure"
	"github.com/xyield/xrpl-go/model/client/common"
)

func (c *WebsocketClient) formatRequest(req common.XRPLRequest, id int, marker any) ([]byte, error) {
	m := make(map[string]any)
	m["id"] = id
	m["command"] = req.Method()
	if marker != nil {
		m["marker"] = marker
	}
	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &m})
	err := dec.Decode(req)
	if err != nil {
		return nil, err
	}

	return json.Marshal(m)
}

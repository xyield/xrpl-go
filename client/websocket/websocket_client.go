package websocket

import (
	"bytes"
	"encoding/json"

	"github.com/gorilla/websocket"
	"github.com/mitchellh/mapstructure"
	"github.com/xyield/xrpl-go/client"
	"github.com/xyield/xrpl-go/model/client/common"
)

var _ client.Client = (*WebsocketClient)(nil)

type websocketConfig struct {
	URL string
}

type WebsocketClient struct {
	// websocket setup
	cfg *websocketConfig
	// hub  *hub
	// conn *websocket.Conn
}

func (c *WebsocketClient) SendRequest(req common.XRPLRequest) (client.XRPLResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	// websocket specific implementation
	id := "1"
	// resChan := make(responseChan)
	// defer close(resChan)
	// c.hub.registerConnection(id, resChan)

	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.URL, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	msg, err := c.formatRequest(req, id, nil)
	if err != nil {
		return nil, err
	}

	// v := <-resChan
	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}

	_, v, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	jDec := json.NewDecoder(bytes.NewReader(v))
	jDec.UseNumber()
	var m map[string]any
	err = jDec.Decode(&m)
	if err != nil {
		return nil, err
	}
	var res WebSocketClientXrplResponse
	dec, _ := mapstructure.NewDecoder(&mapstructure.DecoderConfig{TagName: "json", Result: &res})
	err = dec.Decode(&m)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

/*
Creates a new websocket client with cfg.

This client will open and close a websocket connection for each request.

TODO: implement a websocket connection pool to handle subscriptions
*/
func NewWebsocketClient(cfg *websocketConfig) (client *WebsocketClient, err error) {
	return &WebsocketClient{
		cfg: cfg,
	}, nil
}

// func (c *WebsocketClient) open() error {
// 	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.URL, nil)
// 	if err != nil {
// 		return err
// 	}
// 	c.conn = conn
// 	return nil
// }

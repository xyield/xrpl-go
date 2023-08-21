package websocket

import (
	"bytes"
	"encoding/json"
	"errors"
	"sync/atomic"

	"github.com/gorilla/websocket"
	"github.com/xyield/xrpl-go/client"
)

var _ client.Client = (*WebsocketClient)(nil)

var ErrIncorrectId = errors.New("incorrect id")

type WebsocketConfig struct {
	URL string
}

type WebsocketClient struct {
	cfg       *WebsocketConfig
	idCounter atomic.Uint32
}

func (c *WebsocketClient) SendRequest(req client.XRPLRequest) (client.XRPLResponse, error) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}

	id := c.idCounter.Add(1)

	conn, _, err := websocket.DefaultDialer.Dial(c.cfg.URL, nil)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	msg, err := c.formatRequest(req, int(id), nil)
	if err != nil {
		return nil, err
	}

	err = conn.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		return nil, err
	}

	_, v, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	jDec := json.NewDecoder(bytes.NewReader(json.RawMessage(v)))
	jDec.UseNumber()
	var res WebSocketClientXrplResponse
	err = jDec.Decode(&res)
	if err != nil {
		return nil, err
	}

	if res.ID != int(id) {
		return nil, ErrIncorrectId
	}
	if err := res.CheckError(); err != nil {
		return nil, err
	}

	return &res, nil
}

/*
Creates a new websocket client with cfg.

This client will open and close a websocket connection for each request.
*/
func NewWebsocketClient(cfg *WebsocketConfig) *WebsocketClient {
	return &WebsocketClient{
		cfg: cfg,
	}
}

func NewClient(cfg *WebsocketConfig) *client.XRPLClient {
	wcl := &WebsocketClient{
		cfg: cfg,
	}
	return client.NewXRPLClient(wcl)
}

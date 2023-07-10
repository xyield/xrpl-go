package websocketclient

import (
	"github.com/xyield/xrpl-go/client"
	"github.com/xyield/xrpl-go/model/client/common"
)

var _ client.Client = (*WebsocketClient)(nil)

type WebsocketClient struct {
	// websocket setup
}

func (c *WebsocketClient) SendRequest(reqParams common.XRPLRequest, responseStruct common.XRPLResponse) error {

	// websocket specific implementation
	return nil
}

func NewWebsocketClient() *WebsocketClient {
	return &WebsocketClient{}
}

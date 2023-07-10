package client

import (
	"github.com/xyield/xrpl-go/model/client/common"
)

// both jsonrpc and websocket clients satisfy this
type Client interface {
	SendRequest(reqParams common.XRPLRequest, responseStruct common.XRPLResponse) error
}

type XRPLClient struct {
	Client
	Account Account
	// every method type will be added here
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		Client: cl,
		// client is set which also holds the config object
		Account: &AccountImpl{Client: cl},
	}
}

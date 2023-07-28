package client

import (
	"github.com/xyield/xrpl-go/model/client/common"
)

type Client interface {
	SendRequest(req common.XRPLRequest) (XRPLResponse, error)
}

type XRPLClient struct {
	Client
	Account Account
}

type XRPLResponse interface {
	GetResult(v any)
	// Warnings()

}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		Client:  cl,
		Account: &accountImpl{client: cl},
	}
}

// func (c *XRPLClient) GetAccountChannels(req *account.AccountChannelsRequest) *account.AccountChannelsResponse {
// 	return c.SendRequest(req).(*account.AccountChannelsResponse)
// }

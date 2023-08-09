package client

import "github.com/xyield/xrpl-go/model/client/common"

type Client interface {
	SendRequest(reqParams common.XRPLRequest) (common.XRPLResponse, error)
}

type XRPLClient struct {
	Account Account
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		Account: &accountImpl{Client: cl},
	}
}

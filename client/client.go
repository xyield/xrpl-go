package client

import (
	"github.com/xyield/xrpl-go/model/client/common"
)

type Client interface {
	SendRequest(reqParams common.XRPLRequest) (common.XRPLResponse, error)
}

type XRPLClient struct {
	Client
	Account Account
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		Client:  cl,
		Account: &accountImpl{Client: cl},
	}
}

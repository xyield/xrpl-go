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
	CheckError() error
	// Warnings()

}

type XRPLResponseWarning struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		Client:  cl,
		Account: &accountImpl{client: cl},
	}
}

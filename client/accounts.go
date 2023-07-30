package client

import (
	"github.com/xyield/xrpl-go/model/client/account"
)

type Account interface {
	GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error)
	GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, error)
}

type accountImpl struct {
	client Client
}

func (a *accountImpl) GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var acr account.AccountChannelsResponse
	res.GetResult(&acr)
	return &acr, nil
}

func (a *accountImpl) GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, err
	}
	var air account.AccountInfoResponse
	res.GetResult(&air)
	return &air, nil
}

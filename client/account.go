package client

import (
	"github.com/xyield/xrpl-go/model/client/account"
)

type Account interface {
	GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error)
}

type accountImpl struct {
	Client Client
}

func (a *accountImpl) GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, err
	}

	result, err := a.Client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var channelResponse account.AccountChannelsResponse
	err = result.GetResult(&channelResponse)
	if err != nil {
		return nil, err
	}

	return &channelResponse, nil
}

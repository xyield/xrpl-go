package client

import (
	"github.com/xyield/xrpl-go/model/client/account"
)

type Account interface {
	// return result struct, fill xrpl response for warnings etc, error
	GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, XRPLResponse, error)
}

type accountImpl struct {
	Client Client
}

func (a *accountImpl) GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, XRPLResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, nil, err
	}

	result, err := a.Client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}

	var channelResponse account.AccountChannelsResponse
	err = result.GetResult(&channelResponse)
	if err != nil {
		return nil, nil, err
	}

	return &channelResponse, result, nil
}

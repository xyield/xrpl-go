package client

import (
	"github.com/xyield/xrpl-go/model/client/account"
)

type Account interface {
	GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error)
}

type AccountImpl struct {
	Client Client
}

type AccountChannelsMissingAccountError struct {
	ErrorString string
}

func (e *AccountChannelsMissingAccountError) Error() string {
	return "Account value is missing"
}

func (a *AccountImpl) GetAccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, error) {

	// check required params are there + validate others + serialise (if required)
	if req.Account == "" {
		return nil, &AccountChannelsMissingAccountError{}
	}

	// serialise params will happen here??

	response := &account.AccountChannelsResponse{}

	err := a.Client.SendRequest(req, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

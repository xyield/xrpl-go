package client

import (
	"github.com/xyield/xrpl-go/model/client/account"
)

type Account interface {
	GetAccountChannels(req *account.AccountChannelsRequest, params XRPLPaginatedRequest) ([]account.AccountChannelsResponse, []XRPLResponse, error)
	GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error)
}

type accountImpl struct {
	client Client
}

func (a *accountImpl) GetAccountChannels(req *account.AccountChannelsRequest, params XRPLPaginatedRequest) ([]account.AccountChannelsResponse, []XRPLResponse, error) {

	// TODO; set timer to exit recurssion if continues too long?

	err := req.Validate()
	if err != nil {
		return nil, nil, err
	}

	XRPLResponsePages, err := a.client.SendRequestPaginated(req, params.Limit, params.Paginated)
	if err != nil {
		return nil, nil, err
	}

	acrPages := []account.AccountChannelsResponse{}

	// loop through pages and get result
	for _, page := range XRPLResponsePages {

		var acr account.AccountChannelsResponse

		err = page.GetResult(&acr)
		if err != nil {
			return nil, nil, err
		}

		// append result to array
		acrPages = append(acrPages, acr)
	}

	return acrPages, XRPLResponsePages, nil
}

func (a *accountImpl) GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var air account.AccountInfoResponse
	err = res.GetResult(&air)
	if err != nil {
		return nil, nil, err
	}
	return &air, res, nil
}

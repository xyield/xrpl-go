package client

import (
	"fmt"

	"github.com/xyield/xrpl-go/model/client/account"
)

type Account interface {
	GetAccountChannels(req *account.AccountChannelsRequest) ([]account.AccountChannelsResponse, XRPLResponse, error)
	GetAccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error)
}

type accountImpl struct {
	client Client
}

func (a *accountImpl) GetAccountChannels(req *account.AccountChannelsRequest) ([]account.AccountChannelsResponse, XRPLResponse, error) {

	err := req.Validate()
	if err != nil {
		return nil, nil, err
	}

	pages := []account.AccountChannelsResponse{}
	xrplResult, err := GetPages(a, req, &pages)
	if err != nil {
		return nil, nil, err
	}

	return pages, xrplResult, nil
}

// TODO: have a way for user to choose if they re-try with marker? That would require the 10 min rule being followed - trigger go routine to run and watch timer
// Call this method instead of SendRequest when response is paginated
func GetPages(a *accountImpl, req *account.AccountChannelsRequest, pages *[]account.AccountChannelsResponse) (XRPLResponse, error) {

	// get first page of results
	result, err := a.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Paginated response %v : ", result)

	// map results to struct
	var acr account.AccountChannelsResponse
	err = result.GetResult(&acr)
	if err != nil {
		return nil, err
	}

	// add page to array
	*pages = append(*pages, acr)

	// check if marker present and make new call if exists
	if acr.Marker != nil {
		req.Marker = acr.Marker
		return GetPages(a, req, pages)
	}

	return result, nil
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

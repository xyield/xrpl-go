package client

import (
	"github.com/CreatureDev/xrpl-go/model/client/account"
)

type Account interface {
	AccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, XRPLResponse, error)
	AccountCurrencies(req *account.AccountCurrenciesRequest) (*account.AccountCurrenciesResponse, XRPLResponse, error)
	AccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error)
	AccountLines(req *account.AccountLinesRequest) (*account.AccountLinesResponse, XRPLResponse, error)
	AccountNFTs(req *account.AccountNFTsRequest) (*account.AccountNFTsResponse, XRPLResponse, error)
	AccountObjects(req *account.AccountObjectsRequest) (*account.AccountObjectsResponse, XRPLResponse, error)
	AccountOffers(req *account.AccountOffersRequest) (*account.AccountOffersResponse, XRPLResponse, error)
	AccountTransactions(req *account.AccountTransactionsRequest) (*account.AccountTransactionsResponse, XRPLResponse, error)
}

type accountImpl struct {
	client Client
}

func (a *accountImpl) AccountChannels(req *account.AccountChannelsRequest) (*account.AccountChannelsResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var acr account.AccountChannelsResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, nil, err
	}
	return &acr, res, nil
}

func (a *accountImpl) AccountCurrencies(req *account.AccountCurrenciesRequest) (*account.AccountCurrenciesResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var acr account.AccountCurrenciesResponse
	err = res.GetResult(&acr)
	if err != nil {
		return nil, nil, err
	}
	return &acr, res, nil
}

func (a *accountImpl) AccountInfo(req *account.AccountInfoRequest) (*account.AccountInfoResponse, XRPLResponse, error) {
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

func (a *accountImpl) AccountLines(req *account.AccountLinesRequest) (*account.AccountLinesResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var alr account.AccountLinesResponse
	err = res.GetResult(&alr)
	if err != nil {
		return nil, nil, err
	}
	return &alr, res, nil
}

func (a *accountImpl) AccountNFTs(req *account.AccountNFTsRequest) (*account.AccountNFTsResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var anr account.AccountNFTsResponse
	err = res.GetResult(&anr)
	if err != nil {
		return nil, nil, err
	}
	return &anr, res, nil
}

func (a *accountImpl) AccountObjects(req *account.AccountObjectsRequest) (*account.AccountObjectsResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var aor account.AccountObjectsResponse
	err = res.GetResult(&aor)
	if err != nil {
		return nil, nil, err
	}
	return &aor, res, nil
}

func (a *accountImpl) AccountOffers(req *account.AccountOffersRequest) (*account.AccountOffersResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var aor account.AccountOffersResponse
	err = res.GetResult(&aor)
	if err != nil {
		return nil, nil, err
	}
	return &aor, res, nil
}

func (a *accountImpl) AccountTransactions(req *account.AccountTransactionsRequest) (*account.AccountTransactionsResponse, XRPLResponse, error) {
	res, err := a.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var atr account.AccountTransactionsResponse
	err = res.GetResult(&atr)
	if err != nil {
		return nil, nil, err
	}
	return &atr, res, nil
}

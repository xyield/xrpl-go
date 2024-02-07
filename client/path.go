package client

import "github.com/CreatureDev/xrpl-go/model/client/path"

type Path interface {
	BookOffers(req *path.BookOffersRequest) (*path.BookOffersResponse, XRPLResponse, error)
	DepositAuthorized(req *path.DepositAuthorizedRequest) (*path.DepositAuthorizedResponse, XRPLResponse, error)
	NFTokenBuyOffers(req *path.NFTokenBuyOffersRequest) (*path.NFTokenBuyOffersResponse, XRPLResponse, error)
	NFTokenSellOffers(req *path.NFTokenSellOffersRequest) (*path.NFTokenSellOffersResponse, XRPLResponse, error)
	PathFind(req *path.PathFindRequest) (*path.PathFindResponse, XRPLResponse, error)
	RipplePathFind(req *path.RipplePathFindRequest) (*path.RipplePathFindResponse, XRPLResponse, error)
}

type pathImpl struct {
	client Client
}

func (p *pathImpl) BookOffers(req *path.BookOffersRequest) (*path.BookOffersResponse, XRPLResponse, error) {
	res, err := p.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var bor path.BookOffersResponse
	err = res.GetResult(&bor)
	if err != nil {
		return nil, nil, err
	}
	return &bor, res, nil
}

func (p *pathImpl) DepositAuthorized(req *path.DepositAuthorizedRequest) (*path.DepositAuthorizedResponse, XRPLResponse, error) {
	res, err := p.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var dar path.DepositAuthorizedResponse
	err = res.GetResult(&dar)
	if err != nil {
		return nil, nil, err
	}
	return &dar, res, nil
}

func (p *pathImpl) NFTokenBuyOffers(req *path.NFTokenBuyOffersRequest) (*path.NFTokenBuyOffersResponse, XRPLResponse, error) {
	res, err := p.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var nbr path.NFTokenBuyOffersResponse
	err = res.GetResult(&nbr)
	if err != nil {
		return nil, nil, err
	}
	return &nbr, res, nil
}

func (p *pathImpl) NFTokenSellOffers(req *path.NFTokenSellOffersRequest) (*path.NFTokenSellOffersResponse, XRPLResponse, error) {
	res, err := p.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var nsr path.NFTokenSellOffersResponse
	err = res.GetResult(&nsr)
	if err != nil {
		return nil, nil, err
	}
	return &nsr, res, nil
}

func (p *pathImpl) PathFind(req *path.PathFindRequest) (*path.PathFindResponse, XRPLResponse, error) {
	res, err := p.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var pfr path.PathFindResponse
	err = res.GetResult(&pfr)
	if err != nil {
		return nil, nil, err
	}
	return &pfr, res, nil
}

func (p *pathImpl) RipplePathFind(req *path.RipplePathFindRequest) (*path.RipplePathFindResponse, XRPLResponse, error) {
	res, err := p.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var rpfr path.RipplePathFindResponse
	err = res.GetResult(&rpfr)
	if err != nil {
		return nil, nil, err
	}
	return &rpfr, res, nil
}

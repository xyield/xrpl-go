package client

import "github.com/CreatureDev/xrpl-go/model/client/clio"

type Clio interface {
	ServerInfo(*clio.ServerInfoRequest) (*clio.ServerInfoResponse, XRPLResponse, error)
	Ledger(*clio.LedgerRequest) (*clio.LedgerResponse, XRPLResponse, error)
	NFTHistory(*clio.NFTHistoryRequest) (*clio.NFTHistoryResponse, XRPLResponse, error)
	NFTInfo(*clio.NFTInfoRequest) (*clio.NFTInfoResponse, XRPLResponse, error)
}

type clioImpl struct {
	client Client
}

func (c *clioImpl) ServerInfo(req *clio.ServerInfoRequest) (*clio.ServerInfoResponse, XRPLResponse, error) {
	res, err := c.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var sr clio.ServerInfoResponse
	err = res.GetResult(&sr)
	if err != nil {
		return nil, nil, err
	}
	return &sr, res, nil
}

func (c *clioImpl) Ledger(req *clio.LedgerRequest) (*clio.LedgerResponse, XRPLResponse, error) {
	res, err := c.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var lr clio.LedgerResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, nil, err
	}
	return &lr, res, nil
}

func (c *clioImpl) NFTHistory(req *clio.NFTHistoryRequest) (*clio.NFTHistoryResponse, XRPLResponse, error) {
	res, err := c.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var nr clio.NFTHistoryResponse
	err = res.GetResult(&nr)
	if err != nil {
		return nil, nil, err
	}
	return &nr, res, nil
}

func (c *clioImpl) NFTInfo(req *clio.NFTInfoRequest) (*clio.NFTInfoResponse, XRPLResponse, error) {
	res, err := c.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var nr clio.NFTInfoResponse
	err = res.GetResult(&nr)
	if err != nil {
		return nil, nil, err
	}
	return &nr, res, nil
}

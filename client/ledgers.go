package client

import (
	"github.com/xyield/xrpl-go/model/client/ledger"
)

type Ledger interface {
	LedgerClosed(req *ledger.LedgerClosedRequest) (*ledger.LedgerClosedRequest, XRPLResponse, error)
	LedgerCurrent(req *ledger.LedgerCurrentRequest) (*ledger.LedgerCurrentRequest, XRPLResponse, error)
	LedgerData(req *ledger.LedgerDataRequest) (*ledger.LedgerDataRequest, XRPLResponse, error)
	LedgerEntry(req *ledger.LedgerEntryRequest) (*ledger.LedgerEntryRequest, XRPLResponse, error)
	Ledger(req *ledger.LedgerRequest) (*ledger.LedgerRequest, XRPLResponse, error)
}

type ledgerImpl struct {
	client Client
}

func (l *ledgerImpl) LedgerClosed(req *ledger.LedgerClosedRequest) (*ledger.LedgerClosedResponse, XRPLResponse, error) {
	res, err := l.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var lcr ledger.LedgerClosedResponse
	err = res.GetResult(&lcr)
	if err != nil {
		return nil, nil, err
	}
	return &lcr, res, nil
}

func (l *ledgerImpl) LedgerCurrent(req *ledger.LedgerCurrentRequest) (*ledger.LedgerCurrentResponse, XRPLResponse, error) {
	res, err := l.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var lcr ledger.LedgerCurrentResponse
	err = res.GetResult(&lcr)
	if err != nil {
		return nil, nil, err
	}
	return &lcr, res, nil
}

func (l *ledgerImpl) LedgerData(req *ledger.LedgerDataRequest) (*ledger.LedgerDataResponse, XRPLResponse, error) {
	res, err := l.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ldr ledger.LedgerDataResponse
	err = res.GetResult(&ldr)
	if err != nil {
		return nil, nil, err
	}
	return &ldr, res, nil
}

func (l *ledgerImpl) LedgerEntry(req *ledger.LedgerEntryRequest) (*ledger.LedgerEntryResponse, XRPLResponse, error) {
	res, err := l.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ler ledger.LedgerEntryResponse
	err = res.GetResult(&ler)
	if err != nil {
		return nil, nil, err
	}
	return &ler, res, nil
}

func (l *ledgerImpl) Ledger(req *ledger.LedgerRequest) (*ledger.LedgerResponse, XRPLResponse, error) {
	res, err := l.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var lr ledger.LedgerResponse
	err = res.GetResult(&lr)
	if err != nil {
		return nil, nil, err
	}
	return &lr, res, nil
}

package client

import (
	"github.com/CreatureDev/xrpl-go/model/client/transactions"
)

type Transaction interface {
	SubmitMultisigned(*transactions.SubmitMultisignedRequest) (*transactions.SubmitMultisignedResponse, XRPLResponse, error)
	Submit(*transactions.SubmitRequest) (*transactions.SubmitResponse, XRPLResponse, error)
	TransactionEntry(*transactions.TransactionEntryRequest) (*transactions.TransactionEntryResponse, XRPLResponse, error)
	Tx(*transactions.TxRequest) (*transactions.TxResponse, XRPLResponse, error)
}

type transactionImpl struct {
	client Client
}

func (t *transactionImpl) SubmitMultisigned(req *transactions.SubmitMultisignedRequest) (*transactions.SubmitMultisignedResponse, XRPLResponse, error) {
	res, err := t.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var smr transactions.SubmitMultisignedResponse
	err = res.GetResult(&smr)
	if err != nil {
		return nil, nil, err
	}
	return &smr, res, nil
}

func (t *transactionImpl) Submit(req *transactions.SubmitRequest) (*transactions.SubmitResponse, XRPLResponse, error) {
	res, err := t.client.SendRequest(req)
	if err != nil {
		return nil, res, err
	}
	var sr transactions.SubmitResponse
	err = res.GetResult(&sr)
	if err != nil {
		return nil, res, err
	}
	return &sr, res, nil
}

func (t *transactionImpl) TransactionEntry(req *transactions.TransactionEntryRequest) (*transactions.TransactionEntryResponse, XRPLResponse, error) {
	res, err := t.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ter transactions.TransactionEntryResponse
	err = res.GetResult(&ter)
	if err != nil {
		return nil, nil, err
	}
	return &ter, res, nil
}

func (t *transactionImpl) Tx(req *transactions.TxRequest) (*transactions.TxResponse, XRPLResponse, error) {
	res, err := t.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var tr transactions.TxResponse
	err = res.GetResult(&tr)
	if err != nil {
		return nil, nil, err
	}
	return &tr, res, nil
}

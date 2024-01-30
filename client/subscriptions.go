package client

import (
	"github.com/CreatureDev/xrpl-go/model/client/subscription"
)

type Subscription interface {
	Subscribe(*subscription.SubscribeRequest) (*subscription.SubscribeResponse, XRPLResponse, error)
	Unsubscribe(*subscription.UnsubscribeRequest) (*subscription.UnsubscribeResponse, XRPLResponse, error)
}

type subscriptionImpl struct {
	client Client
}

func (s *subscriptionImpl) Subscribe(req *subscription.SubscribeRequest) (*subscription.SubscribeResponse, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var sr subscription.SubscribeResponse
	err = res.GetResult(&sr)
	if err != nil {
		return nil, nil, err
	}
	return &sr, res, nil
}

func (s *subscriptionImpl) Unsubscribe(req *subscription.UnsubscribeRequest) (*subscription.UnsubscribeResponse, XRPLResponse, error) {
	res, err := s.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var ur subscription.UnsubscribeResponse
	err = res.GetResult(&ur)
	if err != nil {
		return nil, nil, err
	}
	return &ur, res, nil
}

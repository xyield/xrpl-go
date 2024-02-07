package client

import (
	"github.com/CreatureDev/xrpl-go/model/client/channel"
)

type Channel interface {
	ChannelAuthorize(req *channel.ChannelAuthorizeRequest) (*channel.ChannelAuthorizeResponse, XRPLResponse, error)
	ChannelVerify(req *channel.ChannelVerifyRequest) (*channel.ChannelVerifyResponse, XRPLResponse, error)
}

type channelImpl struct {
	client Client
}

func (c *channelImpl) ChannelAuthorize(req *channel.ChannelAuthorizeRequest) (*channel.ChannelAuthorizeResponse, XRPLResponse, error) {
	res, err := c.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var car channel.ChannelAuthorizeResponse
	err = res.GetResult(&car)
	if err != nil {
		return nil, nil, err
	}
	return &car, res, nil
}

func (c *channelImpl) ChannelVerify(req *channel.ChannelVerifyRequest) (*channel.ChannelVerifyResponse, XRPLResponse, error) {
	res, err := c.client.SendRequest(req)
	if err != nil {
		return nil, nil, err
	}
	var cvr channel.ChannelVerifyResponse
	err = res.GetResult(&cvr)
	if err != nil {
		return nil, nil, err
	}
	return &cvr, res, nil
}

package accountmethods

import xrplclient "github.com/xyield/xrpl-go/xrpl-jsonrpc-client"

type AccountMethodsInterface interface {
	Channels(params AccountChannelsParams) (interface{}, error) // TODO: make parameters interface to pass in here, return a model
}

type AccountMethods struct {
	Cfg *xrplclient.Config
}

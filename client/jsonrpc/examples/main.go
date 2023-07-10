package jsonrpcexamples

import (
	"log"

	"github.com/xyield/xrpl-go/client"
	jsonrpcclient "github.com/xyield/xrpl-go/client/jsonrpc"
	"github.com/xyield/xrpl-go/model/client/account"
)

func main() {

	cfg, err := client.NewJsonRpcConfig("http://testnode/")
	if err != nil {
		log.Panicln(err)
	}

	// Initialise new json client with json config
	jsonrpc := jsonrpcclient.NewJsonRpcClient(cfg)

	// create new XRPL client with the json client
	xrplClient := client.NewXRPLClient(jsonrpc)

	var req *account.AccountChannelsRequest

	xrplClient.Account.GetAccountChannels(req)

}

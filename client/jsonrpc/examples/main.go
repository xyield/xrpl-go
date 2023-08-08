package jsonrpcexamples

import (
	"log"

	"github.com/xyield/xrpl-go/client"
	jsonrpcclient "github.com/xyield/xrpl-go/client/jsonrpc"
	"github.com/xyield/xrpl-go/model/client/account"
)

func main() {

	// init new config object with desired node address
	cfg, err := client.NewJsonRpcConfig("http://testnode/")
	if err != nil {
		log.Panicln(err)
	}

	// Initialise new json client with json config
	jsonrpcClient := jsonrpcclient.NewJsonRpcClient(cfg)

	// create new XRPL client with the json client
	xrplClient := client.NewXRPLClient(jsonrpcClient)

	var req *account.AccountChannelsRequest

	// call the desired methodx
	xrplClient.Account.GetAccountChannels(req)

}

package main

import (
	"fmt"
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
	client := jsonrpcclient.NewClient(cfg)

	// call the desired method
	var req *account.AccountChannelsRequest
	ac, xrplRes, err := client.Account.AccountChannels(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Full XRPL response: %v\n", xrplRes)
	fmt.Printf("Results mapped to struct: %v\n", ac)
}

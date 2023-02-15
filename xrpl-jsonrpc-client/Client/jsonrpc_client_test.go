package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
	jsonrpcclient "github.com/xyield/xrpl-go/xrpl-jsonrpc-client"
	accountmethods "github.com/xyield/xrpl-go/xrpl-jsonrpc-client/AccountMethods"
)

func TestClientCreation(t *testing.T) {

	t.Run("Set config with valid port + ip", func(t *testing.T) {

		cfg, _ := jsonrpcclient.NewConfig("url")

		client := New(cfg)

		assert.Equal(t, &XrplJsonRpcClient{Config: cfg, AccountMethods: &accountmethods.AccountMethods{
			Cfg: cfg,
		}}, client)
	})
}

package jsonrpcmodels

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xyield/xrpl-go/client"
	"github.com/xyield/xrpl-go/model/client/account"
)

func TestGetResult(t *testing.T) {
	t.Run("correctly decodes", func(t *testing.T) {

		jr := JsonRpcResponse{
			Result: AnyJson{
				"account":      "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
				"ledger_hash":  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
				"ledger_index": json.Number(strconv.FormatInt(71766343, 10)),
			},
			Warning: "none",
			Warnings: []client.XRPLResponseWarning{{
				Id:      "1",
				Message: "message",
			},
			},
		}

		expected := account.AccountChannelsResponse{
			Account:     "rf1BiGeXwwQoi8Z2ueFYTEXSwuJYfV2Jpn",
			LedgerHash:  "27F530E5C93ED5C13994812787C1ED073C822BAEC7597964608F2C049C2ACD2D",
			LedgerIndex: 71766343,
		}

		var acr account.AccountChannelsResponse
		err := jr.GetResult(&acr)

		assert.NoError(t, err)
		assert.Equal(t, expected, acr)
	})
}

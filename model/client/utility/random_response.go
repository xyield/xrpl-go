package utility

import "github.com/xyield/xrpl-go/model/transactions/types"

type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}

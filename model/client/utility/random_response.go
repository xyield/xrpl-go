package utility

import "github.com/xyield/xrpl-go/binary-codec/types"

type RandomResponse struct {
	Random types.Hash256 `json:"random"`
}

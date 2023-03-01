package transactions

import (
	"github.com/xyield/xrpl-go/model/transactions/types"
)

type Signer struct {
	Account       types.Address
	TxnSignature  []byte
	SigningPubKey []byte
}

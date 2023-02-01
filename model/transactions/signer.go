package transactions

import (
	. "github.com/xyield/xrpl-go/model/transactions/types"
)

type Signer struct {
	Account       Address
	TxnSignature  []byte
	SigningPubKey []byte
}

package transactions

type Signer struct {
	Account       Address
	TxnSignature  []byte
	SigningPubKey []byte
}

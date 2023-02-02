package transactions

import (
	"encoding/json"

	. "github.com/xyield/xrpl-go/model/ledger"
)

type SignerListSet struct {
	BaseTx
	SignerQuorum  uint
	SignerEntries []SignerEntry
}

func (*SignerListSet) TxType() TxType {
	return SignerListSetTx
}

func UnmarshalSignerListSetTx(data json.RawMessage) (Tx, error) {
	var ret SignerListSet
	if err := json.Unmarshal(data, &ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

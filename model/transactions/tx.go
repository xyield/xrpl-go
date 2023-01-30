package transactions

import (
	"encoding/json"
	"fmt"
)

type Tx interface {
	TxType() TxType
}

func UnmarshalTx(data json.RawMessage) (Tx, error) {
	// TODO
	return nil, fmt.Errorf("Unimplemented")
}

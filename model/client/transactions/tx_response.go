package transactions

import (
	"encoding/json"

	"github.com/CreatureDev/xrpl-go/model/client/common"
	"github.com/CreatureDev/xrpl-go/model/transactions"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

type TxResponse struct {
	Date        uint                `json:"date"`
	Hash        types.Hash256       `json:"hash"`
	LedgerIndex common.LedgerIndex  `json:"ledger_index"`
	Meta        transactions.TxMeta `json:"meta"`
	Validated   bool                `json:"validated"`
	Tx          transactions.Tx     `json:",omitempty"`
}

// Custom marshal in order to embed transaction fields on the lowerst level
func (r TxResponse) MarshalJSON() ([]byte, error) {
	type txrHelper TxResponse
	c := txrHelper(r)
	c.Tx = nil
	first, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	second, err := json.Marshal(r.Tx)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(first, &data)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(second, &data)
	if err != nil {
		return nil, err
	}
	return json.Marshal(data)
}

func (r *TxResponse) UnmarshalJSON(data []byte) error {
	type txrHelper struct {
		Date        uint               `json:"date"`
		Hash        types.Hash256      `json:"hash"`
		LedgerIndex common.LedgerIndex `json:"ledger_index"`
		Meta        json.RawMessage    `json:"meta"`
		Validated   bool               `json:"validated"`
	}
	var h txrHelper
	err := json.Unmarshal(data, &h)
	if err != nil {
		return err
	}
	*r = TxResponse{
		Date:        h.Date,
		Hash:        h.Hash,
		LedgerIndex: h.LedgerIndex,
		Validated:   h.Validated,
	}
	r.Meta, err = transactions.UnmarshalTxMeta(h.Meta)
	if err != nil {
		return err
	}
	r.Tx, err = transactions.UnmarshalTx(data)
	return err
}

package transactions

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/model/transactions"
)

type SubmitMultisignedRequest struct {
	Tx       transactions.Tx `json:"tx_json"`
	FailHard bool            `json:"fail_hard"`
}

func (*SubmitMultisignedRequest) Method() string {
	return "submit_multisigned"
}

func (r *SubmitMultisignedRequest) UnmarshalJSON(data []byte) error {
	type sHelper struct {
		Tx       json.RawMessage `json:"tx_json"`
		FailHard bool            `json:"fail_hard"`
	}
	var h sHelper
	if err := json.Unmarshal(data, &h); err != nil {
		return err
	}
	*r = SubmitMultisignedRequest{
		FailHard: h.FailHard,
	}
	tx, err := transactions.UnmarshalTx(h.Tx)
	if err != nil {
		return err
	}
	r.Tx = tx

	return nil
}

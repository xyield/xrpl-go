package transactions

import "fmt"

type SubmitRequest struct {
	TxBlob   string `json:"tx_blob"`
	FailHard bool   `json:"fail_hard,omitempty"`
}

func (*SubmitRequest) Method() string {
	return "submit"
}

func (s *SubmitRequest) Validate() error {
	if s.TxBlob == "" {
		return fmt.Errorf("submit request: missing txblob")
	}
	return nil
}

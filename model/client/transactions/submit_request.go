package transactions

type SubmitRequest struct {
	TxBlob   string `json:"tx_blob"`
	FailHard bool   `json:"fail_hard,omitempty"`
}

func (*SubmitRequest) Method() string {
	return "submit"
}

func (*SubmitRequest) Validate() error {
	return nil
}

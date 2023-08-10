package status

type ValidatorInfoResponse struct {
	Domain       string `json:"domain,omitempty"`
	EphemeralKey string `json:"ephemeral_key,omitempty"`
	Manifest     string `json:"manifest,omitempty"`
	MasterKey    string `json:"master_key"`
	Seq          int    `json:"seq,omitempty"`
}

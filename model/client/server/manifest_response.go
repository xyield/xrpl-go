package server

type ManifestResponse struct {
	Details   ManifestDetails `json:"details,omitempty"`
	Manifest  string          `json:"manifest,omitempty"`
	Requested string          `json:"requested"`
}

type ManifestDetails struct {
	Domain       string `json:"domain"`
	EphemeralKey string `json:"ephemeral_key"`
	MasterKey    string `json:"master_key"`
	Seq          uint   `json:"seq"`
}

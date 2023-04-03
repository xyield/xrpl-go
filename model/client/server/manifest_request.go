package server

type ManifestRequest struct {
	PublicKey string `json:"public_key"`
}

func (*ManifestRequest) Method() string {
	return "manifest"
}

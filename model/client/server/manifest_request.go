package server

import "fmt"

type ManifestRequest struct {
	PublicKey string `json:"public_key"`
}

func (*ManifestRequest) Method() string {
	return "manifest"
}

func (r *ManifestRequest) Validate() error {
	if r.PublicKey == "" {
		return fmt.Errorf("manifest request: public key not set")
	}
	return nil
}

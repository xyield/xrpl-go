package key

type ValidationCreateResponse struct {
	ValidationKey       string `json:"validation_key"`
	ValidationPublicKey string `json:"validation_public_key"`
	ValidationSeed      string `json:"validation_seed"`
}

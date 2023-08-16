package status

type ValidatorsResponse struct {
	LocalStaticKeys      []string          `json:"local_static_keys"`
	PublisherLists       []PublisherList   `json:"publisher_lists"`
	SigningKeys          map[string]string `json:"signing_keys"`
	TrustedValidatorKeys []string          `json:"trusted_validator_keys"`
	ValidationQuorum     int               `json:"validation_quorum"`
	ValidatorListExpires string            `json:"validator_list_expires"`
}

type PublisherList struct {
	Available       bool     `json:"available"`
	Expiration      string   `json:"expiration"`
	List            []string `json:"list"`
	PubkeyPublisher string   `json:"pubkey_publisher"`
	Seq             uint     `json:"seq"`
	Version         uint     `json:"version"`
}

package key

type ValidationCreateRequest struct {
	Secret string `json:"secret,omitempty"`
}

func (*ValidationCreateRequest) Method() string {
	return "validation_create"
}

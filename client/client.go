package client

type Client interface {
	SendRequest(reqParams XRPLRequest) (XRPLResponse, error)
}

type XRPLClient struct {
	Account Account
}

type XRPLRequest interface {
	Method() string
}

type XRPLResponse interface {
	GetResult(v any) error
}

type XRPLResponseWarning struct {
	Id      string `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		Account: &accountImpl{Client: cl},
	}
}

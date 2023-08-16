package client

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
}

type XRPLClient struct {
	Client
	Account Account
}

type XRPLRequest interface {
	Method() string
	Validate() error
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
		Client:  cl,
		Account: &accountImpl{client: cl},
	}
}

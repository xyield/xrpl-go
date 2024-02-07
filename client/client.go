package client

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
}

type XRPLClient struct {
	client       Client
	Account      Account
	Channel      Channel
	Ledger       Ledger
	Path         Path
	Subscription Subscription
	Transaction  Transaction
}

type XRPLRequest interface {
	Method() string
	Validate() error
}

type XRPLResponse interface {
	GetResult(v any) error
	GetError() error
}

type XRPLResponseWarning struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		client:       cl,
		Account:      &accountImpl{client: cl},
		Channel:      &channelImpl{client: cl},
		Ledger:       &ledgerImpl{client: cl},
		Path:         &pathImpl{client: cl},
		Subscription: &subscriptionImpl{client: cl},
		Transaction:  &transactionImpl{client: cl},
	}
}

func (c *XRPLClient) Client() Client {
	return c.client
}

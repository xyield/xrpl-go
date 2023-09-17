package client

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
	SendRequestPaginated(reqParams XRPLPaginatedRequest, limit int, pagination bool) (XRPLPaginatedResponse, error)
}

type XRPLClient struct {
	client  Client
	Account Account
}

type XRPLRequest interface {
	Method() string
	Validate() error
}

type XRPLResponse interface {
	GetResult(v any) error
	GetMarker() any
}

type XRPLPaginatedParams struct {
	Limit     int
	Paginated bool
}

type XRPLPaginatedRequest interface {
	Method() string
	Validate() error
	SetMarker(m any)
}

type XRPLPaginatedResponse interface {
	GetXRPLPages() []XRPLResponse
}

type XRPLResponseWarning struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Details any    `json:"details,omitempty"`
}

func NewXRPLClient(cl Client) *XRPLClient {
	return &XRPLClient{
		client:  cl,
		Account: &accountImpl{client: cl},
	}
}

func (c *XRPLClient) Client() Client {
	return c.client
}

package client

type Client interface {
	SendRequest(req XRPLRequest) (XRPLResponse, error)
	SendRequestPaginated(reqParams XRPLRequest, limit int, pagination bool) ([]XRPLResponse, error)
}

type XRPLClient struct {
	client  Client
	Account Account
}

type XRPLRequest interface {
	Method() string
	Validate() error
	SetMarker(m any) // TODO: take out of interface as only pag ones need it - make new ones for the req and response
}

type XRPLResponse interface {
	GetResult(v any) error
	GetMarker() any
}

type XRPLPaginatedRequest struct {
	Limit     int
	Paginated bool
}

// type XRPLPaginatedResponse interface { // structs in both clients will impl this
// 	GetMarker() any
// }

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

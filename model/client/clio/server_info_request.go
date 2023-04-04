package clio

type ServerInfoRequest struct {
}

func (*ServerInfoRequest) Method() string {
	return "server_info"
}

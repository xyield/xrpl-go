package server

type ServerInfoRequest struct {
}

func (*ServerInfoRequest) Method() string {
	return "server_info"
}

func (*ServerInfoRequest) Validate() error {
	return nil
}

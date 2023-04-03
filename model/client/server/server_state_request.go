package server

type ServerStateRequest struct {
}

func (*ServerStateRequest) Method() string {
	return "server_state"
}

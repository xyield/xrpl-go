package server

type StopRequest struct {
}

func (*StopRequest) Method() string {
	return "stop"
}

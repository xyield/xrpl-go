package server

type StopRequest struct {
}

func (*StopRequest) Method() string {
	return "stop"
}

func (*StopRequest) Validate() error {
	return nil
}

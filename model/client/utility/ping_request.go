package utility

type PingRequest struct{}

func (*PingRequest) Method() string {
	return "ping"
}

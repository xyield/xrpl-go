package data

type LogrotateRequest struct {
}

func (*LogrotateRequest) Method() string {
	return "logrotate"
}

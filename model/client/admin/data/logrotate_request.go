package data

type LogrotateRequest struct {
}

func (*LogrotateRequest) Method() string {
	return "logrotate"
}

func (*LogrotateRequest) Validate() error {
	return nil
}

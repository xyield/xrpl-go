package data

type LogLevelRequest struct {
	Severity  string `json:"severity,omitempty"`
	Partition string `json:"partition,omitempty"`
}

func (*LogLevelRequest) Method() string {
	return "log_level"
}

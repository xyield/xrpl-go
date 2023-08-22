package data

const (
	Fatal LogSeverity = "fatal"
	Error LogSeverity = "error"
	Warn  LogSeverity = "warn"
	Info  LogSeverity = "info"
	Debug LogSeverity = "debug"
	Trace LogSeverity = "trace"
)

type LogSeverity string

type LogLevelRequest struct {
	Severity  LogSeverity `json:"severity,omitempty"`
	Partition string      `json:"partition,omitempty"`
}

func (*LogLevelRequest) Method() string {
	return "log_level"
}

func (*LogLevelRequest) Validate() error {
	return nil
}

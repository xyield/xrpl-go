package common

type XRPLRequest interface {
	Method() string
	Validate() error
}

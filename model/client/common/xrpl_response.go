package common

type XRPLResponse interface {
	GetResult(v any) error
}

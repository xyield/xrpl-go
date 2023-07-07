package server

type FeeRequest struct {
}

func (*FeeRequest) Method() string {
	return "fee"
}

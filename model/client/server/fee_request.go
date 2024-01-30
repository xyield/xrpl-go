package server

type FeeRequest struct {
}

func (*FeeRequest) Method() string {
	return "fee"
}

func (*FeeRequest) Validate() error {
	return nil
}

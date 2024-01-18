package utility

type RandomRequest struct{}

func (*RandomRequest) Method() string {
	return "random"
}

func (*RandomRequest) Validate() error {
	return nil
}

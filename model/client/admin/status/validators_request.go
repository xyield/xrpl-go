package status

type ValidatorsRequest struct {
}

func (*ValidatorsRequest) Method() string {
	return "validators"
}

func (*ValidatorsRequest) Validate() error {
	return nil
}

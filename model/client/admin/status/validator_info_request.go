package status

type ValidatorInfoRequest struct {
}

func (*ValidatorInfoRequest) Method() string {
	return "validator_info"
}

func (*ValidatorInfoRequest) Validate() error {
	return nil
}

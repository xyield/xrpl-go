package status

type ValidatorsRequest struct {
}

func (*ValidatorsRequest) Method() string {
	return "validators"
}

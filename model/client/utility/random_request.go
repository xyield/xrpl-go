package utility

type RandomRequest struct{}

func (*RandomRequest) Method() string {
	return "random"
}

package status

type PrintRequest struct {
}

func (*PrintRequest) Method() string {
	return "print"
}

func (*PrintRequest) Validate() error {
	return nil
}

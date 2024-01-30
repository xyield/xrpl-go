package status

type FetchInfoRequest struct {
	Clear bool `json:"clear"`
}

func (*FetchInfoRequest) Method() string {
	return "fetch_info"
}

func (*FetchInfoRequest) Validate() error {
	return nil
}

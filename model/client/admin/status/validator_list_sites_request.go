package status

type ValidatorListSitesRequest struct {
}

func (*ValidatorListSitesRequest) Method() string {
	return "validator_list_sites"
}

func (*ValidatorListSitesRequest) Validate() error {
	return nil
}

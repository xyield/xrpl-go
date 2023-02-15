package accountmethods

import (
	jsonrpc "github.com/xyield/xrpl-go/xrpl-jsonrpc-client/jsonrpc"
)

// impl RequestParams. User will pass this struct into the channels method
type AccountChannelsParams struct {
	Account            string `json:"account"`
	DestinationAccount string `json:"destination_account,omitempty"`
	LedgerHash         string `json:"ledger_hash,omitempty"`
	LedgerIndex        string `json:"ledger_index,omitempty"` // could also be unsigned int?
	Limit              int    `json:"limit,omitempty"`
	// marker             Marker `json:"marker,omitempty"`
}

type AccountChannelsResponse struct {
}

type AccountChannelsMissingAccountError struct {
	ErrorString string
}

func (e *AccountChannelsMissingAccountError) Error() string {
	return "Account value is missing"
}

func (a *AccountMethods) Channels(params AccountChannelsParams) (interface{}, error) {

	// TODO: check required params are there + validate others

	if params.Account == "" {
		return nil, &AccountChannelsMissingAccountError{}
	}

	// TODO: CreateRequest will be different for each method - some will need serialising
	body, err := jsonrpc.CreateRequest("account_currencies", params)
	if err != nil {
		return nil, err
	}

	response := AccountChannelsResponse{}

	err = jsonrpc.SendRequest(body, a.Cfg, response)
	if err != nil {
		return nil, err
	}

	// Return response struct if all gone successfully
	return response, nil
}

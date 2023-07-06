package accountmethods

import (
	"encoding/json"

	"github.com/xyield/xrpl-go/xrpl-jsonrpc-client/jsonrpc"
)

// impl RequestParams. User will pass this struct into the channels method
type AccountChannelsParams struct {
	Account            string `json:"account"`
	DestinationAccount string `json:"destination_account,omitempty"`
	LedgerHash         string `json:"ledger_hash,omitempty"`
	LedgerIndex        string `json:"ledger_index,omitempty"` // could also be unsigned int?
	Limit              int    `json:"limit,omitempty"`
}

type AccountChannelsResponse struct {
}

func (r *AccountChannelsResponse) UnmarshallJSON(data []byte) error {
	type Alias AccountChannelsResponse
	var aux Alias
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	*r = AccountChannelsResponse(aux)
	return nil
}

type AccountChannelsMissingAccountError struct {
	ErrorString string
}

func (e *AccountChannelsMissingAccountError) Error() string {
	return "Account value is missing"
}

func (a *AccountMethods) Channels(params AccountChannelsParams) (interface{}, error) {

	// check required params are there + validate others + serialise (if required)
	if params.Account == "" {
		return nil, &AccountChannelsMissingAccountError{}
	}

	// serialise params will happen here

	body, err := jsonrpc.CreateRequest("account_currencies", params)
	if err != nil {
		return nil, err
	}

	response := &AccountChannelsResponse{}

	_, err = jsonrpc.SendRequest(body, a.Cfg, response)
	if err != nil {
		return nil, err
	}

	// Return response struct if all gone successfully - unmarshalled into response in SendRequest()
	return response, nil
}

package common

// all response structs will implement this interface
// used in sendRequest to marshall response body into response struct
// example impl found in account_channels_response.go
type XRPLResponse interface {
	UnmarshallJSON([]byte) error
}

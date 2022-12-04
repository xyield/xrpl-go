package types

import (
	addresscodec "github.com/xyield/xrpl-go/address-codec"
	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type AccountID struct{}

// Serializes the given json value to an AccountID byte slice.
func (a *AccountID) FromJson(value any) ([]byte, error) {

	_, accountID, err := addresscodec.DecodeClassicAddressToAccountID(value.(string))

	if err != nil {
		return nil, err
	}

	//AccountIDs that appear as stand-alone fields (such as Account and Destination)
	// are length-prefixed despite being a fixed 160 bits in length. As a result,
	// the length indicator for these fields is always the byte 0x14.
	//
	// AccountIDs that appear as children of special fields (Amount issuer and PathSet account) are not length-prefixed.
	// So in Amount and PathSet fields, don't use the length indicator 0x14.

	return accountID, nil
}

func (a *AccountID) FromParser(p *serdes.BinaryParser) ([]byte, error) {
	return nil, nil
}

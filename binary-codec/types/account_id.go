package types

import (
	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

type AccountID struct{}

func (a *AccountID) SerializeJson(value any) ([]byte, error) {

	_, accountID, err := addresscodec.DecodeClassicAddressToAccountID(value.(string))

	if err != nil {
		return nil, err
	}
	return append([]byte{0x14}, accountID...), nil
}

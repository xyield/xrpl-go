package types

import "encoding/hex"

type Blob struct{}

func (b *Blob) SerializeJson(json any) ([]byte, error) {
	v, err := hex.DecodeString(json.(string))
	if err != nil {
		return nil, err
	}
	return v, nil
}

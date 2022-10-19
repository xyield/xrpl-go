package types

import (
	"fmt"
)

type STArray struct{}

func (t *STArray) SerializeJson(json any) ([]byte, error) {
	if _, ok := json.([]any); !ok {
		return nil, fmt.Errorf("not a slice of objects")
	}

	var sink []byte
	for _, v := range json.([]any) {
		st := &STObject{}
		b, err := st.SerializeJson(v)
		if err != nil {
			return nil, err
		}
		sink = append(sink, b...)
	}
	sink = append(sink, 0xF1)

	return sink, nil
}

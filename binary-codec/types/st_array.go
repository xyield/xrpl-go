package types

import (
	"fmt"

	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

type STArray struct{}

func (t *STArray) FromJson(json any) ([]byte, error) {
	if _, ok := json.([]any); !ok {
		return nil, fmt.Errorf("not a slice of objects")
	}

	var sink []byte
	for _, v := range json.([]any) {
		st := &STObject{}
		b, err := st.FromJson(v)
		if err != nil {
			return nil, err
		}
		sink = append(sink, b...)
	}
	sink = append(sink, 0xF1)

	return sink, nil
}

func (t *STArray) FromParser(p *serdes.BinaryParser, opts ...int) (any, error) {
	return nil, nil
}

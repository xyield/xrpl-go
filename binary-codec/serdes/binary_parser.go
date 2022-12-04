package serdes

import (
	"fmt"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

type BinaryParser struct {
	data   []byte
	cursor int
}

func NewBinaryParser(d []byte) *BinaryParser {
	return &BinaryParser{
		data: d,
	}
}

func (p *BinaryParser) ReadField() (*definitions.FieldInstance, error) {
	fh, err := p.readFieldHeader()
	if err != nil {
		return nil, err
	}
	fn, err := definitions.Get().GetFieldNameByFieldHeader(*fh)
	if err != nil {
		return nil, err
	}
	f, err := definitions.Get().GetFieldInstanceByFieldName(fn)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	fmt.Println(f)
	return nil, nil
}

func (p *BinaryParser) readFieldHeader() (*definitions.FieldHeader, error) {
	typeCode := p.readByte()
	fieldCode := typeCode & 15
	typeCode = typeCode >> 4
	return &definitions.FieldHeader{
		TypeCode:  int32(typeCode),
		FieldCode: int32(fieldCode),
	}, nil
}

func (p *BinaryParser) readByte() byte {
	b := p.data[p.cursor]
	p.cursor++
	return b
}

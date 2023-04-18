package serdes

import (
	"errors"

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
	return f, nil
}

func (p *BinaryParser) readFieldHeader() (*definitions.FieldHeader, error) {
	typeCode, _ := p.readByte()
	fieldCode := typeCode & 15
	typeCode = typeCode >> 4

	if typeCode == 0 {
		typeCode, _ = p.readByte()
		if typeCode == 0 || typeCode < 16 {
			return nil, errors.New("invalid typecode")
		}
	}

	if fieldCode == 0 {
		fieldCode, _ := p.readByte()
		if fieldCode == 0 || fieldCode < 16 {
			return nil, errors.New("invalid fieldcode")
		}
	}
	return &definitions.FieldHeader{
		TypeCode:  int32(typeCode),
		FieldCode: int32(fieldCode),
	}, nil
}

func (p *BinaryParser) readByte() (byte, error) {
	b := p.data[p.cursor]
	p.cursor++
	return b, nil
}

func (p *BinaryParser) ReadBytes(n int) ([]byte, error) {
	var bytes []byte
	for i := 0; i < n; i++ {
		b, err := p.readByte()
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, b)
	}
	return bytes, nil
}

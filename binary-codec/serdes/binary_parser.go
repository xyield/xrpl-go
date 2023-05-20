package serdes

import (
	"errors"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

var (
	ErrParserOutOfBound error = errors.New("parser out of bounds")
)

type BinaryParser struct {
	data []byte
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
	// Read the first byte of the field header
	typeCode, _ := p.ReadByte()

	// The field code is the last 4 bits of the first byte
	fieldCode := typeCode & 15
	typeCode = typeCode >> 4

	// Read the type code if it's not in the first byte
	if typeCode == 0 {
		typeCode, _ = p.ReadByte()
		if typeCode == 0 || typeCode < 16 {
			return nil, errors.New("invalid typecode")
		}
	}

	// Read the field code if it's not in the first byte
	if fieldCode == 0 {
		fieldCode, _ = p.ReadByte()
		if fieldCode == 0 || fieldCode < 16 {
			return nil, errors.New("invalid fieldcode")
		}
	}

	// Return the field header
	return &definitions.FieldHeader{
		TypeCode:  int32(typeCode),
		FieldCode: int32(fieldCode),
	}, nil
}

func (p *BinaryParser) ReadByte() (byte, error) {
	if len(p.data) < 1 {
		return 0, ErrParserOutOfBound
	}
	b := p.data[0]
	p.data = p.data[1:]
	return b, nil
}

func (p *BinaryParser) Peek() (byte, error) {
	if len(p.data) < 1 {
		return 0, ErrParserOutOfBound
	}
	return p.data[0], nil
}

func (p *BinaryParser) ReadBytes(n int) ([]byte, error) {
	var bytes []byte
	for i := 0; i < n; i++ {
		b, err := p.ReadByte()
		if err != nil {
			return nil, err
		}
		bytes = append(bytes, b)
	}
	return bytes, nil
}

func (p *BinaryParser) HasMore() bool {
	return len(p.data) != 0
}

func (p *BinaryParser) ReadVariableLength() (int, error) {
	b1, err := p.ReadByte()
	if err != nil {
		return 0, err
	}
	if b1 < 193 {
		return int(b1), nil
	}
	if b1 > 192 && b1 < 241 {
		b2, err := p.ReadByte()
		if err != nil {
			return 0, err
		}
		return 193 + ((int(b1) - 193) * 256) + int(b2), nil
	}
	if b1 > 240 && b1 < 255 {
		b2, err := p.ReadByte()
		if err != nil {
			return 0, err
		}
		b3, err := p.ReadByte()
		if err != nil {
			return 0, err
		}
		return 12481 + ((int(b1) - 241) * 65536) + (int(b2) * 256) + int(b3), nil
	}
	return 0, nil
}

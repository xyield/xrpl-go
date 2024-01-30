package serdes

import (
	"errors"

	"github.com/CreatureDev/xrpl-go/binary-codec/definitions"
)

var ErrLengthPrefixTooLong = errors.New("length of value must not exceed 918744 bytes of data")

type binarySerializer struct {
	sink []byte
}

func NewSerializer() *binarySerializer {
	return &binarySerializer{}
}

func (s *binarySerializer) put(v []byte) {
	s.sink = append(s.sink, v...)
}

func (s *binarySerializer) GetSink() []byte {
	return s.sink
}

func (s *binarySerializer) WriteFieldAndValue(fi definitions.FieldInstance, value []byte) error {
	h, err := encodeFieldID(fi.FieldName)

	if err != nil {
		return err
	}

	s.put(h)

	if fi.IsVLEncoded {
		vl, err := encodeVariableLength(len(value))
		if err != nil {
			return err
		}
		s.put(vl)
	}

	s.put(value)

	if fi.Type == "STObject" {
		s.put([]byte{0xE1})
	}
	return nil
}

func encodeVariableLength(len int) ([]byte, error) {
	if len <= 192 {
		return []byte{byte(len)}, nil
	}
	if len < 12480 {
		len -= 193
		b1 := byte((len >> 8) + 193)
		b2 := byte((len & 0xFF))
		return []byte{b1, b2}, nil
	}
	if len <= 918744 {
		len -= 12481
		b1 := byte((len >> 16) + 241)
		b2 := byte((len >> 8) & 0xFF)
		b3 := byte(len & 0xFF)
		return []byte{b1, b2, b3}, nil
	}
	return nil, ErrLengthPrefixTooLong
}

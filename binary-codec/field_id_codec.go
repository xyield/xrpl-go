package binarycodec

import (
	"encoding/hex"

	"github.com/xyield/xrpl-go/binary-codec/definitions"
)

func Encode(fieldName string) ([]byte, error) {
	fh, err := definitions.Get().GetFieldHeaderByFieldName(fieldName)
	if err != nil {
		return nil, err
	}
	var b []byte
	if fh.TypeCode < 16 && fh.FieldCode < 16 {
		return append(b, (byte(fh.TypeCode<<4))|byte(fh.FieldCode)), nil
	}
	if fh.TypeCode >= 16 && fh.FieldCode < 16 {
		return append(b, (byte(fh.FieldCode)), byte(fh.TypeCode)), nil
	}
	if fh.TypeCode < 16 && fh.FieldCode >= 16 {
		return append(b, byte(fh.TypeCode<<4), byte(fh.FieldCode)), nil
	}
	if fh.TypeCode >= 16 && fh.FieldCode >= 16 {
		return append(b, 0, byte(fh.TypeCode), byte(fh.FieldCode)), nil
	}
	return nil, nil
}

func Decode(h string) (string, error) {
	b, err := hex.DecodeString(h)
	if err != nil {
		return "", err
	}
	if len(b) == 1 {
		return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(b[0]>>4), int32(b[0]&byte(15))))
	}
	if len(b) == 2 {
		firstByteHighBits := b[0] >> 4
		firstByteLowBits := b[0] & byte(15)
		if firstByteHighBits == 0 {
			return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(b[1]), int32(firstByteLowBits)))
		}
		return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(firstByteHighBits), int32(b[1])))
	}
	if len(b) == 3 {
		return definitions.Get().GetFieldNameByFieldHeader(definitions.CreateFieldHeader(int32(b[1]), int32(b[2])))
	}
	return "", nil
}

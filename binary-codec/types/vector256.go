package types

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/CreatureDev/xrpl-go/binary-codec/serdes"
	"github.com/CreatureDev/xrpl-go/model/transactions/types"
)

const HashLengthBytes = 32

// ErrInvalidVector256Type represents an error when a Vector256 is constructed from an unexpected type.
type ErrInvalidVector256Type struct {
	Got string
}

// Error implements the error interface, providing a descriptive error message for ErrInvalidVector256Type.
func (e *ErrInvalidVector256Type) Error() string {
	return fmt.Sprintf("Invalid type to construct Vector256 from. Expected []string, got %v", e.Got)
}

// Vector256 represents a 256 bit vector.
type Vector256 struct{}

// FromJson converts a JSON value into a serialized byte slice representing a Vector256.
// The input value is assumed to be an array of strings representing Hash256 values.
// If the serialization fails, an error is returned.
func (v *Vector256) FromJson(json any) ([]byte, error) {
	switch json := json.(type) {
	case []string:
		return vector256FromValue(json)
	case []types.Hash256:
		var values []string
		for _, hash := range json {
			values = append(values, string(hash))
		}
		return vector256FromValue([]string(values))
	default:
		return nil, &ErrInvalidVector256Type{fmt.Sprintf("%T", json)}
	}
}

// vector256FromValue takes a slice of strings representing Hash256 values,
// serializes them, and returns the combined byte slice. If an error occurs during serialization, it is returned.
func vector256FromValue(value []string) ([]byte, error) {
	b := make([]byte, 0)
	for _, s := range value {
		hash256, err := NewHash256().FromJson(s)

		if err != nil {
			return nil, err
		}

		b = append(b, hash256...)

	}
	return b, nil
}

// ToJson takes a BinaryParser and optional parameters, and converts the serialized byte data
// back into an array of JSON string values representing Hash256 values.
// If the parsing fails, an error is returned.
func (v *Vector256) ToJson(p *serdes.BinaryParser, opts ...int) (any, error) {

	b, _ := p.ReadBytes(opts[0])
	var value []string

	for i := 0; i < len(b); i += HashLengthBytes {
		value = append(value, strings.ToUpper(hex.EncodeToString(b[i:i+HashLengthBytes])))
	}

	return value, nil
}

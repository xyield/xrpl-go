package types

import (
	"errors"
	"fmt"

	addresscodec "github.com/xyield/xrpl-go/address-codec"
	"github.com/xyield/xrpl-go/binary-codec/serdes"
)

const (
	typeAccount  = 0x01
	typeCurrency = 0x10
	typeIssuer   = 0x20

	pathsetEndByte    = 0x00
	pathSeparatorByte = 0xFF
)

// PathSet type declaration
type PathSet struct{}

// ErrInvalidPathSet is an error that's thrown when an invalid path set is provided.
var ErrInvalidPathSet error = errors.New("invalid type to construct PathSet from. Expected []any of []any")

// FromJson attempts to serialize a path set from a JSON representation of a slice of paths to a byte array.
// It returns the byte array representation of the path set, or an error if the provided json does not represent a valid path set.
func (p PathSet) FromJson(json any) ([]byte, error) {

	if _, ok := json.([]any)[0].([]any); !ok {
		return nil, ErrInvalidPathSet
	}

	if !isPathSet(json.([]any)) {
		return nil, ErrInvalidPathSet
	}

	return newPathSet(json.([]any)), nil
}

// ToJson decodes a path set from a binary representation using a provided binary parser, then translates it to a JSON representation.
// It returns a slice representing the JSON format of the path set, or an error if the path set could not be decoded or if an invalid step is encountered.
func (p PathSet) ToJson(parser *serdes.BinaryParser, opts ...int) (any, error) {
	var pathSet []any

	for parser.HasMore() {
		path, err := parsePath(parser)
		if err != nil {
			return nil, err
		}

		if len(path) > 0 {
			for i, step := range path {
				stepMap, ok := step.(map[string]any)
				if !ok {
					return nil, fmt.Errorf("step is not of type map[string]any")
				}
				if _, ok := stepMap["account"]; ok {
					stepMap["type"] = 1
					stepMap["type_hex"] = "0000000000000001"
				}
				if _, ok := stepMap["currency"]; ok {
					stepMap["type"] = 16
					stepMap["type_hex"] = "0000000000000010"
				}
				path[i] = stepMap
			}
			pathSet = append(pathSet, path)
		}
	}

	return pathSet, nil
}

// isPathSet determines if an array represents a valid path set.
// It checks if the array is either empty or if its first element is a valid path step.
func isPathSet(v []any) bool {
	return len(v) == 0 || len(v[0].([]any)) == 0 || isPathStep(v[0].([]any)[0].(map[string]any))
}

// isPathStep determines if a map represents a valid path step.
// It checks if any of the keys "account", "currency" or "issuer" are present in the map.
func isPathStep(v map[string]any) bool {
	return v["account"] != nil || v["currency"] != nil || v["issuer"] != nil
}

// newPathStep creates a path step from a map representation.
// It generates a byte array representation of the path step, encoding account, currency, and issuer information as appropriate.
func newPathStep(v map[string]any) []byte {

	dataType := 0x00
	b := make([]byte, 0)

	if v["account"] != nil {
		_, account, _ := addresscodec.DecodeClassicAddressToAccountID(v["account"].(string))
		b = append(b, account...)
		dataType |= typeAccount
	}
	if v["currency"] != nil {
		currency, _ := serializeIssuedCurrencyCode(v["currency"].(string))
		b = append(b, currency...)
		dataType |= typeCurrency
	}
	if v["issuer"] != nil {
		_, issuer, _ := addresscodec.DecodeClassicAddressToAccountID(v["issuer"].(string))
		b = append(b, issuer...)
		dataType |= typeIssuer
	}

	return append([]byte{byte(dataType)}, b...)
}

// newPath constructs a path from a slice of path steps.
// It generates a byte array representation of the path, encoding each path step in turn.
func newPath(v []any) []byte {
	b := make([]byte, 0)

	for _, step := range v { // for each step in the path (slice of path steps)
		b = append(b, newPathStep(step.(map[string]any))...) // append the path step to the byte array
	}
	return b
}

// newPathSet constructs a path set from a slice of paths.
// It generates a byte array representation of the path set, encoding each path and adding padding and path separators as appropriate.
func newPathSet(v []any) []byte {

	b := make([]byte, 0)
	padding := make([]byte, 20)

	for _, path := range v { // for each path in the path set (slice of paths)
		b = append(b, newPath(path.([]any))...) // append the path to the byte array
		b = append(b, padding...)               // append 20 empty bytes to the byte array between paths
		b = append(b, pathSeparatorByte)        // between each path, append a path separator byte
	}

	b[len(b)-1] = pathsetEndByte // replace last path separator with path set end byte

	return b

}

// parsePathStep decodes a path step from a binary representation using a provided binary parser.
// It returns a map representing the path step, or an error if the path step could not be decoded.
func parsePathStep(parser *serdes.BinaryParser) (map[string]any, error) {
	dataType, err := parser.ReadByte()
	if err != nil {
		return nil, err
	}

	step := make(map[string]any)

	operations := []struct {
		typeKey byte
		key     string
	}{
		{typeAccount, "account"},
		{typeCurrency, "currency"},
		{typeIssuer, "issuer"},
	}

	for _, op := range operations {
		if dataType&op.typeKey != 0 {
			bytes, err := parser.ReadBytes(20) // AccountID or Currency size
			if err != nil {
				return nil, err
			}

			if op.typeKey == typeCurrency {
				value, err := deserialiseCurrencyCode(bytes)
				if err != nil {
					return nil, err
				}
				step[op.key] = value
			} else {
				value := addresscodec.Encode(bytes, []byte{addresscodec.AccountAddressPrefix}, addresscodec.AccountAddressLength)
				step[op.key] = value
			}
		}
	}

	return step, nil
}

// parsePath decodes a path from a binary representation using a provided binary parser.
// It returns a slice representing the path, or an error if the path could not be decoded.
func parsePath(parser *serdes.BinaryParser) ([]any, error) {
	var path []any

	for parser.HasMore() {
		peek, err := parser.Peek()
		if err != nil {
			return nil, err
		}

		if peek == pathsetEndByte || peek == pathSeparatorByte {
			_, err := parser.ReadByte()
			if err != nil {
				return nil, err
			}
			break
		}

		step, err := parsePathStep(parser)
		if err != nil {
			return nil, err
		}
		path = append(path, step)
	}

	return path, nil
}

// parsePathSet decodes a path set from a binary representation using a provided binary parser.
// It returns a slice representing the path set, or an error if the path set could not be decoded.
func parsePathSet(parser *serdes.BinaryParser) ([]any, error) {
	var pathSet []any

	for parser.HasMore() {
		path, err := parsePath(parser)
		if err != nil {
			return nil, err
		}
		if len(path) > 0 {
			pathSet = append(pathSet, path)
		}
	}

	return pathSet, nil
}

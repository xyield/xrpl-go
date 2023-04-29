package types

import (
	"errors"

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

type PathSet struct{}

var ErrInvalidPathSet error = errors.New("invalid type to construct PathSet from. Expected []any of []any")

// Serializes a path set from a json representation of a slice of paths to a byte array
func (p PathSet) FromJson(json any) ([]byte, error) {

	if _, ok := json.([]any)[0].([]any); !ok {
		return nil, ErrInvalidPathSet
	}

	if !isPathSet(json.([]any)) {
		return nil, ErrInvalidPathSet
	}

	return newPathSet(json.([]any)), nil
}

func (p PathSet) FromParser(parser *serdes.BinaryParser, opts ...int) (any, error) {
	return nil, nil
}

// determine if an array represents a valid path set
func isPathSet(v []any) bool {
	return len(v) == 0 || len(v[0].([]any)) == 0 || isPathStep(v[0].([]any)[0].(map[string]any))
}

// determine if a map represents a valid path step
func isPathStep(v map[string]any) bool {
	return v["account"] != nil || v["currency"] != nil || v["issuer"] != nil
}

// creates a path step from a map representation
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

// constructs a path from a slice of path steps
func newPath(v []any) []byte {
	b := make([]byte, 0)

	for _, step := range v { // for each step in the path (slice of path steps)
		b = append(b, newPathStep(step.(map[string]any))...) // append the path step to the byte array
	}
	return b
}

// constructs a path set from a slice of paths
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

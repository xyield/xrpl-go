package types

import (
	"errors"

	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

const (
	typeAccount  = 0x01
	typeCurrency = 0x10
	typeIssuer   = 0x20

	pathsetEndByte    = 0x00
	pathSeparatorByte = 0xFF
)

type PathSet []byte

type PathStep []byte

type Path []byte

// determine if an array represents a valid path set
func isPathSet(v [][]map[string]string) bool {
	return len(v) == 0 || len(v[0]) == 0 || isPathStep(v[0][0])
}

// determine if a map represents a valid path step
func isPathStep(v map[string]string) bool {
	return v["account"] != "" || v["currency"] != "" || v["issuer"] != ""
}

// creates a path step from a map representation
func newPathStep(v map[string]string) PathStep {

	dataType := 0x00
	b := make([]byte, 0)

	if v["account"] != "" {
		_, account, _ := addresscodec.DecodeClassicAddressToAccountID(v["account"])
		b = append(b, account...)
		dataType |= typeAccount
	}
	if v["currency"] != "" {
		currency, _ := SerializeIssuedCurrencyCode(v["currency"])
		b = append(b, currency...)
		dataType |= typeCurrency
	}
	if v["issuer"] != "" {
		_, issuer, _ := addresscodec.DecodeClassicAddressToAccountID(v["issuer"])
		b = append(b, issuer...)
		dataType |= typeIssuer
	}

	return append([]byte{byte(dataType)}, b...)
}

// constructs a path from a slice of path steps
func newPath(v []map[string]string) Path {
	b := make([]byte, 0)
	for _, step := range v {
		b = append(b, newPathStep(step)...)
	}
	return b
}

func newPathSet(v [][]map[string]string) (PathSet, error) {

	if !isPathSet(v) {
		return nil, errors.New("invalid path set")
	}

	b := make([]byte, 0)

	for _, path := range v {
		for _, step := range path {
			b = append(b, newPathStep(step)...)
		}
		b = append(b, pathSeparatorByte)
	}

	return append(b, pathsetEndByte), nil

}

package keypairs

import (
	"crypto/rand"

	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

var r randomizer

func init() {
	r.Reader = rand.Reader
}

type CryptoImplementation interface {
	deriveKeypair(decodedSeed string, validator bool) (string, string, error)
}

func GenerateSeed(entropy string, alg addresscodec.CryptoAlgorithm) (string, error) {
	var pe []byte
	if entropy == "" {
		b, err := r.generateBytes(addresscodec.FamilySeedLength)
		pe = b
		if err != nil {
			return "", err
		}
	} else {
		pe = []byte(entropy)[:addresscodec.FamilySeedLength]
	}
	return addresscodec.EncodeSeed(pe, alg)
}

// func DeriveKeypair(seed string)

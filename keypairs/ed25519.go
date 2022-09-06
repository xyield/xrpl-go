package keypairs

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"strings"

	addresscodec "github.com/xyield/xrpl-go/address-codec"
)

type ed25519Alg struct{}

func (c *ed25519Alg) deriveKeypair(decodedSeed []byte, validator bool) (string, string, error) {
	rawPriv := sha512Half(decodedSeed)
	pubKey, privKey, err := ed25519.GenerateKey(bytes.NewBuffer(rawPriv))
	if err != nil {
		return "", "", err
	}
	pubKey = append([]byte{addresscodec.ED25519Prefix}, pubKey...)
	public := strings.ToUpper(hex.EncodeToString(pubKey))
	privKey = append([]byte{addresscodec.ED25519Prefix}, privKey...)
	private := strings.ToUpper(hex.EncodeToString(privKey[:32+len([]byte{addresscodec.ED25519Prefix})]))
	return private, public, nil
}

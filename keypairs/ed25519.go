package keypairs

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"

	addresscodec "github.com/CreatureDev/xrpl-go/address-codec"
)

type ed25519Alg struct{}

func (c *ed25519Alg) deriveKeypair(decodedSeed []byte, validator bool) (string, string, error) {
	if validator {
		return "", "", &ed25519ValidatorError{}
	}
	rawPriv := sha512Half(decodedSeed)
	pubKey, privKey, err := ed25519.GenerateKey(bytes.NewBuffer(rawPriv))
	if err != nil {
		return "", "", err
	}
	pubKey = append([]byte{addresscodec.ED25519}, pubKey...)
	public := formatKey(pubKey)
	privKey = append([]byte{addresscodec.ED25519}, privKey...)
	private := formatKey(privKey[:32+len([]byte{addresscodec.ED25519})])
	return private, public, nil
}

func (c *ed25519Alg) sign(msg, privKey string) (string, error) {
	b, err := hex.DecodeString(privKey)
	if err != nil {
		return "", err
	}
	rawPriv := ed25519.NewKeyFromSeed(b[1:])
	signedMsg := ed25519.Sign(rawPriv, []byte(msg))
	return formatKey(signedMsg), nil
}

func (c *ed25519Alg) validate(msg, pubkey, sig string) bool {
	return ed25519.Verify(ed25519.PublicKey(deformatKey(pubkey)[1:]), []byte(msg), deformatKey(sig))
}

type ed25519ValidatorError struct{}

func (e *ed25519ValidatorError) Error() string {
	return "validator keypairs can not use Ed25519"
}

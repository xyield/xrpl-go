package addresscodec

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"

	"golang.org/x/crypto/ripemd160"
)

const (
	AccountAddressLength   = 20
	AccountPublicKeyLength = 33
	FamilySeedLength       = 16
	NodePublicKeyLength    = 33

	AccountAddressPrefix   = 0x00
	AccountPublicKeyPrefix = 0x23
	FamilySeedPrefix       = 0x21
	NodePublicKeyPrefix    = 0x1C

	ED25519Prefix          = 0xED
	ED25519PrefixHexString = "ED"
)

type EncodeLengthError struct {
	Instance string
	Input    int
	Expected int
}

func (e *EncodeLengthError) Error() string {
	return fmt.Sprintf("`%v` length should be %v not %v", e.Instance, e.Expected, e.Input)
}

type InvalidClassicAddressError struct {
	Input string
}

func (e *InvalidClassicAddressError) Error() string {
	return fmt.Sprintf("`%v` is an invalid classic address", e.Input)
}

func EncodeClassicAddressFromPublicKeyHex(pubkeyhex string, typePrefix []byte) (string, error) {

	if len(typePrefix) != 1 {
		return "", &EncodeLengthError{Instance: "TypePrefix", Expected: 1, Input: len(typePrefix)}
	}

	pubkey, err := hex.DecodeString(pubkeyhex)

	if err != nil {
		return "", &EncodeLengthError{Instance: "PublicKey", Expected: AccountPublicKeyLength, Input: len(pubkey)}
	}

	accountID := sha256RipeMD160(pubkey)

	if len(accountID) != AccountAddressLength {
		return "", &EncodeLengthError{Instance: "AccountID", Expected: AccountAddressLength, Input: len(accountID)}
	}

	checkSum := createCheckSum(append(typePrefix, accountID...))[:4]

	if len(checkSum) != 4 {
		return "", &EncodeLengthError{Instance: "CheckSum", Expected: 4, Input: len(checkSum)}
	}

	payload := append(typePrefix, accountID...)
	address := EncodeBase58((append(payload, checkSum...)))

	if !IsValidClassicAddress(address) {
		return "", &InvalidClassicAddressError{Input: address}
	}

	return address, nil
}

func DecodeClassicAddressToAccountID(cAddress string) (typePrefix, accountID []byte, err error) {

	if len(DecodeBase58(cAddress)) != 25 {
		return nil, nil, &InvalidClassicAddressError{Input: cAddress}
	}

	return DecodeBase58(cAddress)[:1], DecodeBase58(cAddress)[1:21], nil

}

func IsValidClassicAddress(cAddress string) bool {
	_, _, c := DecodeClassicAddressToAccountID(cAddress)

	return c == nil
}

func EncodeNodePublicKey(pubkeyhex string, typePrefix []byte) (string, error) {
	return "", nil
}

func EncodeSeed(entropy string, versionType hash.Hash) (string, error) {
	return "", nil
}

func DecodeSeed(seed string) (string, error) {
	return "", nil
}

func sha256RipeMD160(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	ripemd160 := ripemd160.New()
	ripemd160.Write(sha256.Sum(nil))

	return ripemd160.Sum(nil)
}

func createCheckSum(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	sha256sha256 := sha256.Sum(nil)
	sha256.Reset()
	sha256.Write(sha256sha256)

	return sha256.Sum(nil)
}

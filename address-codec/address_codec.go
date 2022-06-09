package addresscodec

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"

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

	ED25519Prefix = 0xED
)

type CryptoAlgorithm uint32

const (
	Undefined CryptoAlgorithm = iota
	ED25519
	SECP256K1
)

func (c CryptoAlgorithm) String() string {
	switch c {
	case ED25519:
		return "ed25519"
	case SECP256K1:
		return "secp256k1"
	}
	return "unknown"
}

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

func Encode(b []byte, typePrefix []byte, expectedLength int) string {

	if len(b) != expectedLength {
		return ""
	}

	return CheckEncode(b, typePrefix[0])
}

func Decode(b58string string, typePrefix []byte) []byte {

	prefixLength := len(typePrefix)

	if !bytes.Equal(DecodeBase58(b58string)[:prefixLength], typePrefix) {
		return nil
	}

	return DecodeBase58(b58string)[prefixLength:]
}

func EncodeClassicAddressFromPublicKeyHex(pubkeyhex string, typePrefix []byte) (string, error) {

	if len(typePrefix) != 1 {
		return "", &EncodeLengthError{Instance: "TypePrefix", Expected: 1, Input: len(typePrefix)}
	}

	pubkey, err := hex.DecodeString(pubkeyhex)

	if len(pubkey) != AccountPublicKeyLength {
		pubkey = append([]byte{ED25519Prefix}, pubkey...)
	}

	if err != nil {
		return "", &EncodeLengthError{Instance: "PublicKey", Expected: AccountPublicKeyLength, Input: len(pubkey)}
	}

	accountID := sha256RipeMD160(pubkey)

	if len(accountID) != AccountAddressLength {
		return "", &EncodeLengthError{Instance: "AccountID", Expected: AccountAddressLength, Input: len(accountID)}
	}

	address := CheckEncode(accountID, AccountAddressPrefix)

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

func EncodeSeed(entropy []byte, encodingType CryptoAlgorithm) (string, error) {

	if len(entropy) != FamilySeedLength {
		return "", &EncodeLengthError{Instance: "Entropy", Input: len(entropy), Expected: FamilySeedLength}
	}

	switch encodingType {
	case ED25519:
		prefix := []byte{ED25519Prefix}
		return Encode(entropy, prefix, FamilySeedLength), nil
	case SECP256K1:
		prefix := []byte{FamilySeedLength}
		return Encode(entropy, prefix, FamilySeedLength), nil
	default:
		return "", errors.New("encoding type must be `ed25519` or `secp256k1`")
	}

}

func DecodeSeed(seed string) ([]byte, CryptoAlgorithm, error) {

	decodedResult := Decode(seed, []byte{ED25519Prefix})

	return decodedResult, ED25519, nil
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

package addresscodec

import (
	"bytes"
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
	ED25519PrefixHexString = "ED" //hex prefix for 32 byte ED25519 hex strings, to make length 33 bytes after decoding
)

type EncodeLengthError struct {
	Instance string
	Input    int
}

func (e *EncodeLengthError) Error() string {
	return fmt.Sprintf("%v length should be %v", e.Instance, e.Input)
}

func EncodeClassicAddressFromPublicKeyHex(pubkeyhex string, typePrefix []byte) (string, error) {

	pubkey, _ := hex.DecodeString(pubkeyhex)

	if len(pubkey) != AccountPublicKeyLength {
		return "", &EncodeLengthError{Instance: "PublicKey", Input: AccountPublicKeyLength}
	}

	accountID := sha256RipeMD160(pubkey)

	if len(accountID) != AccountAddressLength {
		return "", &EncodeLengthError{Instance: "AccountID", Input: AccountAddressLength}
	}

	payload := append(typePrefix, accountID...)

	if len(payload) != 21 {
		return "", &EncodeLengthError{Instance: "Payload", Input: 21}
	}

	checkSum := createCheckSum(payload)[:4]

	if len(checkSum) != 4 {
		return "", &EncodeLengthError{Instance: "CheckSum", Input: 4}
	}

	address := EncodeBase58((append(payload, checkSum...)))

	if len(address) != 34 { //can they be different lengths?
		return "", &EncodeLengthError{Instance: "Address", Input: 34}
	}

	if !bytes.Equal(accountID, DecodeBase58(address)[1:21]) {
		return "", &EncodeLengthError{Instance: "DecodedAddress", Input: 20}
	}

	return address, nil
}

func DecodeClassicAddressToAccountID(cAddress string) (typePrefix, accountID []byte, err error) {

	if len(DecodeBase58(cAddress)[1:21]) != AccountAddressLength {
		return nil, nil, &EncodeLengthError{Instance: "DecodedAddress", Input: 20}
	}

	return DecodeBase58(cAddress)[:1], DecodeBase58(cAddress)[1:21], nil

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

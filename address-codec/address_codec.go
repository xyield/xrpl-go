package addresscodec

import (
	"crypto/sha256"
	"encoding/hex"
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

var (
	ED25519PrefixByteSlice, _ = hex.DecodeString(ED25519PrefixHexString)
)

func Sha256RipeMD160(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	ripemd160 := ripemd160.New()
	ripemd160.Write(sha256.Sum(nil))

	return ripemd160.Sum(nil)
}

func CreateCheckSum(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	sha256sha256 := sha256.Sum(nil)
	sha256.Reset()
	sha256.Write(sha256sha256)

	return sha256.Sum(nil)
}

func EncodeAddressFromPublicKeyHex(pubkeyhex string, typePrefix []byte) (string, error) {

	pubkey, _ := hex.DecodeString(pubkeyhex)

	if len(pubkey) != AccountPublicKeyLength {
		return "", &EncodeLengthError{Instance: "PublicKey", Input: AccountPublicKeyLength}
	}

	accountID := Sha256RipeMD160(pubkey)

	if len(accountID) != AccountAddressLength {
		return "", &EncodeLengthError{Instance: "AccountID", Input: AccountAddressLength}
	}

	payload := append(typePrefix, accountID...)

	if len(payload) != 21 {
		return "", &EncodeLengthError{Instance: "Payload", Input: 21}
	}

	checkSum := CreateCheckSum(payload)[:4]

	if len(checkSum) != 4 {
		return "", &EncodeLengthError{Instance: "CheckSum", Input: 4}
	}

	address := EncodeBase58((append(payload, checkSum...)))

	if len(address) != 34 { //can they be different lengths?
		return "", &EncodeLengthError{Instance: "Address", Input: 34}
	}

	return address, nil
}

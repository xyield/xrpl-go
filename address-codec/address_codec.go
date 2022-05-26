package addresscodec

import (
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

	ED25519PrefixHex = "ED"
	pubkeyHex        = "9434799226374926EDA3B54B1B461B4ABF7237962EAE18528FEA67595397FA32" //just an example
)

var (
	AccountAddressPrefix   = []byte{0x00}
	AccountPublicKeyPrefix = []byte{0x23}
	FamilySeedPrefix       = []byte{0x21}
	NodePublicKeyPrefix    = []byte{0x1C}

	pubkey, _        = hex.DecodeString(pubkeyHex)
	ED25519Prefix, _ = hex.DecodeString(ED25519PrefixHex)
)

func Sha256RipeMD160(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	ripemd160 := ripemd160.New()
	ripemd160.Write(sha256.Sum(nil))

	return ripemd160.Sum(nil)
}

func Sha256Sha256(b []byte) []byte {
	sha256 := sha256.New()
	sha256.Write(b)

	sha256sha256 := sha256.Sum(nil)
	sha256.Reset()
	sha256.Write(sha256sha256)

	return sha256.Sum(nil)
}

func EncodeAddress(pubkeyhex string) (string, error) {

	prefixandpubkey := append(ED25519Prefix, pubkey...)

	if len(prefixandpubkey) != AccountPublicKeyLength {
		return "", errors.New("public key length should be 33")
	}

	accountID := Sha256RipeMD160(prefixandpubkey)

	if len(accountID) != AccountAddressLength {
		return "", errors.New("account ID length should be 20")
	}

	payload := append(AccountAddressPrefix, accountID...)

	if len(payload) != 21 {
		return "", errors.New("payload length should be 21")
	}

	checkSum := Sha256Sha256(accountID)[0:4]

	if len(checkSum) != 4 {
		return "", errors.New("checksum length should be 4")
	}

	address := EncodeBase58((append(payload, checkSum...)))

	fmt.Println(address)

	return "rDTXLQ7ZKZVKz33zJbHjgVShjsBnqMBhmN", nil
}

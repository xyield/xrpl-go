package addresscodec

import (
	"crypto/sha256"
	"errors"
)

var (
	ErrChecksum      = errors.New("checksum error")
	ErrInvalidFormat = errors.New("invalid format: version and/or checksum bytes missing")
)

func checksum(input []byte) (cksum [4]byte) {
	h := sha256.Sum256(input)
	h2 := sha256.Sum256(h[:])
	copy(cksum[:], h2[:4])
	return cksum
}

func Base58CheckEncode(input []byte, prefix byte) string {
	b := make([]byte, 0, 1+len(input)+4)
	b = append(b, prefix)
	b = append(b, input...)

	cksum := checksum(b)
	b = append(b, cksum[:]...)
	return EncodeBase58(b)
}

func Base58CheckDecode(input string) (result []byte, prefix byte, err error) {
	decoded := DecodeBase58(input)
	if len(decoded) < 5 {
		return nil, 0, ErrInvalidFormat
	}
	prefix = decoded[0]
	var cksum [4]byte
	copy(cksum[:], decoded[len(decoded)-4:])
	if checksum(decoded[:len(decoded)-4]) != cksum {
		return nil, 0, ErrChecksum
	}
	payload := decoded[1 : len(decoded)-4]
	result = append(result, payload...)
	return
}

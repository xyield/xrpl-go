package types

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidHashType = errors.New("invalid type for Hash, expected string")
)

// ErrInvalidHashLength struct is used when the hash length does not meet the expected value.
type ErrInvalidHashLength struct {
	Expected int
}

// Error method for ErrInvalidHashLength formats the error message.
func (e *ErrInvalidHashLength) Error() string {
	return fmt.Sprintf("invalid hash length expected length %v", e.Expected)
}

// hashI interface combines the SerializedType interface and getLength method for hashes.
type hashI interface {
	SerializedType
	getLength() int
}

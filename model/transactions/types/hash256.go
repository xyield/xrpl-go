package types

import "fmt"

type Hash256 string

func (h Hash256) Validate() error {
	if h == "" {
		return fmt.Errorf("hash256 value not set")
	}
	if len(h) != 64 {
		return fmt.Errorf("hash256 length was not expected 64 characters")
	}
	return nil
}

package types

import (
	"fmt"
	"strings"
)

type Address string

func (a Address) Validate() error {
	characters := []string{"0", "O", "I", "l"}
	if len(a) == 0 {
		return fmt.Errorf("missing xrpl address")
	}
	if len(a) < 25 || len(a) > 35 {
		return fmt.Errorf("invalid xrpl address length")
	}
	if a[0] != 'r' {
		return fmt.Errorf("invalid xrpl address prefix '%c'", a[0])
	}
	for _, c := range characters {
		if strings.Contains(string(a), c) {
			return fmt.Errorf("xrpl address contains invalid character '%s'", c)
		}
	}
	// TODO checksum
	return nil
}

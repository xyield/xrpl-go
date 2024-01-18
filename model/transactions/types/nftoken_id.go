package types

import "fmt"

type NFTokenID Hash256

func (id NFTokenID) Validate() error {
	h := Hash256(id)
	if err := h.Validate(); err != nil {
		return fmt.Errorf("nftoken id: %w", err)
	}
	return nil
}

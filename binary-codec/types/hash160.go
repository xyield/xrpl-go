package types

// Hash160 struct represents a 160-bit hash.
type Hash160 struct {
	hashI
}

// NewHash160 is a constructor for creating a new 160-bit hash.
func NewHash160() *Hash160 {
	return &Hash160{
		newHash(20),
	}
}

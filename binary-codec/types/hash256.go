package types

// Hash256 struct represents a 256-bit hash.
type Hash256 struct {
	hashI
}

// NewHash256 is a constructor for creating a new 256-bit hash.
func NewHash256() *Hash256 {
	return &Hash256{
		newHash(32),
	}
}

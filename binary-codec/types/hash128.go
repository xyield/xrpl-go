package types

// Hash128 struct represents a 128-bit hash.
type Hash128 struct {
	hashI
}

// NewHash128 is a constructor for creating a new 128-bit hash.
func NewHash128() *Hash128 {
	return &Hash128{
		newHash(16),
	}
}

package types

type Hash128 struct {
	hashI
}

func NewHash128() *Hash128 {
	return &Hash128{
		newHash(16),
	}
}

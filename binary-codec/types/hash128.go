package types

type Hash128 struct {
	hash
}

func (h *Hash128) getLength() int {
	return 16
}

func NewHash128() *Hash128 {
	return &Hash128{
		newHash(16),
	}
}

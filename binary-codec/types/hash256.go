package types

type Hash256 struct {
	hash
}

func NewHash256() *Hash256 {
	return &Hash256{
		newHash(32),
	}
}

func (h *Hash256) getLength() int {
	return 32
}

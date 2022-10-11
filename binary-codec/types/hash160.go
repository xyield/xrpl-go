package types

type Hash160 struct {
	hash
}

func (h *Hash160) getLength() int {
	return 20
}

func NewHash160() *Hash160 {
	return &Hash160{
		newHash(20),
	}
}

package types

type Hash256 struct {
	hashI
}

func NewHash256() *Hash256 {
	return &Hash256{
		newHash(32),
	}
}

package types

type Hash160 struct {
	hashI
}

func NewHash160() *Hash160 {
	return &Hash160{
		newHash(20),
	}
}

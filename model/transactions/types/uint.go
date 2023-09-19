package types

// Uint is a helper function that allocates a new uint value
// to store v and returns a pointer to it.
func Uint(v uint) *uint {
	p := new(uint)
	*p = v
	return p
}

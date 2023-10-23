package types

// FlagsI is an interface for types that can be converted to a uint.
type FlagsI interface {
	ToUint() uint32
}

type Flag uint32

func (f *Flag) ToUint() uint32 {
	return uint32(*f)
}

// SetFlag is a helper function that allocates a new uint value
// to store v and returns a pointer to it.
func SetFlag(v uint32) *Flag {
	p := new(uint32)
	*p = v
	return (*Flag)(p)
}

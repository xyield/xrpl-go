package types

import "hash"

func CalcHash(buf []byte, hasher hash.Hash) []byte {
	_, _ = hasher.Write(buf)
	return hasher.Sum(nil)
}

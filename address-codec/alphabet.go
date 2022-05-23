package addresscodec

type Alphabet struct {
	decode [128]int8
	encode [58]byte
}

func NewAlphabet(s string) *Alphabet {
	if len(s) != 58 {
		panic("base58 alphabet must be 58 bytes long")
	}
	ret := new(Alphabet)
	copy(ret.encode[:], s)
	for i := range ret.decode {
		ret.decode[i] = -1
	}

	distinct := 0
	for i, b := range ret.encode {
		if ret.decode[b] == -1 {
			distinct++
		}
		ret.decode[b] = int8(i)
	}

	if distinct != 58 {
		panic("provided alphabet does not consist of 58 distinct characters")
	}

	return ret
}

var BTCAlphabet = NewAlphabet("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")
var XRPAlphabet = NewAlphabet("rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz")

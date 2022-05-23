package addresscodec

import (
	"fmt"
	"math/big"
)

const (
	alphabet = "rpshnaf39wBUDNEGHJKLM4PQRST7VWXYZ2bcdeCg65jkm8oFqi1tuvAxyz"

	alphabetIdx0 = 1
)

func EncodeBase58(bin []byte) string {
	return FastBase58EncodingAlphabet(bin, XRPAlphabet)
}

func FastBase58EncodingAlphabet(bin []byte, alphabet *Alphabet) string {

	size := len(bin)

	zcount := 0
	for zcount < size && bin[zcount] == 0 {
		zcount++
	}

	size = zcount +
		(size-zcount)*555/406 + 1

	out := make([]byte, size)

	var i, high int
	var carry uint32

	high = size - 1

	for _, b := range bin {
		i = size - 1
		for carry = uint32(b); i > high || carry != 0; i-- {
			carry = carry + 256*uint32(out[i])
			out[i] = byte(carry % 58)
			carry /= 58
		}
		high = i
	}

	for i = zcount; i < size && out[i] == 0; i++ {
	}

	val := out[i-zcount:]
	size = len(val)
	for i = 0; i < size; i++ {
		out[i] = alphabet.encode[val[i]]
	}

	return string(out[:size])

}

var bigRadix = [...]*big.Int{
	big.NewInt(0),
	big.NewInt(58),
	big.NewInt(58 * 58),
	big.NewInt(58 * 58 * 58),
	big.NewInt(58 * 58 * 58 * 58),
	big.NewInt(58 * 58 * 58 * 58 * 58),
	big.NewInt(58 * 58 * 58 * 58 * 58 * 58),
	big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58),
	big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58 * 58),
	big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58 * 58 * 58),
	bigRadix10,
}

var bigRadix10 = big.NewInt(58 * 58 * 58 * 58 * 58 * 58 * 58 * 58 * 58 * 58)

func EncodeBase58Original(d []byte) string {

	bn := new(big.Int)
	bn.SetBytes(d)

	maxlen := int(float64(len(d))*1.365658237309761) + 1
	answer := make([]byte, 0, maxlen)
	mod := new(big.Int)

	for bn.Sign() > 0 {
		bn.DivMod(bn, bigRadix10, mod)

		if bn.Sign() == 0 {
			m := mod.Int64()
			for m > 0 {
				answer = append(answer, alphabet[m%58])
				m /= 58
			}
		} else {
			m := mod.Int64()
			for i := 0; i < 10; i++ {
				answer = append(answer, alphabet[m%58])
				m /= 58
			}
		}
	}

	for _, i := range d {
		if i != 0 {
			break
		}
		answer = append(answer, alphabetIdx0)
	}

	alen := len(answer)
	for i := 0; i < alen/2; i++ {
		answer[i], answer[alen-1-i] = answer[alen-1-i], answer[i]
	}

	return string(answer)

}

func DecodeBase58(str string) ([]byte, error) {
	return FastBase58DecodingAlphabet(str, XRPAlphabet)
}

func FastBase58DecodingAlphabet(str string, alphabet *Alphabet) ([]byte, error) {
	if len(str) == 0 {
		return nil, fmt.Errorf("zero length string")
	}

	zero := alphabet.encode[0]
	b58len := len(str)

	var zcount int
	for i := 0; i < b58len && str[i] == zero; i++ {
		zcount++
	}

	var t, c uint64

	binu := make([]byte, 2*((b58len*406/555)+1))
	outi := make([]uint32, (b58len+3)/4)

	for _, r := range str {
		if r > 127 {
			return nil, fmt.Errorf("high-bit set on invalid digit")
		}
		if alphabet.decode[r] == -1 {
			return nil, fmt.Errorf("invalid base58 digit (%q)", r)
		}

		c = uint64(alphabet.decode[r])

		for j := len(outi) - 1; j >= 0; j-- {
			t = uint64(outi[j])*58 + c
			c = t >> 32
			outi[j] = uint32(t & 0xffffffff)
		}
	}

	mask := (uint(b58len%4) * 8)
	if mask == 0 {
		mask = 32
	}
	mask -= 8

	outLen := 0
	for j := 0; j < len(outi); j++ {
		for mask < 32 {
			binu[outLen] = byte(outi[j] >> mask)
			mask -= 8
			outLen++
		}
		mask = 24
	}

	for msb := zcount; msb < len(binu); msb++ {
		if binu[msb] > 0 {
			return binu[msb-zcount : outLen], nil
		}
	}

	return binu[:outLen], nil
}

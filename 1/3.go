package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	enc, _ := hex.DecodeString("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736")
	done := false
	maxScore := -1
	res := "make([]byte, len(enc))"
	buf := make([]byte, len(enc))
	for i := byte(0); !done; {
		for pos, val := range enc {
			buf[pos] = val ^ i
		}
		score := score(buf)
		if score > maxScore {
			maxScore = score
			res = string(i) + ": " + string(buf)
		}
		if i < 255 {
			i += 1
		} else {
			done = true
		}
	}
	fmt.Println(res)
}

func score(in []byte) int {
	ret := 0
	freq := map[byte]int{}
	for _, x := range in {
		if val, ok := freq[x]; ok {
			freq[x] = val + 1
		} else {
			freq[x] = 1
		}
	}
	n := len(in)
	if freq['e']+freq['E'] > n/20 {
		ret += freq['e'] + freq['E']
	}
	if freq['o']+freq['O'] > n/20 {
		ret += freq['o'] + freq['O']
	}
	if freq['a']+freq['A'] > n/20 {
		ret += freq['a'] + freq['A']
	}
	if freq['i']+freq['I'] > n/20 {
		ret += freq['i'] + freq['I']
	}
	if _, ok := freq[' ']; !ok {
		ret -= 10
	}
	nonPrint := []byte{'&', '$', '%', '{', '}', '\n', '+', '=', '*', '^', '/', '\\', '`'}
	for _, bad := range nonPrint {
		if _, exists := freq[bad]; exists {
			ret -= 10
		}
	}

	return ret
}

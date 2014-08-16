package utils

import (
	"bufio"
	"encoding/hex"
	"os"
)

func Xor(left, right string) string {
	a, _ := hex.DecodeString(left)
	b, _ := hex.DecodeString(right)
	res := make([]byte, len(a))
	for i, _ := range a {
		res[i] = a[i] ^ b[i]
	}
	return hex.EncodeToString(res)
}

func DecodeHex(a string) ([]byte, error) {
	return hex.DecodeString(a)
}

func FindCipher(a string, atLeast int) []string {
	enc, _ := DecodeHex(a)
	done := false
	best := map[string]int{}
	bestCap := 3
	var minBest string
	var minBestScore int
	buf := make([]byte, len(enc))
	for i := byte(0); !done; {
		for pos, val := range enc {
			buf[pos] = val ^ i
		}
		score := score(buf)
		if score > atLeast {
			key := string(i) + ": '" + string(buf) + "'"
			if len(best) < bestCap {
				best[key] = score
			} else {
				if score > minBestScore {
					delete(best, minBest)
					best[key] = score
				}
				minBest, minBestScore = findMin(best)
			}
		}
		if i < 255 {
			i += 1
		} else {
			done = true
		}
	}
	res := []string{}
	for k := range best {
		res = append(res, k)
	}
	return res
}

func findMin(m map[string]int) (string, int) {
	var min string
	var minScore int
	started := false
	for k, v := range m {
		if !started {
			min = k
			minScore = v
			started = true
		}
		if v < minScore {
			min = k
			minScore = v
		}

	}
	return min, minScore
}

var letters string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func score(in []byte) int {
	freq := map[byte]int{}
	for _, x := range in {
		if val, ok := freq[x]; ok {
			freq[x] = val + 1
		} else {
			freq[x] = 1
		}
	}
	letterCount := 0
	for _, l := range []byte(letters) {
		letterCount += freq[l]
	}
	ret := letterCount
	ret += freq[' ']
	nonPrint := []byte{'&', '$', '%', '{', '}', '+', '=', '*', '^', '/', '\\', '@', '(', ')', '[', ']', '_', '#', '<', '>', '-'}
	for _, bad := range nonPrint {
		if _, exists := freq[bad]; exists {
			ret -= 10
		}
	}
	return ret
}

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Encrypt(message, key []byte) []byte {
	res := make([]byte, len(message))
	i := 0
	for pos, x := range message {
		res[pos] = x ^ key[i]
		i = (i + 1) % len(key)
	}
	return res
}

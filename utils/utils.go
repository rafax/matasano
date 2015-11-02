package utils

import (
	"bufio"
	"encoding/hex"
	"os"
	"sort"
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

func FindCipher(enc []byte, atLeast int) []KeyEncoding {
	done := false
	best := map[KeyEncoding]int{}
	bestCap := 3
	var minBest KeyEncoding
	var minBestScore int
	buf := make([]byte, len(enc))
	for i := byte(0); !done; {
		for pos, val := range enc {
			buf[pos] = val ^ i
		}
		score := Score(buf)
		if score > atLeast {
			key := KeyEncoding{Key: string(i), Encoding: string(buf), Score: score}
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
	res := []KeyEncoding{}
	for k := range best {
		res = append(res, k)
	}
	sort.Sort(sort.Reverse(KeyEncodings(res)))
	return res
}

type KeyEncoding struct {
	Key      string
	Encoding string
	Score    int
}

type KeyEncodings []KeyEncoding

func (a KeyEncodings) Len() int           { return len(a) }
func (a KeyEncodings) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KeyEncodings) Less(i, j int) bool { return a[i].Score < a[j].Score }

func findMin(m map[KeyEncoding]int) (KeyEncoding, int) {
	var min KeyEncoding
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

func Score(in []byte) int {
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

func XorEncrypt(message, key []byte) []byte {
	res := make([]byte, len(message))
	for pos, x := range message {
		res[pos] = x ^ key[pos%len(key)]
	}
	return res
}

func HammingDistance(a, b []byte) int {
	cnt := 0
	for pos := range a {
		for i := uint(0); i < 8; i++ {
			if a[pos]&(1<<i) != b[pos]&(1<<i) {
				cnt++
			}
		}
	}
	return cnt
}

package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"math"
	"sort"

	"../utils"
)

var (
	MinKeyLength int = 2
	MaxKeyLength int = 40
)

func main() {
	in := flag.String("f", "1/in/6.txt", "Input file")
	text, err := readInput(in)
	if err != nil {
		panic(err)
	}
	dist := make(KeyScores, 0, MaxKeyLength)
	for i := MinKeyLength; i <= int(math.Min(float64(MaxKeyLength), float64(len(text)/4))); i++ {
		score := avgDistance(text, i)
		dist = append(dist, KeyScore{i, score})
	}
	scored := []utils.KeyEncoding{}
	for _, v := range dist {
		ranges := divide(text, v.KeySize)
		key, err := buildKey(ranges, v.KeySize)
		if err != nil {
			continue
		}
		decoded := utils.XorEncrypt(text, key)
		scored = append(scored, utils.KeyEncoding{Key: string(key), Encoding: string(decoded), Score: utils.Score(decoded)})
	}
	sort.Sort(sort.Reverse(utils.KeyEncodings(scored)))
	fmt.Println(scored[0].Encoding)
}

// buildKey tries to find a XOR cipher for each of the ranges and combines them to form a key
func buildKey(ranges [][]byte, keySize int) ([]byte, error) {
	key := make([]byte, keySize)
	for i, r := range ranges {
		ciphers := utils.FindCipher(r, 20)
		// No cipher for this range so we can't build the key
		if len(ciphers) == 0 {
			return nil, errors.New("No key found for one of the ranges")
		}
		c := ciphers[0]
		key[i] = []byte(c.Key)[0]
	}
	return key, nil
}

func avgDistance(text []byte, i int) float32 {
	return (float32(utils.HammingDistance(text[0:i], text[i:2*i]))/float32(i) +
		float32(utils.HammingDistance(text[i:2*i], text[2*i:3*i]))/float32(i) +
		float32(utils.HammingDistance(text[2*i:3*i], text[i:2*i]))/float32(i))
}

// divide splits the text byte array into keySize of byte ranges
// divide(ABDC,2) -> [[A,C],[B,D]]
func divide(text []byte, keySize int) [][]byte {
	ranges := make([][]byte, keySize)
	for i, v := range text {
		ranges[i%keySize] = append(ranges[i%keySize], v)
	}
	return ranges
}

// readInput reads the input from specified file, converts it to byte array and base64 decodes it
func readInput(inFile *string) ([]byte, error) {
	lines, err := utils.ReadLines(*inFile)
	if err != nil {
		return nil, err
	}
	base64text := []byte{}
	for _, line := range lines {
		base64text = append(base64text, []byte(line)...)
	}
	text := make([]byte, base64.StdEncoding.DecodedLen(len(base64text)))
	base64.StdEncoding.Decode(text, base64text)
	return text, nil
}

type KeyScore struct {
	KeySize int
	Score   float32
}

type KeyScores []KeyScore

func (a KeyScores) Len() int           { return len(a) }
func (a KeyScores) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KeyScores) Less(i, j int) bool { return a[i].Score < a[j].Score }

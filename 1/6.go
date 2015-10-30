package main

import (
	"encoding/base64"
	"fmt"
	"sort"

	"../utils"
)

func main() {
	lines, _ := utils.ReadLines("in/6.txt")
	base64text := []byte{}
	for _, line := range lines {
		base64text = append(base64text, []byte(line)...)
	}
	text := make([]byte, base64.StdEncoding.DecodedLen(len(base64text)))
	base64.StdEncoding.Decode(text, base64text)
	dist := make(KeyScores, 0, 40)
	for i := 2; i <= 40; i++ {
		score := avgDistance(text, i)
		dist = append(dist, KeyScore{i, score})
	}
	fmt.Printf("%v\n", dist)
	sort.Sort(dist)
	for _, v := range dist[0:5] {
		fmt.Printf("%v\n", v.KeySize)
		ranges := divide(text, v.KeySize)
		for _,r:= range ranges {
			utils.FindCipher(r,20)
		}
	}
}

func avgDistance(text []byte, i int) float32 {
	return float32(utils.HammingDistance(text[0:i], text[i:2*i])) / float32(i)
}

func divide(text []byte, keySize int) [][]byte {
	ranges := make([][]byte, keySize)
	for i, v := range text {
		ranges[i%keySize] = append(ranges[i%keySize], v)
	}
	return ranges
}

type KeyScore struct {
	KeySize int
	Score   float32
}

type KeyScores []KeyScore

func (a KeyScores) Len() int           { return len(a) }
func (a KeyScores) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a KeyScores) Less(i, j int) bool { return a[i].Score < a[j].Score }

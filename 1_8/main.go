package main

import (
	"encoding/hex"
	"flag"
	"fmt"

	"github.com/rafax/matasano/utils"
)

var (
	keySize int = 16
)

func main() {
	in := flag.String("f", "1_8/in.txt", "Input file")
	flag.Parse()
	src, _ := readInput(in)
	fmt.Println(len(src))
	for _, l := range src {
		if hasRepeatedRanges(l, keySize) {
			fmt.Println(l)
		}
	}
}

func hasRepeatedRanges(l []byte, keySize int) bool {
	cnt := map[string]int{}
	for i := 0; i < len(l)/keySize; i++ {
		cnt[string(l[i*keySize:(i+1)*keySize])]++
	}
	for k, v := range cnt {
		if v >= 2 {
			fmt.Printf("---> %v repeats %v times\n", []byte(k), v)
			return true
		}
	}
	return false
}

func readInput(inFile *string) ([][]byte, error) {
	lines, err := utils.ReadLines(*inFile)
	if err != nil {
		return nil, err
	}
	fmt.Println(len(lines))
	src := [][]byte{}
	for _, line := range lines {
		c, _ := hex.DecodeString(line)
		src = append(src, c)
	}
	return src, nil
}

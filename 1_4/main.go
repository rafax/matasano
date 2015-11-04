package main

import (
	"flag"
	"fmt"
	"sort"

	"github.com/rafax/matasano/utils"
)

func main() {
	var ip = flag.Int("m", 22, "Minimum score")
	flag.Parse()
	lines, err := utils.ReadLines("1_4/in.txt")
	if err != nil {
		panic(err)
	}
	all := []utils.KeyEncoding{}
	for _, line := range lines {
		dec, _ := utils.DecodeHex(line)
		candidates := utils.FindCipher([]byte(dec), *ip)
		if len(candidates) > 0 {
			all = append(all, candidates...)
		}
	}
	sort.Sort(sort.Reverse(utils.KeyEncodings(all)))
	fmt.Println(all[0])
}

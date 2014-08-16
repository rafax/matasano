package main

import (
	"../utils"
	"flag"
	"fmt"
)

func main() {
	var ip = flag.Int("m", 20, "Minimum score")
	flag.Parse()

	lines, err := utils.ReadLines("in/4.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		candidates := utils.FindCipher(line, *ip)
		if len(candidates) > 0 {
			fmt.Println(candidates)
		}
	}
}

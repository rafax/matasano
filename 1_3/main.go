package main

import (
	"fmt"

	"../utils"
)

func main() {
	lines, err := utils.ReadLines("1_3/in.txt")
	if err != nil {
		panic(err)
	}
	dec, _ := utils.DecodeHex(lines[0])
	candidates := utils.FindCipher(dec, 30)
	fmt.Println(candidates)
}

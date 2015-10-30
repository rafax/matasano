package main

import (
	"../utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadLines("1/in/3.txt")
	if err != nil {
		panic(err)
	}
	dec, _ := utils.DecodeHex(lines[0])
	candidates := utils.FindCipher(dec, 30)
	fmt.Println(candidates)
}

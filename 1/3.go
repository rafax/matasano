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
	candidates := utils.FindCipher(lines[0], 30)
	fmt.Println(candidates)
}

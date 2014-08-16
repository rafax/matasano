package main

import (
	"../utils"
	"fmt"
)

func main() {
	lines, err := utils.ReadLines("in/4.txt")
	if err != nil {
		panic(err)
	}
	for _, line := range lines {
		candidates := utils.FindCipher(line, 0)
		if len(candidates) > 0 {
			fmt.Println(line)
			fmt.Println(candidates)
		}
	}
}

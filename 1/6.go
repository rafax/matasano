package main

import (
	"fmt"

	"../utils"
)

func main() {
	lines, _ := utils.ReadLines("in/6.txt")
	text := []byte{}
	for _, line := range lines {
		text = append(text, []byte(line)...)
	}
	for i := 2; i <= 40; i++ {
		score := avgDistance(text, i)
		fmt.Printf("%d : %f\n", i, score)
	}
}

func avgDistance(text []byte, i int) float32 {
	single := float32(utils.HammingDistance(text[0:i], text[i:2*i])) / float32(i)
	score := single + float32(utils.HammingDistance(text[i:2*i], text[2*i:3*i]))/float32(i)
	score += float32(utils.HammingDistance(text[2*i:3*i], text[3*i:4*i])) / float32(i)
	score /= 3
	return score
}

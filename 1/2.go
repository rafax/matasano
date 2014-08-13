package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	a, _ := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	b, _ := hex.DecodeString("686974207468652062756c6c277320657965")
	e := make([]byte, len(a))
	for i, _ := range a {
		e[i] = a[i] ^ b[i]
	}
	fmt.Println(hex.EncodeToString(e))

}

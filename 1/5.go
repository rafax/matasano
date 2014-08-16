package main

import (
	"../utils"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(hex.EncodeToString(utils.Encrypt([]byte("Burning 'em, if you ain't quick and nimble\nI go crazy when I hear a cymbal"), []byte("ICE"))))
}

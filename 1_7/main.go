package main

import (
	"crypto/aes"
	"encoding/base64"
	"flag"
	"fmt"

	"github.com/rafax/matasano/utils"
)

func main() {
	in := flag.String("f", "in/1_7.txt", "Input file")
	flag.Parse()
	text, _ := readInput(in)
	key := []byte("YELLOW SUBMARINE")
	ciph, _ := aes.NewCipher(key)
	bs := ciph.BlockSize()
	decrypted := make([]byte, len(text))
	tmp := decrypted
	for len(text) > 0 {
		ciph.Decrypt(tmp, text[:bs])
		text = text[bs:]
		tmp = tmp[bs:]
	}
	fmt.Println(string(decrypted))

}

// readInput reads the input from specified file, converts it to byte array and base64 decodes it
func readInput(inFile *string) ([]byte, error) {
	lines, err := utils.ReadLines(*inFile)
	if err != nil {
		return nil, err
	}
	base64text := []byte{}
	for _, line := range lines {
		base64text = append(base64text, []byte(line)...)
	}
	text := make([]byte, base64.StdEncoding.DecodedLen(len(base64text)))
	base64.StdEncoding.Decode(text, base64text)
	return text, nil
}

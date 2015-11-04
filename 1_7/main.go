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
	fmt.Println(*in)
	//text, _ := readInput(in)
	src := []byte("ATOS PORTOS KAKA DEMONA DEMONA!!")
	key := []byte("YELLOW SUBMARINE")
	ciph, _ := aes.NewCipher(key)
	bs := ciph.BlockSize()
	dst := make([]byte, len(src))
	decrypted := make([]byte, len(src))
	tmp := dst
	for len(src) > 0 {
		ciph.Encrypt(tmp, src[:bs])
		src = src[bs:]
		tmp = tmp[bs:]
	}
	fmt.Println(dst)
	tmp = decrypted
	for len(dst) > 0 {
		ciph.Decrypt(tmp, dst[:bs])
		dst = dst[bs:]
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

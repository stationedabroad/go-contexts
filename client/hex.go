package main

import (
	"encoding/hex"
	"fmt"
	"log"
)

func main() {
	src := []byte("48656c6c6f20476f7068657221")

	dst := make([]byte, hex.DecodedLen(len(src)))
	n, err := hex.Decode(dst, src)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)
	fmt.Printf("%s\nstring is: %v\n\x54\x72\x65\x65\x5f", dst, string(src))

}
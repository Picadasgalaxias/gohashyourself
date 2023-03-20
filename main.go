package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func main() {
	h := sha256.Sum256(nil)
	var input [32]byte

	for i := 1; ; i++ {
		input = h
		h = sha256.Sum256(input[:])

		if bytes.Equal(h[:(3)], []byte{0x69, 0x69, 0x69}) {
			fmt.Printf("INPUT:\t\t\t%x\n", input)
			fmt.Printf("GO HASH YOURSELF:\t%x iteration: %d\n", h, i+1)
		}
	}
}

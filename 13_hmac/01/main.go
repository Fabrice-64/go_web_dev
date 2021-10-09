package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

func main() {
	c := getCode("Hello World")
	fmt.Print("First Encryption: ", c, "\n")
	c1 := getCode("hello World")
	fmt.Print("2nd Encryption: ", c1, "\n")
	fmt.Print("\n")
}

func getCode(s string) string {
	h := hmac.New(sha256.New, []byte("hello")) // Includes a private key
	//io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

package main

import (
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	s := "Ceci est encore une phrase où se trouvent, oui, des caractères accentués\nOK"
	s64 := base64.StdEncoding.EncodeToString([]byte(s))

	fmt.Println(s64)
	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		log.Fatalln("Sorry, big mistake", err)
	}
	fmt.Println(string(bs))
}

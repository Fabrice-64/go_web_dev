package main

import (
	"encoding/json"
	"log"
	"os"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{
		// First letter of the fields should be in caps, otherwise throws an error
		State: true,
		Pictures: []string{
			"one.jpg",
			"two.jpg",
			"three.jpg",
		},
	}
	bs, err := json.Marshal(m)
	if err != nil {
		log.Fatalln("Error: ", err)
	}
	os.Stdout.Write(bs)
}

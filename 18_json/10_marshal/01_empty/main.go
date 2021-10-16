package main

import (
	"encoding/json"
	"fmt"
)

type model struct {
	State    bool
	Pictures []string
}

func main() {
	m := model{}
	fmt.Println(m)

	bs, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(bs))
}

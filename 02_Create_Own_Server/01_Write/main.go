package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		io.WriteString(conn, "\nHello from TCP Server!\n")
		fmt.Fprintf(conn, "How is your day?")
		fmt.Fprintln(conn, "%v", "Well I hope")

		conn.Close()

	}

}

// 1er terminal : go run main.go
// Ouvrir 2e terminal: telnet localhost 8080

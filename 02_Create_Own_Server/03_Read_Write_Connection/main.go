package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

// Interagir avec le serveur via le terminal

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))
	if err != nil {
		log.Fatalln("CONN TIMEOUT")
	}
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintf(conn, "You said that: %s\n", ln)
	}
	defer conn.Close()
	fmt.Println("Code got here")
}

// Ouvrir dans le terminal de l'ordi (pas celui de VS Code) avec telnet localhost 8080
// Saisir du texte

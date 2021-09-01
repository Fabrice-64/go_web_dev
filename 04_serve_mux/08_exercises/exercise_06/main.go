package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

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
		serve(conn)
	}
}

func serve(c net.Conn) {
	defer c.Close()
	scanner := bufio.NewScanner(c)
	var rURI, rMethod string
	var i int
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Print(ln)
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD: ", rMethod)
			fmt.Println("URI: :", rURI)
		}
		if ln == "" {
			fmt.Println("END OF REQUEST HEADERS")
			break
		}
		i++
	}
	body := "CHECK OUT THE RESPONSE BODY PAYLOAD"
	body += "\n"
	body += rMethod
	body += "\n"
	body += rURI
	body += "\n"
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Printf("Content-Length: %d\r\n", len(body))
	fmt.Print("Content-Type:  text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}

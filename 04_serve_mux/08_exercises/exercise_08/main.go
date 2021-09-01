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
		go serve(conn)
	}

}

func serve(c net.Conn) {
	defer c.Close()
	var i int
	var rMethod, rURI string
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD: ", rMethod)
			fmt.Println("URI: ", rURI)
		}
		if ln == "" {
			fmt.Println("END OF HEADER REQUEST")
			break
		}
		i++
	}
	switch {
	case rMethod == "GET" && rURI == "/":
		handleIndex(c)
	case rMethod == "GET" && rURI == "/apply":
		handleApply(c)
	case rMethod == "POST" && rURI == "/apply":
		handleApplyPost(c)
	}
}

func handleIndex(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="eng">
		<head>
			<meta charset="UTF-8">
			<title>Exercise 08</title>
		</head>
		<body>
			<h1>GET INDEX</h1>
			<a href="/">Index</a><br>
			<a href="/apply">Apply</a><br>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}

func handleApply(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="eng">
		<head>
			<meta charset="UTF-8">
			<title>Exercise 08</title>
		</head>
		<body>
			<h1>APPLY</h1>
			<a href="/">Index</a><br>
			<a href="/apply">Apply</a><br>
			<form method="POST" action="/apply">
				<input type="hidden" value="This is my value">
				<input type="submit" value="submit">
			</form>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)

}

func handleApplyPost(c net.Conn) {
	body := `
		<!DOCTYPE html>
		<html lang="eng">
		<head>
			<meta charset="UTF-8">
			<title>Exercise 08</title>
		</head>
		<body>
			<h1>POST APPLY</h1>
			<a href="/">Index</a><br>
			<a href="/apply">Apply</a><br>
		</body>
		</html>
	`
	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

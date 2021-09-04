package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/castle", castle)
	http.ListenAndServe(":8080", nil)
}

func castle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// From external server
	io.WriteString(w, `
	<!--File from External server-->
	<img src="https://cdn01.eviivo.media/media/images/2/2/22d29137-ea27-4848-a339-9e798b2e0a50/22d29137-ea27-4848-a339-9e798b2e0a50-l.jpg">
	`)
}

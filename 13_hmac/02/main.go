package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/authenticate", auth)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-session")
	if err != nil {
		c = &http.Cookie{
			Name:  "my-session",
			Value: "",
		}
		if req.Method == http.MethodPost {
			e := req.FormValue("email")
			c.Value = e + "|" + getCode(e)
		}
	}
	http.SetCookie(w, c)
	io.WriteString(w, `<!DOCTYPE html>
		<html lang="eng">
			<head>
				<meta charset="UTF-8">
				<title>INDEX</title>
			</head>
			<body>
				<h1>You are on the Index Page</h1>
				<form method="POST">
					<input type="email" name="email" placeholder="email">
					<input type="submit" value="Valider">
				</form>
				<a href="/authenticate">Validate this value`+c.Value+`</a>
			</body>
		</html>
	`)
}

func auth(w http.ResponseWriter, req *http.Request) {
	c, err := req.Cookie("my-session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	xs := strings.Split(c.Value, "|")
	email := xs[0]
	codeRcvd := xs[1]
	// ajouter un caractère en dessous +"s" pour voir un résultat différent
	codeCheck := getCode(email + "s")

	if codeRcvd != codeCheck {
		fmt.Print("HMAC Codes did not match\n")
		fmt.Println("Check: ", codeCheck)
		fmt.Println("Received:", codeRcvd)
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	io.WriteString(w, `<!DOCTYPE html>
		<html>
			<body>
				<h4>RECEIVED: `+codeRcvd+`</h4>
				<h4>RECALCULATED: `+codeCheck+`<h4>
			</body>
		</html>`)
}

func getCode(data string) string {
	h := hmac.New(sha256.New, []byte("privatekey"))
	io.WriteString(h, data)
	return fmt.Sprintf("%x", h.Sum(nil))
}

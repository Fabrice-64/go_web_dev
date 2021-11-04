package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("postgres", "postgres://<TO SET LOCALLY>@localhost/bookstore?sslmode=disable")
	if err != nil {
		panic(err)
	}
	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("You are Connected !")
}

type Book struct {
	isbn   string
	title  string
	author string
	price  float64
}

func main() {
	http.HandleFunc("/books", BookIndex)
	http.ListenAndServe(":8080", nil)
}

func BookIndex(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	rows, err := db.Query("SELECT * FROM book")
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
	}
	defer rows.Close()
	bks := make([]Book, 0)
	for rows.Next() {
		bk := Book{}
		err := rows.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}
		bks = append(bks, bk)
	}
	if rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	for _, bk := range bks {
		fmt.Fprintf(w, "%s, %s, %s, €%.2f\n", bk.isbn, bk.title, bk.author, bk.price)
	}
}

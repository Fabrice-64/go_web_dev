package main

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func main() {
	db, err = sql.Open("mysql", "root:password@/my_db?charset=utf8")
	check(err)
	defer db.Close()

	err = db.Ping()
	check(err)

	http.HandleFunc("/", index)
	http.HandleFunc("/amigos", amigos)
	http.HandleFunc("/create", create)
	http.HandleFunc("/insert", insert)
	http.HandleFunc("/read", read)
	http.HandleFunc("/update", update)
	http.HandleFunc("/delete", delete)
	http.HandleFunc("/drop", drop)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)

}

func index(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "You are at Index")
	check(err)
}

func amigos(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT aName FROM amigos`)
	check(err)
	defer rows.Close()

	// Data to be used
	var s, name string
	s = "RETRIEVED RECORDS:\n"
	//query
	for rows.Next() {
		err = rows.Scan(&name)
		check(err)
		s += name + "\n"
	}
	fmt.Fprintln(w, s)
}

func create(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`CREATE TABLE customer(name VARCHAR(20))`)
	check(err)
	defer stmt.Close()
	// Create the table
	r, err := stmt.Exec()
	check(err)
	// Check number of rows affected
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "CREATED TABLE customer", n)
}

func insert(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`INSERT INTO customer VALUES ("Ronan")`)
	check(err)
	defer stmt.Close()
	// Execute the Command
	r, err := stmt.Exec()
	check(err)
	// Check number of rows affected
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "INSERTED ITEM INTO customer;", n)
}

func read(w http.ResponseWriter, req *http.Request) {
	rows, err := db.Query(`SELECT * FROM customer;`)
	check(err)
	defer rows.Close()
	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		check(err)
		fmt.Fprintln(w, "RETRIEVED RECORD:", name)
	}
}

func update(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`UPDATE customer SET name="Walou" WHERE name="Ronan";`)
	check(err)
	defer stmt.Close()
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "UPDATE DONE ", n)
}

func delete(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DELETE FROM customer WHERE name="Walou";`)
	check(err)
	defer stmt.Close()
	r, err := stmt.Exec()
	check(err)
	n, err := r.RowsAffected()
	check(err)
	fmt.Fprintln(w, "DELETED A RECORD", n)
}

func drop(w http.ResponseWriter, req *http.Request) {
	stmt, err := db.Prepare(`DROP TABLE customer;`)
	check(err)
	defer stmt.Close()
	_, err = stmt.Exec()
	check(err)
	fmt.Fprintln(w, "TABLE customer DROPPED")

}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

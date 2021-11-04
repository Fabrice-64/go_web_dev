package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = // to set locally
	password = // to set locally
	dbname   = // to set locally
)

type Book struct {
	Isbn   string
	Title  string
	Author string
	Price  string
}

func main() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()

	err = db.Ping()
	CheckError(err)
	fmt.Println("Connecté, Hourra")

	// Insertion in DB
	//insertStmt := `INSERT INTO "book" ("isbn", "title", "author", "price") VALUES ('9783150099001', 'Le Château', 'Franz Kafka', 12.5)`
	//_, err = db.Exec(insertStmt)
	//CheckError(err)

	//insertStmt := `INSERT INTO "book" ("isbn", "title", "author", "price") VALUES ($1, $2, $3, $4)`
	//_, err = db.Exec(insertStmt, "97831500555", "Le Beau Château", "Franz Kafkaiene", 12.2)
	//CheckError(err)

	// Update a record
	//insertStmt := `UPDATE "book" SET "isbn" = $1, "title"= $2, "author"= $3, "price"= $4 WHERE "isbn" = $1`
	//_, err = db.Exec(insertStmt, "97831500555", "Le Grand Château", "Franz Kafkaiene", 200)
	//CheckError(err)

	// Delete a record
	deleteStmt := `DELETE FROM "book" WHERE "isbn" = $1`
	_, err = db.Exec(deleteStmt, "97831500555")
	CheckError(err)

	// Select Several records
	rows, err := db.Query(`SELECT * FROM "book"`)
	CheckError(err)
	defer rows.Close()
	for rows.Next() {
		bk := Book{}
		err = rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		CheckError(err)
		fmt.Println(bk.Isbn, bk.Author, bk.Price)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	pass     = "pass1234"
	host     = "localhost"
	port     = 5432
	database = "dbName"
)

func main() {
	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, database)
	db, err := sql.Open("posgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("hello", user, "you are now connected to ", host, "in port", port, "to database", database)
}

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	pass     = "pass1234"
	host     = "localhost"
	port     = 5432
	database = "address"
)

var migration string = `

CREATE TABLE address_book (id serial primary key not null, country char(50), city char(50), street text, number int, zip int);

INSERT INTO address_book(country, city, street, number, zip) VALUES ('Israel','Tel Aviv','Shibutei Elkatraz u Banav',143, 123987);
`

type Address struct {
	country string
	city    string
	street  string
	number  int
	zip     int
}

var db *sql.DB

// initialize db connection - if cannot ping it will panic
func init() {
	var err error
	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, database)
	db, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatal("could not connect to db: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("hello", user, "you are now connected to ", host, "in port", port, "to database", database)

	_, err = db.Exec(migration)
	if err != nil {
		log.Println("db migration failed: ", err)
	}
	fmt.Println("db migrations complete successfully")
}

// running api
func main() {
	//http.HandleFunc("/add", addAddress)
	http.HandleFunc("/index", listAddresses)
	http.ListenAndServe(":8080", nil)
}

func listAddresses(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
		log.Println("not found dude... - 405")
	}

	rows, err := db.Query("SELECT country, city, street, number, zip FROM address_book")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println("internal server error - 500")
		return
	}
	defer rows.Close()

	addrs := make([]*Address, 0)
	for rows.Next() {
		addr := new(Address)
		err := rows.Scan(&addr.country, &addr.city, &addr.street, &addr.number, &addr.zip)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		addrs = append(addrs, addr)
	}
	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		log.Println("internal server error - 500")
		return
	}

	for _, addr := range addrs {
		fmt.Fprintf(w, "%s, %s, %s, %v, %v", addr.country, addr.city, addr.street, addr.number, addr.zip)
	}
}

/*
func addAddress() {

}
*/

// package main will run simple web app to show address book
// and add addresses.
// to run this app use Makefile
// to add an address use this command:
// curl -i -X POST -d "Sweden&city=Stokholm&street=HadNess&number=12&zip=54321" localhost:8080/add
package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// connection based on that - compare with docker-conpose or docker commad in Makefile
const (
	user     = "postgres"
	pass     = "pass1234"
	host     = "localhost"
	port     = 5432
	database = "address"
)

// intial db migration - create the table and add one address
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
	http.HandleFunc("/add", addAddress)
	http.HandleFunc("/index", listAddresses)
	http.ListenAndServe(":8080", nil)
}

// list all addresses in the database
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
		fmt.Fprintf(w, "\n%s, %s, %s, %v, %v", addr.country, addr.city, addr.street, addr.number, addr.zip)
	}
}

// api to add additional address
// example: curl -i -X POST -d "Sweden&city=Stokholm&street=HadNess&number=12&zip=54321" localhost:8080/add
// to add an address.
func addAddress(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
		return
	}

	country := r.FormValue("country")
	city := r.FormValue("city")
	street := r.FormValue("street")
	number := r.FormValue("number")
	zip := r.FormValue("zip")

	if country == "" || city == "" || street == "" {
		http.Error(w, http.StatusText(400), 400)
		return
	}

	address, err := db.Exec("INSERT INTO address_book(country, city, street, number, zip) VALUES ($1, $2, $3, $4, $5)", country, city, street, number, zip)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	affRows, err := address.RowsAffected()
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	addstr := fmt.Sprintf("%s %s %s %d %d", country, city, street, number, zip)
	fmt.Fprintf(w, "Address ' %s ' created successfully and %d row added", addstr, affRows)
}

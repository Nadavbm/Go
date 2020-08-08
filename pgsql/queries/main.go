package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	user     = "postgres"
	pass     = "postgres"
	host     = "localhost"
	port     = 5432
	database = "dbName"
)

var migration string = `
CREATE TABLE albums (id serial primary key not null, name char(50) not null,band text not null, year int, genre text, price int);

INSERT INTO albums (name, band, year, genre) VALUES ('crack the sky', 'mastodon', 2008, 'metal');
`

type Album struct {
	name  string
	band  string
	year  int
	genre string
}

func main() {
	conn := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", user, pass, host, database)
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("hello", user, "you are now connected to ", host, "in port", port, "to database", database)

	_, err = db.Exec(migration)
	if err != nil {
		log.Println("db migraiton failed", err)
	}

	rows, err := db.Query("SELECT name, band, year, genre FROM albums;")
	if err != nil {
		log.Println("cannot query database", err)
	}
	defer rows.Close()

	albums := make([]*Album, 0)
	for rows.Next() {
		al := new(Album)
		err := rows.Scan(&al.name, &al.band, &al.year, &al.genre)
		if err != nil {
			log.Fatal("could not assing values from database in struct album: ", err)
		}
		albums = append(albums, al)
	}

	for _, al := range albums {
		fmt.Println("album name: ", al.name, "\tplay by the band", al.band, "year", al.year)
	}

}

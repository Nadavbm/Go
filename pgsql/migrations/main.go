// this is an example for running db migrations
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
CREATE TABLE people (pid serial primary key not null, name char(50) not null);

CREATE TABLE cars (cid serial primary key not null, car_type char(50) not null, car_license text not null);

CREATE TABLE rent_time (tid serial primary key not null, pid int references people(pid), cid int references cars(cid), rental_hours char(50));

INSERT INTO people(name) VALUES ('shimon'), ('nehemia'), ('bezazek');

INSERT INTO cars(car_type, car_license) VALUES ('shevrolet','123-123-33'), ('renault', '123-124-44');

INSERT INTO rent_time (pid, cid, rental_hours) VALUES (1,1, '8.8.2020 17:00 - 12.8.2020 18:00'), (2,1, '20.7.2020 11:00 - 22.7.2020 17:00');
`

var query = `
SELECT * FROM people inner join rent_time on people.pid = rent_time.pid;

SELECT * FROM people a inner join rent_time b on a.pid = b.pid;

SELECT people.name, cars.car_license FROM people LEFT OUTER JOIN cars ON people.pid = cars.cid;

SELECT people.name, cars.car_license FROM people RIGHT OUTER JOIN cars ON people.pid = cars.cid;

SELECT people.name, cars.car_license FROM people FULL OUTER JOIN cars ON people.pid = cars.cid;
`

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

	_, err = db.Exec(migration)
	if err != nil {
		log.Fatalln("db migration failed: ", err)
	}

	_, err = db.Exec(query)
	if err != nil {
		log.Println("could not run query", err)
	}

	fmt.Println("connect to the db and list tables")
}

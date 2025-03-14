package schema

import (
	"database/sql"
	"log"
)

var usert_schema = `
CREATE TABLE users (
	uid INT NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    dob DATE,
    password VARCHAR(255),
	language VARCHAR(255)
)`

type User struct {
	Uid      int    `db:"uid"`
	Name     string `db:"name"`
	Dob      string `db:"dob"`
	Password string `db:"password"`
	Language string `db:"language"`
}

func CreateUsersTables() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	drop := "DROP TABLE IF EXISTS users;"
	_, err = db.Exec(drop + usert_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("User Table Created")
}

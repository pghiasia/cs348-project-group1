package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var schema = `
CREATE TABLE movies (
    name text,
    genre text,
    director text
)`

type Movie struct {
	Name     string `db:"name"`
	Genre    string `db:"genre"`
	Director string `db:"director"`
}

func main() {
	db, err := sqlx.Connect("postgres", "user=postgres dbname=movies sslmode=disable password=cs348 host=localhost")
	if err != nil {
		log.Fatalln(err)
	}

	db.MustExec(schema)
	tx := db.MustBegin()
	tx.MustExec("INSERT INTO movies (name, genre, director) VALUES ($1, $2, $3)", "Madagascar", "Comedy", "Tom McGrath")
	tx.MustExec("INSERT INTO movies (name, genre, director) VALUES ($1, $2, $3)", "John Wick", "Action", "Chad Stahelski")
	tx.MustExec("INSERT INTO movies (name, genre, director) VALUES ($1, $2, $3)", "Forrest Gump", "Romance", "Robert Zemeckis")
	tx.Commit()

	movies := []Movie{}
	db.Select(&movies, "SELECT * FROM movies ORDER BY name ASC")
	m1, m2, m3 := movies[0], movies[1], movies[2]

	fmt.Printf("%#v\n%#v\n%#v", m1, m2, m3)
}

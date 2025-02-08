package schema

import (
	"database/sql"
	"log"
)

var actor_schema = `
CREATE TABLE actors (
	aid VARCHAR(255) NOT NULL PRIMARY KEY,
    birthyear VARCHAR(255),
    deathyear VARCHAR(255),
	name VARCHAR(255)
)`

type Actor struct {
	Aid       string `db:"aid"`
	Birthyear string `db:"birthyear"`
	Deathyear string `db:"deathyear"`
	Name      string `db:"name"`
}

func CreateActorTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    dropActor := "DROP TABLE IF EXISTS actors;"
    dropCast := "DROP TABLE IF EXISTS movie_to_actor;"
	_, err = db.Exec(dropCast + dropActor + actor_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("Actor Table Created")
}

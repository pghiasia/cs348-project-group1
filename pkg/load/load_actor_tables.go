package load

import (
	"database/sql"
	"log"
)

func LoadActorsTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO actors
	SELECT nconst as aid, birthYear, deathYear, primaryName as name 
	FROM read_csv("./test-data/name.basics.tsv")
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Actors Loaded")
}

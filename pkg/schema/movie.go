package schema

import (
	"database/sql"
	"log"
)


var movies_schema = `
CREATE TABLE movies (
	mid VARCHAR(255) NOT NULL PRIMARY KEY,
    title VARCHAR(255),
    genres VARCHAR(255),
    release BIGINT,
	rating FLOAT
)`

var movie_to_actors_schema = `
CREATE TABLE movie_to_actor (
	mid VARCHAR(255),
	aid VARCHAR(255),
	PRIMARY KEY(mid, aid),
	FOREIGN KEY(mid) REFERENCES movies, 
	FOREIGN KEY(aid) REFERENCES actors
)`

type Movie struct {
	Mid     string  `db:"mid"`
	Title   string  `db:"title"`
	Genres   string  `db:"genres"`
	Release string  `db:"release"`
	Rating  float32 `db:"float"`
}

type MovieToActor struct {
	Mid string `db:"mid"`
	Aid string `db:"aid"`
}

func CreateMoviesTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    dropMovies := "DROP TABLE IF EXISTS movies;"
    dropCast := "DROP TABLE IF EXISTS movie_to_actor;"
    _, err = db.Exec(dropCast + dropMovies + movies_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("Movie Table Created")

	_, err = db.Exec(movie_to_actors_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("Movie To Actors Relation Created")
}

package schema

import (
	"database/sql"
	"log"
)

var movies_schema = `
CREATE TABLE movies (
	mid VARCHAR(255) NOT NULL PRIMARY KEY,
    title VARCHAR(255),
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

var movie_to_genre = `
CREATE TABLE movie_to_genre (
	mid VARCHAR(255),
	genrename VARCHAR(255),
	PRIMARY KEY(mid, genrename),
	FOREIGN KEY(mid) REFERENCES movies, 
)`

type Movie struct {
	Mid     string  `db:"mid"`
	Title   string  `db:"title"`
	release string  `db:"release"`
	rating  float32 `db:"float"`
}

type MovieToActor struct {
	Mid string `db:"mid"`
	Aid string `db:"aid"`
}

type MovieToGenre struct {
	Mid       string `db:"mid"`
	GenreName string `db:"genrename"`
}

func CreateMoviesTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(movies_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("Movie Table Created")

	_, err = db.Exec(movie_to_actors_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("Movie To Actors Relation Created")

	_, err = db.Exec(movie_to_genre)
	if err != nil {
		log.Fatal(err)
	}
	println("Movie To Genre Relation Created")
}

package load

import (
	"database/sql"
	"log"
)

func LoadMoviesTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO movies
	SELECT name_basics.tconst as mid, primaryTitle as title, genres, startYear as release, averageRating as rating, numVotes
	FROM read_csv("./test-data/title.basics.tsv") as name_basics
	JOIN read_csv("./test-data/title.ratings.tsv") as ratings
		ON name_basics.tconst = ratings.tconst
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Movies Loaded")
}

func LoadMovieToActorTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO movie_to_actor
	SELECT titles.tconst as mid, names.nconst as aid
	FROM read_csv("./test-data/title.basics.tsv") as titles
	JOIN read_csv("./test-data/title.principals.tsv") as principals
		ON titles.tconst = principals.tconst
		JOIN read_csv("./test-data/name.basics.tsv") as names
			ON principals.nconst = names.nconst
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Movies To Actor Relation Data Loaded")
}

func LoadMovieToGenreTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO movie_to_genre
	SELECT tconst as mid, genres as genrename
	FROM read_csv('./test-data/title.basics.tsv')
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Movies To Genre Relation Data Loaded")
}

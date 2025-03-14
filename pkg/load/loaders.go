package load

import (
	"database/sql"
	"log"
)

func LoadTitlesTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO titles
	SELECT tconst as tID
	FROM read_csv("./test-data/title.basics.tsv")
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Actors Loaded")
}


func LoadPeopleTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO people
	SELECT nconst as pID, primaryProfession, birthYear, deathYear, primaryName as name, knownForTitles
	FROM read_csv("./test-data/name.basics.tsv")
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Actors Loaded")
}

func LoadEpisodesTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO episodes
	SELECT tconst as tID, parentTconst AS seriesID, isAdult, startYear AS releaseYear, originalTitle, averageRating, numVotes, runtimeMinutes, primaryTitle, episodeNumber, seasonNumber
	FROM (read_csv("./test-data/title.basics.tsv") NATURAL JOIN read_csv("./test-data/title.ratings.tsv")) NATURAL JOIN read_csv("./test-data/title.episode.tsv")
	WHERE titleType = 'tvepisode'
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Episode titles Loaded")
}


func LoadSeriesTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO series
	SELECT tconst as tID, isAdult, startYear AS releaseYear, endYear, originalTitle, primaryTitle, runtimeMinutes, averageRating, numVotes
	FROM read_csv("./test-data/title.basics.tsv") NATURAL JOIN read_csv("./test-data/title.ratings.tsv")
	WHERE titleType = 'tvseries'
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Series titles Loaded")
}


func LoadShortTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO short
	SELECT tconst as tID, isAdult, startYear AS releaseYear, primaryTitle, originalTitle, runtimeMinutes, averageRating, numVotes
	FROM read_csv("./test-data/title.basics.tsv") NATURAL JOIN read_csv("./test-data/title.ratings.tsv")
	WHERE titleType = 'short'
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Short titles Loaded")
}


func LoadMovieTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO movie
	SELECT tconst as tID, isAdult, startYear AS releaseYear, primaryTitle, originalTitle, runtimeMinutes, averageRating, numVotes
	FROM read_csv("./test-data/title.basics.tsv") NATURAL JOIN read_csv("./test-data/title.ratings.tsv")
	WHERE titleType = 'movie'
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("Movie titles Loaded")
}

func LoadWorkedOnTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO workedOn
	SELECT tconst as tID, nconst as pID, job as jobTitle, ordering AS creditOrder, category, characters
	FROM read_csv("./test-data/title.principals.tsv")
	`
	_, err = db.Exec(insertion_query)
	if err != nil {
		log.Fatal(err)
	}

	println("workedOn Loaded")
}


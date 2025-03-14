package load

import (
	"database/sql"
	"log"
)

func LoadFavTitlesTable(basicPath string, userPath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO favTitles 
        SELECT * FROM read_csv('?') NATURAL JOIN`, basicPath, userPath)
	if err != nil {
		log.Fatal(err)
	}

	println("Users Loaded")
}


func LoadUsersTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users SELECT * FROM read_csv('?')`, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Users Loaded")
}

func LoadTitlesTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO titles
	SELECT tconst as tID
	FROM read_csv(?, delim='\t')
	`
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Actors Loaded")
}


func LoadPeopleTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO people
	SELECT nconst as pID, primaryName as name, birthYear, deathYear, primaryProfession, knownForTitles
	FROM read_csv(?, delim='\t')
	`
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Actors Loaded")
}

func LoadEpisodesTable(basicPath string, ratingsPath string, episodePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO episodes
	SELECT tconst as tID, parentTconst AS seriesID, isAdult, startYear AS releaseYear, originalTitle, averageRating, numVotes, runtimeMinutes, primaryTitle, episodeNumber, seasonNumber
	FROM read_csv(?, delim='\t') NATURAL JOIN read_csv(?, delim='\t') NATURAL JOIN read_csv(?, delim='\t')
	WHERE titleType = 'tvepisode'
	`
	_, err = db.Exec(insertion_query, basicPath, ratingsPath, episodePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Episode titles Loaded")
}


func LoadSeriesTable(basicPath string, ratingsPath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO series
	SELECT tconst as tID, isAdult, startYear AS releaseYear, endYear, originalTitle, primaryTitle, runtimeMinutes, averageRating, numVotes
	FROM read_csv(?, delim='\t') NATURAL JOIN read_csv(?, delim='\t')
	WHERE titleType = 'tvseries'
	`
	_, err = db.Exec(insertion_query, basicPath, ratingsPath)
	if err != nil {
		log.Fatal(err)
	}

	println("Series titles Loaded")
}


func LoadShortTable(basicPath string, ratingsPath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO short
	SELECT tconst as tID, isAdult, startYear AS releaseYear, primaryTitle, originalTitle, runtimeMinutes, averageRating, numVotes
	FROM read_csv(?, delim='\t') NATURAL JOIN read_csv(?, delim='\t')
	WHERE titleType = 'short'
	`
	_, err = db.Exec(insertion_query, basicPath, ratingsPath)
	if err != nil {
		log.Fatal(err)
	}

	println("Short titles Loaded")
}


func LoadMovieTable(basicPath string, ratingsPath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO movie
	SELECT tconst as tID, isAdult, startYear AS releaseYear, primaryTitle, originalTitle, runtimeMinutes, averageRating, numVotes
	FROM read_csv(?, delim='\t') NATURAL JOIN read_csv(?, delim='\t')
	WHERE titleType = 'movie'
	`
	_, err = db.Exec(insertion_query, basicPath, ratingsPath)
	if err != nil {
		log.Fatal(err)
	}

	println("Movie titles Loaded")
}

func LoadWorkedOnTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO workedOn
	SELECT tconst as tID, nconst as pID, job as jobTitle, ordering AS creditOrder, category, characters
	FROM read_csv(?, delim='\t')
	`
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("workedOn Loaded")
}

func LoadGenresTable(filePath string) {
    db, err := sql.Open("duckdb", "./movie.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    var insertion_query = `
        INSERT INTO genres
        SELECT unnest(string_split(T.genres, ',')) as genre, tconst as tID
        FROM read_csv(?, delim='\t') as T;
    `
    _, err = db.Exec(insertion_query, filePath)
    if err != nil {
        log.Fatal(err)
    }

    println("Genres Loaded")
}

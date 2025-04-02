package load

import (
	"database/sql"
	"log"
)

func LoadRanksTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
        INSERT INTO ranks
        SELECT uID, ranking, tconst AS tID FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='');
    `
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Ranks Loaded")
}

func LoadUsersTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`
        INSERT INTO users 
        SELECT uID, name, DOB, language, password
        FROM read_csv(?)`,
		filePath)
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
    FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='')
	`
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Titles Loaded")
}

func LoadPeopleTable(filePath string) {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var insertion_query = `
	INSERT INTO people
	SELECT nconst as pID, birthYear, deathYear, primaryName as name, knownForTitles, primaryProfession
	FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='')
	`
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("People Loaded")
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
	FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A1 
    NATURAL LEFT OUTER JOIN read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A2 
    NATURAL JOIN read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A3
	WHERE titleType = 'tvEpisode';
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
	SELECT tconst as tID, isAdult, startYear AS releaseYear, endYear, originalTitle, averageRating, numVotes, runtimeMinutes, primaryTitle
	FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A1
    NATURAL LEFT OUTER JOIN read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A2
	WHERE titleType = 'tvSeries' OR titleType = 'tvMiniSeries'
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
	SELECT tconst as tID, isAdult, startYear AS releaseYear, originalTitle, averageRating, numVotes, runtimeMinutes, primaryTitle
	FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A1 
    NATURAL LEFT OUTER JOIN read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A2
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
	SELECT tconst as tID, isAdult, startYear AS releaseYear, originalTitle, averageRating, numVotes, runtimeMinutes, primaryTitle
	FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A1 
    NATURAL LEFT OUTER JOIN read_csv(?, delim='\t', nullstr='\N', quote='', escape='') AS A2
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
        SELECT nconst AS pID, tconst as tID, job as jobTitle, ordering as creditOrder, category, characters
        FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='') JOIN titles ON titles.tid = tconst
        JOIN people ON people.pID = nconst;
    `
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("ALTER TABLE workedOn ADD PRIMARY KEY (tID, creditOrder)")
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
        FROM read_csv(?, delim='\t', nullstr='\N', quote='', escape='') as T;
    `
	_, err = db.Exec(insertion_query, filePath)
	if err != nil {
		log.Fatal(err)
	}

	println("Genres Loaded")
}

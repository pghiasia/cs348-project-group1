package schema

import (
	"database/sql"
	"log"
)

var titles_schema = `
CREATE TABLE titles (
    tID VARCHAR NOT NULL PRIMARY KEY
);`

var people_schema = `
CREATE TABLE people (
	pID VARCHAR NOT NULL PRIMARY KEY,
    birthYear INT,
    deathYear INT,
	name VARCHAR,
	knownForTitles VARCHAR,
	primaryProfession VARCHAR
);`

var ranks_schema = `
CREATE TABLE ranks (
    uID INT NOT NULL,
    ranking INT NOT NULL,
    tID VARCHAR NOT NULL,
    FOREIGN KEY (uID) REFERENCES users(uID),
    FOREIGN KEY (tID) REFERENCES titles(tID),
    PRIMARY KEY (uID, ranking)
);`

// moved primary key constraint for efficiency reason.
var workedOn_schema = `
CREATE TABLE workedOn (
	pID VARCHAR NOT NULL,
	tID VARCHAR NOT NULL,
    jobTitle VARCHAR,
    creditOrder INT NOT NULL,
	category VARCHAR NOT NULL,
	characters VARCHAR,
	FOREIGN KEY (pID) REFERENCES people(pID),
	FOREIGN KEY (tID) REFERENCES titles(tID)
);`


var genres_schema = `
CREATE TABLE genres (
	genre VARCHAR,
	tID VARCHAR,
	PRIMARY KEY (genre, tID),
    FOREIGN KEY (tID) REFERENCES titles(tID),
);`

var users_schema = `
CREATE TABLE users(
	uID INT PRIMARY KEY,
	name VARCHAR NOT NULL UNIQUE,
	DOB DATE NOT NULL,
	language VARCHAR NOT NULL,
	password VARCHAR NOT NULL,
);
`

var episodes_schema = `
CREATE TABLE episodes (
	tID VARCHAR PRIMARY KEY,
    seriesID VARCHAR NOT NULL,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR,
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR,
	episodeNumber INT,
	seasonNumber INT,
	FOREIGN KEY(seriesID) REFERENCES series(tID),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var series_schema = `
CREATE TABLE series (
	tID VARCHAR PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	endYear INT,
	originalTitle VARCHAR,
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR,
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`
var short_schema = `
CREATE TABLE short(
	tID VARCHAR PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR,
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR,
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var movie_schema = `
CREATE TABLE movie (
	tID VARCHAR PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR,
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR,
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0)
);
`

func CreateTables() {
    db, err := sql.Open("duckdb", "./movie.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    dropShort := "DROP TABLE IF EXISTS short;"
    dropMovie := "DROP TABLE IF EXISTS movie;"
    dropEpisodes := "DROP TABLE IF EXISTS episodes;"
    dropSeries := "DROP TABLE IF EXISTS series;"
    dropGenres := "DROP TABLE IF EXISTS genres;"
    dropRanks := "DROP TABLE IF EXISTS ranks;"
	dropTitles := "DROP TABLE IF EXISTS titles;"
    dropWorkedOn := "DROP TABLE IF EXISTS workedOn;"
    dropPeople := "DROP TABLE IF EXISTS people;"
    dropUsers := "DROP TABLE IF EXISTS users;"

    // This is technically not good since db.Exec is supposed to execute a single statement not multiple.
    _, err = db.Exec(dropShort + dropMovie + dropEpisodes + dropSeries + dropRanks  + dropGenres + dropWorkedOn + dropTitles +  dropPeople + dropUsers)
    if err != nil {
        log.Fatal(err)
    }
    println("Pre-existing tables dropped")


    _, err = db.Exec(titles_schema + people_schema + users_schema + series_schema + short_schema + 
		movie_schema + episodes_schema + ranks_schema + genres_schema + workedOn_schema)
    if err != nil {
        log.Fatal(err)
    }
    println("Relations Created")
}

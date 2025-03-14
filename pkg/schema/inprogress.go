package schema

import (
	"database/sql"
	"log"
)

var titles_schema = `
CREATE TABLE titles (
    tID VARCHAR(255) NOT NULL PRIMARY KEY
);`

var people_schema = `
CREATE TABLE people (
	pID VARCHAR(255) NOT NULL PRIMARY KEY,
    birthYear VARCHAR(255),
    deathYear VARCHAR(255),
	name VARCHAR(255),
	knownForTitles VARCHAR(255),
	primaryProfession VARCHAR(255)
);`

var workedOn_schema = `
CREATE TABLE workedOn (
	pID VARCHAR(255) NOT NULL,
	tID VARCHAR(255) NOT NULL,
    jobTitle VARCHAR(255),
    creditOrder VARCHAR(255),
	category VARCHAR(255),
	characters VARCHAR(255),
	PRIMARY KEY (pID, tID),
	FOREIGN KEY (pID) REFERENCES people(pID),
	FOREIGN KEY (tID) REFERENCES titles(tID)
);`


var genres_schema = `
CREATE TABLE genres (
	genre VARCHAR(255),
	tID VARCHAR(255),
	PRIMARY KEY (genre, tID),
    FOREIGN KEY (tID) REFERENCES titles(tID),
);`

var users_schema = `
CREATE TABLE users(
	uID INT PRIMARY KEY,
	name VARCHAR(255) NOT NULL UNIQUE,
	DOB DATE NOT NULL,
	language VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
);
`

var fave_titles_schema = `
CREATE TABLE favTitles (
	uID INT,
	tID VARCHAR(255),
	ranking VARCHAR(255),
	PRIMARY KEY (uID, tID),
    FOREIGN KEY (uID) REFERENCES users(uID),
	FOREIGN KEY (tID) REFERENCES titles(tID),
);
`

var episodes_schema = `
CREATE TABLE episodes (
	tID VARCHAR(255) PRIMARY KEY,
    seriesID VARCHAR(255) NOT NULL,
	isAdult VARCHAR(255),
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes VARCHAR(255),
	primaryTitle VARCHAR(255),
	episodeNumber VARCHAR(255),
	seasonNumber VARCHAR(255),
	FOREIGN KEY(seriesID) REFERENCES series(tID),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var series_schema = `
CREATE TABLE series (
	tID VARCHAR(255) PRIMARY KEY,
	isAdult VARCHAR(255),
    releaseYear VARCHAR(255),
	endYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes VARCHAR(255),
	primaryTitle VARCHAR(255),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`
var short_schema = `
CREATE TABLE short(
	tID VARCHAR(255) PRIMARY KEY,
	isAdult VARCHAR(255),
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes VARCHAR(255),
	primaryTitle VARCHAR(255),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var movie_schema = `
CREATE TABLE movie (
	tID VARCHAR(255) PRIMARY KEY,
	isAdult VARCHAR(255),
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DOUBLE,
	numVotes INT,
	runtimeMinutes VARCHAR(255),
	primaryTitle VARCHAR(255),
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
    dropFavTitles := "DROP TABLE IF EXISTS favTitles;"
    dropGenres := "DROP TABLE IF EXISTS genres;"
	dropTitles := "DROP TABLE IF EXISTS titles;"
    dropWorkedOn := "DROP TABLE IF EXISTS workedOn;"
    dropPeople := "DROP TABLE IF EXISTS people;"
    dropUsers := "DROP TABLE IF EXISTS users;"

    _, err = db.Exec(dropShort + dropMovie + dropEpisodes + dropSeries + dropFavTitles + dropGenres + dropWorkedOn + dropTitles +  dropPeople + dropUsers)
    if err != nil {
        log.Fatal(err)
    }
    println("Pre-existing tables dropped")



    _, err = db.Exec(titles_schema + people_schema + users_schema + series_schema + short_schema + 
		movie_schema + episodes_schema + fave_titles_schema + genres_schema + workedOn_schema)
    if err != nil {
        log.Fatal(err)
    }
    println("people Relation Created")
}

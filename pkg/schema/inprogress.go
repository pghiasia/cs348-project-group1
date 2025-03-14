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
    birthYear INT,
    deathYear INT,
	name VARCHAR(255),
	knownForTitles VARCHAR(255),
	primaryProfession VARCHAR(255)
);`

var workedOn_schema = `
CREATE TABLE workedOn (
	pID VARCHAR(255) NOT NULL,
	tID VARCHAR(255) NOT NULL,
    jobTitle VARCHAR(255),
    creditOrder INT,
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
	uID VARCHAR(255),
	tID VARCHAR(255),
	ranking INT,
	PRIMARY KEY (uID, tID),
    FOREIGN KEY (uID) REFERENCES users(uID),
	FOREIGN KEY (tID) REFERENCES titles(tID),
);
`

var episodes_schema = `
CREATE TABLE episodes (
	tID VARCHAR(255) PRIMARY KEY,
    seriesID VARCHAR(255) NOT NULL,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	episodeNumber INT,
	seasonNumber INT,
	FOREIGN KEY(seriesID) REFERENCES series(tID),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var series_schema = `
CREATE TABLE series (
	tID VARCHAR(255) PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	endYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`
var short_schema = `
CREATE TABLE short(
	tID VARCHAR(255) PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var movie_schema = `
CREATE TABLE movie (
	tID VARCHAR(255) PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
    FOREIGN KEY(tid) REFERENCES titles(tid),
    CHECK (averageRating <= 10 AND averageRating >= 0)
);
`

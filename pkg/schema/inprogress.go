package schema

import (
	"database/sql"
	"log"
)

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
	genre VARCHAR(256),
	tID VARCHAR(256),
	PRIMARY KEY (genre, tID),
    FOREIGN KEY (pID) REFERENCES titles(tID),
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
	tID VARCHAR(256),
	uID VARCHAR(256),
	ranking INT,
	PRIMARY KEY (pID, tID),
    FOREIGN KEY (pID) REFERENCES users(uID),
	FOREIGN KEY (tID) REFERENCES titles(tID),
);
`

var episodes_schema = `
CREATE TABLE episodes (
	tID VARCHAR(256) PRIMARY KEY,
    seriesID VARCHAR(256) NOT NULL,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	episodeNumber INT,
	seasonNumber INT,
	FOREIGN KEY (seriesID) REFERENCES series(tID),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var series_schema = `
CREATE TABLE series (
	tID VARCHAR(256) PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	endYear INT,
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`
var short_schema = `
CREATE TABLE short(
	tID VARCHAR(256) PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
    CHECK (averageRating <= 10 AND averageRating >= 0),
);
`

var movie_schema = `
CREATE TABLE movie (
	tID VARCHAR(256) PRIMARY KEY,
	isAdult BOOLEAN,
    releaseYear INT,
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
    CHECK (averageRating <= 10 AND averageRating >= 0)
);
`

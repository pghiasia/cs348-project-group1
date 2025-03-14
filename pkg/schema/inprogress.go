package schema

import (
	"database/sql"
	"log"
)

var people_schema = `
CREATE TABLE people (
	pID VARCHAR(255) NOT NULL PRIMARY KEY,
    birthYear VARCHAR(255),
    deathYear VARCHAR(255),
	name VARCHAR(255),
	knownForTitles VARCHAR(255),
	primaryProfession VARCHAR(255)
)`

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
)`





var related_genres_schema = `
CREATE TABLE relatedGenres (
	genre VARCHAR(256),
	tID VARCHAR(256),
	PRIMARY KEY (genre, tID),
    FOREIGN KEY (pID) REFERENCES titles(tID),
	FOREIGN KEY (genre) REFERENCES genres(genre),
)	

`

var genres = `
CREATE TABLE genres (
	genre VARCHAR(256) PRIMARY KEY,
)	
`

var users_schema = `
CREATE TABLE users(
	uID VARCHAR(255) PRIMARY KEY,
	name VARCHAR(255),
	DOB DATE NOT NULL,
	language VARCHAR(255),
	password VARCHAR(255) NOT NULL,
	
)
`

var fave_titles_schema = `
CREATE TABLE favTitles (
	tID VARCHAR(256),
	uID VARCHAR(256),
	ranking INT,
	PRIMARY KEY (pID, tID),
    FOREIGN KEY (pID) REFERENCES users(uID),
	FOREIGN KEY (tID) REFERENCES titles(tID),
)
`



var episodes_schema = `
CREATE TABLE episodes (
	CHECK (rating <= 10 AND rating >= 0),
	isAdult BOOLEAN,
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	tID VARCHAR(256) PRIMARY KEY,
	episodeNumber INT,
	seasonNumber INT,
	FOREIGN KEY (tID) REFERENCES titles(tID),
)
`

var series_schema = `
CREATE TABLE series (
CHECK (rating <= 10 AND rating >= 0),
	isAdult BOOLEAN,
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	tID VARCHAR(256) PRIMARY KEY,
	endYear INT,
	FOREIGN KEY (tID) REFERENCES titles(tID),
)
`
var short_schema = `
CREATE TABLE short(
CHECK (rating <= 10 AND rating >= 0),
	isAdult BOOLEAN,
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	tID VARCHAR(256) PRIMARY KEY,
	FOREIGN KEY (tID) REFERENCES titles(tID),
)
`

var movie_schema = `
CREATE TABLE movie (
CHECK (rating <= 10 AND rating >= 0),
	isAdult BOOLEAN,
    releaseYear VARCHAR(255),
	originalTitle VARCHAR(255),
	averageRating DECIMAL(2,1),
	numVotes INT,
	runtimeMinutes INT,
	primaryTitle VARCHAR(255),
	tID VARCHAR(256) PRIMARY KEY,
	FOREIGN KEY (tID) REFERENCES titles(tID),
)
`

var has_episode_schema = `
CREATE TABLE hasEpisode (
	episodeID VARCHAR(256) UNIQUE,
    seriesID VARCHAR(256),
	PRIMARY KEY (episodeID, seriesID),
	FOREIGN KEY (episodeID) REFERENCES episode(tID),
	FOREIGN KEY (seriesID) REFERENCES series(tID),
)
`



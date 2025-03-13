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

var titles_schema = `
CREATE TABLE workedOn (
	isAdult BOOLEAN,
	tID VARCHAR(255) NOT NULL,
    releaseYear VARCHAR(255),
    avgRating DECIMAL(10,10),
	originalTitle VARCHAR(255),
	primaryTitle VARCHAR(255),
	PRIMARY KEY (tID),
)`
package schema

import (
	"database/sql"
	"log"
)

var users_schema = `
CREATE TABLE users (
	uid INT NOT NULL PRIMARY KEY,
    name VARCHAR(255),
    dob DATE,
    password VARCHAR(255),
	language VARCHAR(255)
)`

var users_preferences_genre_schema = `
CREATE TABLE users_pref_genre (
	uid INT NOT NULL,
	genre VARCHAR(255),
	rank INT,
	PRIMARY KEY(uid, rank),
	FOREIGN KEY(uid) REFERENCES users 
)`

var users_preferences_actor_schema = `
CREATE TABLE users_pref_actor (
	uid INT NOT NULL,
	actor VARCHAR(255),
	rank INT,
	PRIMARY KEY(uid, actor),
	FOREIGN KEY(uid) REFERENCES users 
)`

type User struct {
	Uid      int    `db:"uid"`
	Name     string `db:"name"`
	Dob      string `db:"dob"`
	Password string `db:"password"`
	Language string `db:"language"`
}

type UserPrefGenre struct {
	Uid   int    `db:"uid"`
	Genre string `db:"genre"`
	Rank  int    `db:"rank"`
}

type UserPrefActor struct {
	Uid   int    `db:"uid"`
	Actor string `db:"actor"`
	Rank  int    `db:"rank"`
}

func CreateUsersTables() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(users_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("User Table Created")

	_, err = db.Exec(users_preferences_genre_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("User's Pref Genre Table Created")

	_, err = db.Exec(users_preferences_actor_schema)
	if err != nil {
		log.Fatal(err)
	}
	println("User's Pref Actor Table Created")
}

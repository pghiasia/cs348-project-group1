package load

import (
	"database/sql"
	"log"
)

func LoadUsersTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users SELECT * FROM read_csv('./test-data/users.csv')`)
	if err != nil {
		log.Fatal(err)
	}

	println("Users Loaded")
}

func LoadUsersGenreTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users_pref_genre SELECT * FROM read_csv('./test-data/users_pref_genre.csv')`)
	if err != nil {
		log.Fatal(err)
	}

	println("Users Genre Pref Loaded")
}

func LoadUsersActorTable() {
	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`INSERT INTO users_pref_actor SELECT * FROM read_csv('./test-data/users_pref_actor.csv')`)
	if err != nil {
		log.Fatal(err)
	}

	println("Users Actor Pref Loaded")
}

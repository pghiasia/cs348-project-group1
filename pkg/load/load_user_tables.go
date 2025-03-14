package load

import (
	"database/sql"
	"log"
)

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

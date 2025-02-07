package main

import (
	"cs348-project-group1/pkg/load"
	"cs348-project-group1/pkg/schema"
	"database/sql"
	"errors"
	"fmt"
	"log"

	_ "github.com/marcboeker/go-duckdb"
)

func main() {
	//Users Table Creation & Loading
	schema.CreateUsersTables()
	load.LoadUsersTable()
	load.LoadUsersGenreTable()
	load.LoadUsersActorTable()

	//Actors Table Creation & Loading
	schema.CreateActorTable()
	load.LoadActorsTable()

	//Movies Table Creation & Loading
	schema.CreateMoviesTable()
	load.LoadMoviesTable()
	load.LoadMovieToActorTable()
	load.LoadMovieToGenreTable()

	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	user_table := schema.User{}

	rows, nil := db.Query(`SELECT * FROM users`)

	if errors.Is(err, sql.ErrNoRows) {
		log.Println("no rows")
	} else if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		_ = rows.Scan(
			&user_table.Uid, &user_table.Name,
			&user_table.Dob, &user_table.Password,
			&user_table.Language,
		)
		fmt.Printf("uid: %d, name: %s, dob: %s, password: %s, language: %s\n", user_table.Uid, user_table.Name, user_table.Dob, user_table.Password, user_table.Language)
	}
}

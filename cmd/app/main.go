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

func list_users(db *sql.DB) {
	user_table := schema.User{}

	rows, err := db.Query(`SELECT * FROM users`)

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

func list_movie_ratings(db *sql.DB) {
	println("\n\n ******** Listing Movie Ratings ******** \n\n")
	movie_table := schema.Movie{}
	rows, _ := db.Query(`SELECT title, rating FROM movies`)

	for rows.Next() {
		_ = rows.Scan(&movie_table.Title, &movie_table.Rating)
		fmt.Printf("movie %s has a rating of: %.2f\n", movie_table.Title, movie_table.Rating)
	}
}

func list_highest_rating_movie(db *sql.DB) {
	println("\n\n ******** Listing Movie with Highest Rating ******** \n\n")

	var q = `
	SELECT title, rating 
	FROM movies
	WHERE rating = (SELECT MAX(rating) FROM movies);
	`

	movie_table := schema.Movie{}
	row := db.QueryRow(q)
	_ = row.Scan(&movie_table.Title, &movie_table.Rating)
	fmt.Printf("movie %s has the highest rating of: %.2f\n", movie_table.Title, movie_table.Rating)
}


func list_highest_rating_movie_in_actor(db *sql.DB) {
	var first_name string
	var last_name string
    println("\n\n ******** Listing Movie with Highest Rating with respect to an Actor.******** \n\n")

    print("\n\nFirst Name: ")
	fmt.Scanln(&first_name)
    print("\n\nLast Name: ")
	fmt.Scanln(&last_name)

    actor := first_name + " " + last_name

    fmt.Printf("\n\n ******** Listing Movie with Highest Rating with respect to %s ******** \n\n", actor)

	type Output struct {
		Title  string
		Actor  string
		Rating float32
	}


	var q = `
	SELECT m.title, m.rating, a.name
  	FROM movies m NATURAL JOIN movie_to_actor ma NATURAL JOIN actors a
  	WHERE a.name = ?
  	ORDER BY m.rating DESC
  	LIMIT 1;
	`

	output := Output{}

	row := db.QueryRow(q, actor)
    _ = row.Scan(&output.Title, &output.Rating, &output.Actor)

	fmt.Printf("movie %s with actor %s has the highest rating of: %.2f\n", output.Title, output.Actor, output.Rating)
}

func main() {
	//Users Table Creation & Loading
	schema.CreateUsersTables()
	load.LoadUsersTable()

	//Actors Table Creation & Loading
	schema.CreateActorTable()
	load.LoadActorsTable()

	//Movies Table Creation & Loading
	schema.CreateMoviesTable()
	load.LoadMoviesTable()
	load.LoadMovieToActorTable()

	db, err := sql.Open("duckdb", "./movie.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

    list_highest_rating_movie(db);
    //	list_highest_rating_movie_in_actor(db)
}

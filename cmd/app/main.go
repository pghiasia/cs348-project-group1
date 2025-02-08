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

func list_uers(db *sql.DB) {
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
	fmt.Printf("movie %s has a the highest rating of: %.2f\n", movie_table.Title, movie_table.Rating)

}

func list_highest_rating_movie_in_genre(db *sql.DB) {
	var genre string
	print("!! Currently only supports Romance, Documentary, Short, Animation, Comedy, Sport, as we only have these in the sample data, and haven't implemented null checking yet \n Please entre the Genre: ")
	fmt.Scan(&genre)
	println("")

	fmt.Printf("\n\n ******** Listing Movie with Highest Rating in %s ******** \n\n", genre)

	type Output struct {
		Title  string
		Genre  string
		Rating float32
	}

	var q = `
	SELECT m.title, g.genrename, m.rating
  	FROM movies m
  	JOIN movie_to_genre g ON m.mid = g.mid
  	WHERE g.genrename LIKE ?
  	ORDER BY m.rating DESC
  	LIMIT 1;
	`

	output := Output{}

	row := db.QueryRow(q, "%"+genre+"%")
	_ = row.Scan(&output.Title, &output.Genre, &output.Rating)

	fmt.Printf("movie %s in %s has a the highest rating of: %.2f\n", output.Title, output.Genre, output.Rating)
}
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

	list_movie_ratings(db)
	list_highest_rating_movie(db)
	list_highest_rating_movie_in_genre(db)
}

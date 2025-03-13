package movie

import (
	"cs348-project-group1/pkg/schema"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type getMoviesRequest struct {
	Sortby string `json:"sortby"`
}

func GetMovies(c *gin.Context) {
	var requestBody getMoviesRequest

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, httpError{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid Request Body",
		})
		return
	}

	db, err := sql.Open("duckdb", "./movie.db")
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	query := fmt.Sprintf("SELECT mid, title, genres, release, rating, numVotes FROM movies ORDER BY %s DESC", requestBody.Sortby)
	rows, err := db.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	defer rows.Close()

	var movies []schema.Movie
	for rows.Next() {
		var movie schema.Movie
		err := rows.Scan(
			&movie.Mid, &movie.Title, &movie.Genres,
			&movie.Release, &movie.Rating, &movie.NumVotes,
		)
		if err != nil {
			c.JSON(http.StatusInternalServerError, httpError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}
		movies = append(movies, movie)
	}

	c.JSON(http.StatusOK, movies)
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

package movie

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func GetMovies(c *gin.Context) {
	db, err := sql.Open("duckdb", "./movie.db")
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	query := `
	WITH AllTitles AS (
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'movie' AS titleType
		FROM movie
		UNION ALL
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'series' AS titleType
		FROM series
		UNION ALL
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'short' AS titleType
		FROM short
		UNION ALL
		SELECT tID, isAdult, releaseYear, originalTitle, averageRating,
		numVotes, runtimeMinutes, primaryTitle, 'episode' AS titleType
		FROM episodes
	)
	SELECT DISTINCT a.tID, a.primaryTitle, a.releaseYear, a.averageRating, a.isAdult, a.titleType
	FROM AllTitles a
	WHERE
		tID is not null
	LIMIT 20
	`
	genre_filter := c.Query("genre")
	title_filter := c.Query("title_type")

	if genre_filter != "" {
		println(genre_filter)
	}

	if title_filter != "" {
		println(title_filter)
	}

	rows, _ := db.Query(query)
	defer rows.Close()

	type RetMovie struct {
		Tid           string  `db:"tid"`
		PrimaryTitle  string  `db:"primaryTitle"`
		ReleaseYear   int     `db:"releaseYear"`
		AverageRating float64 `db:"averageRating"`
		IsAdult       bool    `db:"isAdult"`
		TitleType     string  `db:"titleType"`
	}

	var results []RetMovie

	for rows.Next() {
		var row RetMovie
		rows.Scan(
			&row.Tid, &row.PrimaryTitle, &row.ReleaseYear,
			&row.AverageRating, &row.IsAdult, &row.TitleType,
		)
		results = append(results, row)
	}

	c.JSON(http.StatusOK, results)
}

package movie

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type updateRatingRequest struct {
	Tid    string  `json:"tid"`
	Rating float64 `json:"rating"`
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
	`
	genre_filter := c.Query("genre")
	title_filter := c.Query("titleType")
	crew_member := c.Query("crewMember")
	title_keyword := c.Query("titleKeyword")
	start_year, end_year := c.Query("startYear"), c.Query("endYear")
	low_rating, high_rating := c.Query("lowRating"), c.Query("highRating")
	is_adult := c.Query("isAdult")

	if genre_filter != "" {
		query += fmt.Sprintf("AND EXISTS (FROM genres rg SELECT 1 WHERE rg.tid = a.tID AND rg.genre = '%s')", genre_filter)
	}

	if title_filter != "" {
		query += fmt.Sprintf("AND a.titleType = '%s'", title_filter)
	}

	if crew_member != "" {
		query += fmt.Sprintf("AND EXISTS (SELECT 1 FROM workedOn w JOIN people p ON w.pID = p.pID WHERE w.tID = a.tID AND p.name = '%s')", crew_member)
	}

	if title_keyword != "" {
		query += fmt.Sprintf("AND a.originalTitle LIKE '%%%s%%'", title_keyword)
	}

	if start_year != "" && end_year != "" {
		query += fmt.Sprintf("AND a.releaseYear BETWEEN %s AND %s", start_year, end_year)
	}

	if low_rating != "" && high_rating != "" {
		query += fmt.Sprintf("AND a.averageRating BETWEEN %s AND %s", low_rating, high_rating)
	}

	if is_adult == "false" {
		query += fmt.Sprintf("AND a.isAdult = 0")
	} else {
		query += fmt.Sprint("AND a.isAdult = 1")
	}

	query += "LIMIT 1000"
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
	if results == nil {
		results = make([]RetMovie, 0)
		c.JSON(http.StatusOK, results)
		return
	}
	c.JSON(http.StatusOK, results)
}

func GetMovie(c *gin.Context) {
	db, err := sql.Open("duckdb", "./movie.db")
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	tid := c.Query("tid")

	type RetMovie struct {
		Genres []string
		Crew   []string
	}

	var genres []string
	var crew []string

	g_query := fmt.Sprintf("SELECT genre FROM genres WHERE tid = '%s'", tid)
	c_query := fmt.Sprintf("SELECT name FROM workedOn w JOIN people p ON w.pID = p.pID WHERE w.tID = '%s'", tid)

	rows, _ := db.Query(g_query)
	for rows.Next() {
		var gname string
		rows.Scan(&gname)
		genres = append(genres, gname)
	}

	rows, _ = db.Query(c_query)
	for rows.Next() {
		var cname string
		rows.Scan(&cname)
		crew = append(crew, cname)
	}

	defer rows.Close()
	results := RetMovie{
		Genres: genres,
		Crew:   crew,
	}

	c.JSON(http.StatusOK, results)
}

func UpdateRating(c *gin.Context) {
	var requestBody updateRatingRequest

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

	rating_string := strconv.FormatFloat(requestBody.Rating, 'f', -1, 64)
	update_movie_query := fmt.Sprintf("UPDATE movie SET numVotes = numVotes + 1, averageRating = ((averageRating * (numVotes - 1)) + %s) / numVotes WHERE tID = '%s'", rating_string, requestBody.Tid)
	db.Exec(update_movie_query)

	update_short_query := fmt.Sprintf("UPDATE short SET numVotes = numVotes + 1, averageRating = ((averageRating * (numVotes - 1)) + %s) / numVotes WHERE tID = '%s'", rating_string, requestBody.Tid)
	db.Exec(update_short_query)

	update_series_query := fmt.Sprintf("UPDATE series SET numVotes = numVotes + 1, averageRating = ((averageRating * (numVotes - 1)) + %s) / numVotes WHERE tID =  '%s'", rating_string, requestBody.Tid)
	db.Exec(update_series_query)

	update_episode_query := fmt.Sprintf("UPDATE episodes SET numVotes = numVotes + 1, averageRating = ((averageRating * (numVotes - 1)) + %s) / numVotes WHERE tID =  '%s'", rating_string, requestBody.Tid)
	db.Exec(update_episode_query)

	c.JSON(http.StatusOK, "rating updated")
}

package user

import (
	"database/sql"
	"net/http"

	"cs348-project-group1/pkg/schema"

	"github.com/gin-gonic/gin"
)

type httpError struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

type signupRequest struct {
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Password string `json:"password"`
	Language string `json:"language"`
}

type signinRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type getUserRequest struct {
	Name string `json:"name"`
}

type userInfo struct {
	Uid      string `json:"uid"`
	Name     string `json:"name"`
	Dob      string `json:"dob"`
	Language string `json:"language"`
}

func getNumUsers(db *sql.DB) (int, error) {
	count := 0
	err := db.QueryRow(`SELECT COUNT(*) as num_user FROM users`).Scan(&count)
	return count, err
}

func userExists(db *sql.DB, name string) (bool, error) {
	var exists bool
	query := "SELECT EXISTS(SELECT * FROM users WHERE name = ?)"

	err := db.QueryRow(query, name).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

func checkCredential(db *sql.DB, name string, password string) (bool, error) {
	var match bool
	query := "SELECT EXISTS(SELECT * FROM users WHERE name = ? AND password = sha256(?))"

	err := db.QueryRow(query, name, password).Scan(&match)
	if err != nil {
		return false, err
	}

	return match, nil
}

func GetUser(c *gin.Context) {
	var requestBody getUserRequest

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

	// check if user exists
	user_exists, err := userExists(db, requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if !user_exists {
		c.JSON(http.StatusBadRequest, httpError{
			StatusCode: http.StatusBadRequest,
			Message:    "User does not exist",
		})
		return
	}

	var user userInfo
	query := "SELECT uid, name, dob, language FROM users WHERE name = ?"
	err = db.QueryRow(query, requestBody.Name).Scan(&user.Uid, &user.Name, &user.Dob, &user.Language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {
	db, err := sql.Open("duckdb", "./movie.db")
	defer db.Close()
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	query := "SELECT uid, name, dob, language FROM users"
	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	defer rows.Close()

	var users []userInfo
	for rows.Next() {
		var user userInfo
		err := rows.Scan(&user.Uid, &user.Name, &user.Dob, &user.Language)
		if err != nil {
			c.JSON(http.StatusInternalServerError, httpError{
				StatusCode: http.StatusInternalServerError,
				Message:    err.Error(),
			})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

func SignIn(c *gin.Context) {
	var requestBody signinRequest

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

	// check if user exists
	user_exists, err := userExists(db, requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if !user_exists {
		c.JSON(http.StatusBadRequest, httpError{
			StatusCode: http.StatusBadRequest,
			Message:    "User does not exist",
		})
		return
	}

	// search for matching credentials in the db
	match, err := checkCredential(db, requestBody.Name, requestBody.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if !match {
		c.JSON(http.StatusBadRequest, httpError{
			StatusCode: http.StatusBadRequest,
			Message:    "Username or password is invalid",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": requestBody.Name,
	})
}

func SignUp(c *gin.Context) {
	var requestBody signupRequest

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

	// check if user exists
	user_exists, err := userExists(db, requestBody.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	if user_exists {
		c.JSON(http.StatusBadRequest, httpError{
			StatusCode: http.StatusBadRequest,
			Message:    "User already exists, please use a different user name",
		})
		return
	}

	// insert user to table with new UID
	num_user, err := getNumUsers(db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	user := schema.User{
		Uid:      num_user + 1,
		Name:     requestBody.Name,
		Dob:      requestBody.Dob,
		Language: requestBody.Language,
		Password: requestBody.Password,
	}

	stmt, err := db.Prepare("INSERT INTO users(uid, name, dob, password, language) VALUES (?, ?, ?, sha256(?), ?)")
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Uid, user.Name, user.Dob, user.Password, user.Language)
	if err != nil {
		c.JSON(http.StatusInternalServerError, httpError{
			StatusCode: http.StatusInternalServerError,
			Message:    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"name": user.Name,
	})
}

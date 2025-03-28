package web

import (
	"cs348-project-group1/web/movie"
	"cs348-project-group1/web/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(cors.Default())
	r.POST("/user/signup", user.SignUp)
	r.POST("/user/signin", user.SignIn)
	r.GET("/user", user.GetUser)
	r.GET("/users", user.GetUsers)
	r.GET("/movies", movie.GetMovies)
	return r
}

package main

import (
	"cs348-project-group1/pkg/load"
	"cs348-project-group1/pkg/schema"
	"cs348-project-group1/web"

	_ "github.com/marcboeker/go-duckdb"
)

func initDb() {
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
}

func main() {
	initDb()
	r := web.InitRouter()
	r.Run(":9888")
}

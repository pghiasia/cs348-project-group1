package main

import (
	"cs348-project-group1/pkg/load"
	"cs348-project-group1/pkg/schema"
	"cs348-project-group1/web"

	_ "github.com/marcboeker/go-duckdb"
)

func initTestDb() {
	//Users Table Creation & Loading
	schema.CreateUsersTables()
	load.LoadUsersTable("./test-data/users.csv")

	//Actors Table Creation & Loading
	schema.CreateActorTable()
	load.LoadActorsTable()

	//Movies Table Creation & Loading
	schema.CreateMoviesTable()
	load.LoadMoviesTable()
	load.LoadMovieToActorTable()
}

func initProdDb() {
	//Users Table Creation & Loading
	schema.CreateUsersTables()
	load.LoadUsersTable("./test-data/usersProd.csv")

	//Actors Table Creation & Loading
	schema.CreateActorTable()
	load.LoadActorsTable()

	//Movies Table Creation & Loading
	schema.CreateMoviesTable()
	load.LoadMoviesTable()
	load.LoadMovieToActorTable()
}


func main() {
	initTestDb()
	r := web.InitRouter()
	r.Run(":9888")
}

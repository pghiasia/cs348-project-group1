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

    load.LoadTitlesTable("./test-data/title.basics.tsv")
    load.LoadPeopleTable("./test-data/name.basics.tsv")
    load.LoadEpisodesTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv", "./test-data/title.episode.tsv")
    load.LoadSeriesTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadShortTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadMovieTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadWorkedOnTable("./test-data/title.principals.tsv")

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

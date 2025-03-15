package main

import (
	"cs348-project-group1/pkg/load"
	"cs348-project-group1/pkg/schema"
	"cs348-project-group1/web"

	_ "github.com/marcboeker/go-duckdb"
)

// Initalizes the Production Database, run this to access production database. 
func initProdDb() {
    schema.CreateTables()
    load.LoadTitlesTable("./production-data/title.basics.tsv")
    load.LoadPeopleTable("./production-data/name.basics.tsv")
	load.LoadUsersTable("./production-data/usersProd.csv")
    load.LoadSeriesTable("./production-data/title.basics.tsv", "./production-data/title.ratings.tsv")
    load.LoadShortTable("./production-data/title.basics.tsv", "./production-data/title.ratings.tsv")
    load.LoadMovieTable("./production-data/title.basics.tsv", "./production-data/title.ratings.tsv")
    load.LoadEpisodesTable("./production-data/title.basics.tsv", "./production-data/title.ratings.tsv", "./production-data/title.episode.tsv")
    load.LoadGenresTable("./production-data/title.basics.tsv")
    load.LoadWorkedOnTable("./production-data/title.principals.tsv")
}

// Initalizes the Test Database, run this to access sample data. 
func initTestDb() {
    schema.CreateTables()
    load.LoadTitlesTable("./test-data/title.basics.tsv")
    load.LoadPeopleTable("./test-data/name.basics.tsv")
	load.LoadUsersTable("./test-data/users.csv")
    load.LoadSeriesTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadShortTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadMovieTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadEpisodesTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv", "./test-data/title.episode.tsv")
    load.LoadGenresTable("./test-data/title.basics.tsv")
    load.LoadWorkedOnTable("./test-data/title.principals.tsv")
}


func main() {
	initProdDb()
	r := web.InitRouter()
	r.Run(":9888")
}

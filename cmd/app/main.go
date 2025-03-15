package main

import (
	"cs348-project-group1/pkg/load"
	"cs348-project-group1/pkg/schema"
	"cs348-project-group1/web"

	_ "github.com/marcboeker/go-duckdb"
)

func initProdDb() {
    schema.CreateTables()
    load.LoadTitlesTable("./bigData/title.basics.tsv")
    load.LoadPeopleTable("./bigData/name.basics.tsv")
	load.LoadUsersTable("./test-data/usersProd.csv")
    load.LoadSeriesTable("./bigData/title.basics.tsv", "./bigData/title.ratings.tsv")
    load.LoadShortTable("./bigData/title.basics.tsv", "./bigData/title.ratings.tsv")
    load.LoadMovieTable("./bigData/title.basics.tsv", "./bigData/title.ratings.tsv")
    load.LoadEpisodesTable("./bigData/title.basics.tsv", "./bigData/title.ratings.tsv", "./bigData/title.episode.tsv")
    load.LoadGenresTable("./bigData/title.basics.tsv")
    load.LoadWorkedOnTable("./bigData/title.principals.tsv")
}

func initTestDb() {
    schema.CreateTables()
    load.LoadTitlesTable("./test-data/title.basics.tsv")
    load.LoadPeopleTable("./test-data/name.basics.tsv")
	load.LoadUsersTable("./test-data/usersProd.csv")
    load.LoadSeriesTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadShortTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadMovieTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv")
    load.LoadEpisodesTable("./test-data/title.basics.tsv", "./test-data/title.ratings.tsv", "./test-data/title.episode.tsv")
    load.LoadGenresTable("./test-data/title.basics.tsv")
    load.LoadWorkedOnTable("./test-data/title.principals.tsv")
}


func main() {
	initTestDb()
	r := web.InitRouter()
	r.Run(":9888")
}

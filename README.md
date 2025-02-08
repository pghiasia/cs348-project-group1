# cs348-project-group1

The Repository for Group #1's CS 348 Winter 2025 Project.

## Milestone 0

- To access the progress pdf for milestone 0, please consult the README file under `docs/milestone-0` directory
- To run the sample application, in the project directory, run `go mod download` and `go run main.go`

## Milestone 1

### Setup

- We have switched the database we use to **duckdb**.
- To setup the go environment, please ensure **go is installed** and run **`go mod download`**
- To get started with our sample appliation, the table initlization and data loading could be done with `go run cmd/app/main.go`

### Basic Movie Feature (Implemneted via function calls)

- The `list_movie_ratings` function created in the main file implements the feature where the users query all movies with their ratings
- The `list_highest_rating_movie` function created in the main file implements the feature the user wants to find the movie with the highest rating
- The `list_highest_rating_movie_in_genre` function created in the main file implementes the feature where the user wants to find the movie with the highest rating in a specific genre`
  - This function takes a user input from the command line of a genre that we currently support

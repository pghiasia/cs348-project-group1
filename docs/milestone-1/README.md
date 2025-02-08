## Milestone 1
The link to the milestone 1 report can be found here: [report.pdf](https://docs.google.com/document/d/1e3PryBrJ6RT2bxPib76dL9ANRlInVWm12EVRgzta7_g/edit?usp=sharing)

### Setup

- We have switched the database we use to **duckdb**, please make sure you have it installed.
- To setup the go environment, please ensure **go is installed** and run **`go mod download`**
- To get started with our sample appliation, the table initlization and data loading could be done with `go run cmd/app/main.go`  

note: running the application will automatically load the data.  
### Basic Features (Implemented via function calls)

Main features:  
R6: list_highest_rating_movie(db *sql.DB)  
    - finds the movie with the max rating.  
R9: list_highest_rating_movie_in_actor(db *sql.DB)  
    - finds the highest ranking movie which involved a specific actor.

note: the top level features directory contains the SQL queries and outputs for functions R6, R7, R8, R9.

Miscellaneous to be possibly used in the future, not part of milestone 1:
- The `list_movie_ratings` function created in the main file implements the feature where the users query all movies with their ratings.
- The `list_users` function created in the main file prints out all the users and their info.

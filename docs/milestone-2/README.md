# Milestone 2

The link to the milestone 2 report can be found here: [report.pdf](https://docs.google.com/document/d/1B7Vtl7qic19OT3wTG5EU1CMJkK214ILyDAvyS6Z2oUo/edit?usp=sharing)

## Changes

- For this milstone, we decided to spin up a server with Go using gin, and moved our features into API endpoints.
- We have also introduced React for our GUI and frontend.

## Setup

- We have switched the database we use to **duckdb**, please make sure you have it installed.
- To setup the go environment, please ensure **go is installed** and run **`go mod download`**
- To be able to run our GUI, please ensure that the latest version of [node](https://nodejs.org/en/download) and [npm](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) are both installed.

## Running The App

- To start up the backend server, run `go run cmd/app/main.go` in the root directory.
- To start up the GUI, run `npm i` and `npm run dev` in the `/app` directory.
- note: running the application will automatically load the data.  

## How To Run The App With: Production Data and Sample Data
- In cmd/app/main.go there are two functions, initProdDb, and initTestDb,
  by default the database is loaded with the production database (the main function calls initProdDb), 
  to change to sample data just replace the call to initProdDb in main to initTestDb.

## Basic Features (Implemented via function calls)
note: the top level features directory contains the SQL queries and outputs for functions R6, R7, R8, R9, R10.
You can learn more about these features in the report. 

## GUI
- Login Page implemented with user signup and signin functions, password & username verifications, and required field checks.
![demo-1](./assets/login-1.png)
![demo-2](./assets/login-2.png)
![demo-3](./assets/login-3.png)

- The Movie Database page currently only supports display the movies and sorting them with a specific field. We can click into each movie card to checkout more details for the movie. More features are currently work in progress for the GUI to support the API calls.
![demo-4](./assets/movie-1.png)
![demo-5](./assets/movie-2.png)
![demo-6](./assets/movie-3.png)

Miscellaneous to be possibly used in the future, not part of milestone 1:
- The `list_movie_ratings` function created in the main file implements the feature where the users query all movies with their ratings.
- The `list_users` function created in the main file prints out all the users and their info.

# cs348-project-group1

The Repository for Group #1's CS 348 Winter 2025 Project. 

## Milestone 0

- To access the progress pdf for milestone 0, please consult the README file under `docs/milestone-0` directory
- To run the sample application, in the project directory, run `go mod download` and `go run main.go`

## Setup

- To run the database and the sample application, ensure that [Docker](https://www.docker.com/get-started/) is installed.
- In the `setup/` directory, run `docker compose up -d` or `docker-compose up -d` if you are running an older version of docker compose.
- Postgres is running at `localhost:5432` and pgAdmin is running at `localhost:5433`
  - Postgres Credentials:
    - User: postgres
    - Password: cs348
  - pgAdmin Credentials:
    - Email: placeholder@email.com
    - Password: cs348

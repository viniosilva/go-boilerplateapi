# Go Boilerplate API

A Go-based API boilerplate for APIs

## Tech Stack

- [Go](https://go.dev/)
- [Make](https://www.gnu.org/software/make/)
- [Project Layout](https://github.com/golang-standards/project-layout)
- [Swagger](https://swagger.io/)
- [Docker](https://www.docker.com/)
- [PostgreSQL](https://www.postgresql.org/)

## Run Locally

Install dependencies

```bash
  make
```

Configure variables environment

```bash
  cp .env.example .env
```

Start the server

```bash
  make dev
```

Local swagger [[here]](http://localhost:8000/swagger/index.html)


## Running Tests

To run tests, run the following command

```bash
  # unit tests
  make test

  # e2e tests
  make test-e2e
```

## Deployment

[UNDER CONSTRUCTION]

To deploy this project run

```bash
  make deploy
```
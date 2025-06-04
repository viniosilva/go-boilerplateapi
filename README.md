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

Start infra

```
  make infra-up
```

Configure data sources using docker compose network:

- Loki doc [[here]](https://grafana.com/docs/grafana/latest/datasources/loki/configure-loki-data-source/)
- Tempo doc [[here]](https://grafana.com/docs/grafana/latest/datasources/tempo/configure-tempo-data-source/)
- Prometheus doc [[here]](https://grafana.com/docs/grafana/latest/datasources/prometheus/configure-prometheus-data-source/)

Start the server

```bash
  make dev
```

### Links

- Local swagger UI [[here]](http://localhost:8000/swagger/index.html)
- Local grafana UI [[here]](http://localhost:3000)
- Local jaeger UI [[here]](http://localhost:9090)
- Local prometheus UI [[here]](http://localhost:16686)

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
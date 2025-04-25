include .env
export

MIGRATE_CMD := docker compose run --rm migrate
MIGRATE_URL := postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@db/$(POSTGRES_DB)?sslmode=disable
MIGRATE_PATH := -path=/migrations/

DOCKER_COMPOSE_TEST_CMD := docker compose -f docker-compose.yml -f test/docker-compose.override.yml

all: install-migrate install-swag install-lint install-lefthook install-hot-reload install-mock

install-migrate:
	go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest
install-lint:
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.1.2
install-lefthook:
	go install github.com/evilmartians/lefthook@latest
install-hot-reload:
	go install github.com/air-verse/air@latest
install-mock:
	go install go.uber.org/mock/mockgen@latest

dev: migrate-up
	docker compose down api
	air

dev-docker: migrate-up
	docker compose up -d

.PHONY: test
test:
	go test -cover -v \
		./internal/application/... \
		./internal/domain/... \
		./internal/presentation/...

test-e2e: test-up
	# go test -v ./test/...
	make test-down

test-up:
	$(DOCKER_COMPOSE_TEST_CMD) run --rm migrate $(MIGRATE_PATH) -database="$(MIGRATE_URL)" up
	$(DOCKER_COMPOSE_TEST_CMD) up -d

test-down:
	$(DOCKER_COMPOSE_TEST_CMD) down --volumes

migrate-up:
	$(MIGRATE_CMD) $(MIGRATE_PATH) -database="$(MIGRATE_URL)" up

migrate-down:
	$(MIGRATE_CMD) $(MIGRATE_PATH) -database="$(MIGRATE_URL)" down

migrate-force:
	$(MIGRATE_CMD) $(MIGRATE_PATH) -database="$(MIGRATE_URL)" force $(version)

migrate-version:
	$(MIGRATE_CMD) $(MIGRATE_PATH) -database="$(MIGRATE_URL)" version

migrate-create:
	migrate create -ext sql -dir internal/infrastructure/migrations $(name)

up: migrate-up
	docker compose up -d
	docker compose logs api

down:
	docker compose down

swag:
	swag init --parseDependency --parseInternal --generalInfo cmd/api/main.go --output docs/

lint:
	golangci-lint run ./...

lint-fix:
	golangci-lint run --fix ./...

.PHONY: mock
mock:
	mockgen -source=internal/domain/customer/customer_repository.go -destination=mock/customer_repository_mock.go -package=mock


precommit:
	lefthook run pre-commit
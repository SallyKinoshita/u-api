include .env

init:
	docker compose down -v
	@make up

rebuild-up:
	@make down
	@make build
	@make up

reup:
	@make down
	@make up

restart:
	@make stop
	@make start

build:
	docker compose build

up:
	docker compose up -d

start:
	docker compose start

down:
	docker compose down

stop:
	docker compose stop

ps:
	docker compose ps

exec-api:
	docker compose exec $(API_HOST) /bin/bash

go-check:
	@make go-fmt
	@make go-lint

go-fmt:
	docker compose exec $(API_HOST) go fmt ./...

go-lint:
	docker compose exec $(API_HOST) sh -c 'go list ./... | grep -v internal/gen | grep -v internal/di | xargs staticcheck -f stylish'

go-mod-tidy:
	docker compose exec $(API_HOST) go mod tidy

go-test:
	docker compose exec $(API_HOST) go test ./internal/... -v -cover -coverprofile=../cover.out

go-test-coverage:
	docker compose exec $(API_HOST) go tool cover -html=../cover.out -o cover.html

go-update-schema:
	oapi-codegen -package "openapi" docs/openapi.yml > internal/gen/openapi/doc.go

db-migrate-diff:
	docker pull arigaio/atlas:latest
	docker run --rm --net=host -v $(CURDIR):/app arigaio/atlas \
		migrate diff $(ARG) \
			--dir "file://app/migrations" \
			--to "file://app/atlas.hcl" \
			--dev-url "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@127.0.0.1:$(MYSQL_PORT)/$(MYSQL_DATABASE)" \
			--format '{{ sql . "  " }}'

db-migrate-apply:
	docker pull arigaio/atlas:latest
	docker run --rm --net=host -v $(CURDIR):/app arigaio/atlas \
		migrate apply --dir "file://app/migrations" --url "mysql://$(MYSQL_USER):$(MYSQL_PASSWORD)@127.0.0.1:$(MYSQL_PORT)/$(MYSQL_DATABASE)"
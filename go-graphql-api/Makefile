include .env

.DEFAULT_GOAL := help
.PHONY: build help migrate-up migrate-down seed migrate-create

build: ## Build app
	@go build -o build/gql-server cmd/gql-server/main.go

migrate-up: ## Migrations up
	@migrate -database ${DB_URL} -path db/migrations up

migrate-down: ## Migrations down
	@migrate -database ${DB_URL} -path db/migrations down

seed: ## Seed DB
	@go run scripts/seeder.go

migrate-create: ## Create migration (for documentation purposes)
	@migrate create -ext sql -dir db/migrations -seq template

## broken because of include (Makefile prefix in MAKEFILE_LIST)
help: ## This help dialog.
	@egrep '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

MIGRATE=./migrate.darwin-amd64 -path db/migrations -database postgres://postgres:manager@localhost:5432/task_manager?sslmode=disable
TEST_MIGRATE=./migrate.darwin-amd64 -path db/migrations -database postgres://postgres:manager@localhost:5432/test_task_manager?sslmode=disable

.PHONY: build start stop down migrate-up migrate-down

build: ## Build docker containers
	docker-compose build

start: ## Start docker containers
	docker-compose up -d

stop: ## Stop docker containers
	docker-compose stop

down: ## Down docker containers
	docker-compose down

migrate-up: ## Run migrations
	$(MIGRATE) up

migrate-down: ## Rollback migrations
	$(MIGRATE) down

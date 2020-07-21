MIGRATE=./migrate -path db/migrations -database postgres://postgres:manager@db:5432/task_manager?sslmode=disable
TEST_MIGRATE=./migrate.darwin-amd64 -path db/migrations -database postgres://postgres:manager@localhost:5432/test_task_manager?sslmode=disable

.PHONY: build start stop down migrate-up migrate-down test

build: ## Build docker containers
	docker-compose build

start: ## Start docker containers
	docker-compose up -d

stop: ## Stop docker containers
	docker-compose stop

down: ## Down docker containers
	docker-compose down --volumes

migrate-up: ## Run migrations
	docker-compose exec app $(MIGRATE) up

migrate-down: ## Rollback migrations
	docker-compose exec app $(MIGRATE) down

test: ## Request test
	docker-compose -f docker-compose.test.yml up -d --build
	sleep 2
	$(TEST_MIGRATE) up
	go test -v ./...
	docker-compose -f docker-compose.test.yml down --volumes
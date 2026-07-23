.PHONY: build run test clean dev migrate seed docker-build docker-up lint

APP_NAME=satudata-api
BUILD_DIR=./bin

# ========== Development ==========

dev: ## Run development server with hot reload
	@echo "Starting development server..."
	@if command -v air > /dev/null; then \
		air; \
	else \
		go run ./cmd/api/main.go; \
	fi

build: ## Build binary
	@echo "Building $(APP_NAME)..."
	@mkdir -p $(BUILD_DIR)
	CGO_ENABLED=0 go build -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/main.go
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

run: build ## Build and run
	@echo "Running $(APP_NAME)..."
	@$(BUILD_DIR)/$(APP_NAME)

test: ## Run all tests
	@echo "Running tests..."
	go test -v -race -count=1 ./...

test-coverage: ## Run tests with coverage
	@echo "Running tests with coverage..."
	go test -v -race -count=1 -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

lint: ## Run linter
	@echo "Running linter..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run ./...; \
	else \
		echo "golangci-lint not installed. Install with: brew install golangci-lint"; \
	fi

clean: ## Clean build artifacts
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# ========== Database ==========

migrate: ## Run database migrations
	@echo "Running migrations..."
	@echo "Create tables manually by running SQL files in ./migrations/"
	@echo "Or use: cat migrations/*.sql | psql -h localhost -U satudata -d satudata"

migrate-up: ## Run migrations (using psql)
	@for f in migrations/*.sql; do \
		echo "Running migration: $$f"; \
		psql "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -f $$f; \
	done

seed: ## Seed database with sample data
	go run ./cmd/seed/main.go

# ========== Docker ==========

docker-build: ## Build Docker image
	@echo "Building Docker image..."
	docker build -t $(APP_NAME):latest .

docker-up: ## Start all services with Docker Compose
	@echo "Starting Docker Compose services..."
	docker compose up -d

docker-down: ## Stop all services
	@echo "Stopping Docker Compose services..."
	docker compose down

docker-logs: ## View Docker Compose logs
	docker compose logs -f

# ========== Production ==========

build-prod: ## Build for production (linux amd64)
	@echo "Building production binary..."
	@mkdir -p $(BUILD_DIR)
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BUILD_DIR)/$(APP_NAME) ./cmd/api/main.go
	@echo "Production build complete: $(BUILD_DIR)/$(APP_NAME)"

# ========== Help ==========

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# Variables
BINARY_NAME=job-tracker
MAIN_PATH=cmd/service/main.go
DOCKER_COMPOSE=docker compose
GO=go

# Colors for help
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: help run build test clean fmt vet lint docker-build docker-up docker-down docker-restart docker-logs db-up db-down db-restart db-logs

# Default target
.DEFAULT_GOAL := help

## help: Show this help message
help:
	@echo "$(BLUE)Available targets:$(NC)"
	@echo ""
	@echo "$(BLUE)Development:$(NC)"
	@grep -E '^## [a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'
	@echo ""
	@echo "$(BLUE)Examples:$(NC)"
	@echo "  make run          # Run the application locally"
	@echo "  make build        # Build the binary"
	@echo "  make docker-up    # Start all services with Docker"

## run: Run the application locally
run:
	$(GO) run $(MAIN_PATH)

## build: Build the application binary
build:
	$(GO) build -o $(BINARY_NAME) $(MAIN_PATH)

## test: Run tests
test:
	$(GO) test -v ./...

## clean: Remove build artifacts
clean:
	rm -f $(BINARY_NAME)
	$(GO) clean

## fmt: Format Go code
fmt:
	$(GO) fmt ./...

## vet: Run go vet
vet:
	$(GO) vet ./...

## lint: Run golangci-lint (if installed)
lint:
	@if command -v golangci-lint >/dev/null 2>&1; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it from https://golangci-lint.run/"; \
	fi

## mod-download: Download Go dependencies
mod-download:
	$(GO) mod download

## mod-tidy: Tidy Go modules
mod-tidy:
	$(GO) mod tidy

## mod-verify: Verify Go modules
mod-verify:
	$(GO) mod verify

# Docker targets
## docker-build: Build Docker images
docker-build:
	$(DOCKER_COMPOSE) build

## docker-up: Start all services with Docker Compose
docker-up:
	$(DOCKER_COMPOSE) up -d

## docker-down: Stop all Docker Compose services
docker-down:
	$(DOCKER_COMPOSE) down

## docker-restart: Restart all Docker Compose services
docker-restart: docker-down docker-up

## docker-logs: View logs from all services
docker-logs:
	$(DOCKER_COMPOSE) logs -f

## docker-logs-app: View logs from app service only
docker-logs-app:
	$(DOCKER_COMPOSE) logs -f app

## docker-logs-db: View logs from database service only
docker-logs-db:
	$(DOCKER_COMPOSE) logs -f db

## docker-ps: Show running Docker Compose services
docker-ps:
	$(DOCKER_COMPOSE) ps

# Database targets
## db-up: Start only the database service
db-up:
	$(DOCKER_COMPOSE) up -d db

## db-down: Stop the database service
db-down:
	$(DOCKER_COMPOSE) stop db

## db-restart: Restart the database service
db-restart: db-down db-up

## db-logs: View database service logs
db-logs:
	$(DOCKER_COMPOSE) logs -f db

## db-shell: Access PostgreSQL shell
db-shell:
	$(DOCKER_COMPOSE) exec db psql -U myuser -d job_tracker

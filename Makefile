#!/bin/zsh

install:
	@echo "Started installing process."
	@go mod tidy
	@go install
	@echo "Finished installed."

up:
	@go run cmd/migrate/main.go up

drop:
	@go run cmd/drop/main.go up

down:
	@go run cmd/migrate/main.go down

migration:
	@migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))
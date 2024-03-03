.PHONY: run
run: build
	@./bin/main

.PHONY: build
build:
	@tailwindcss build -i view/css/input.css -o public/style.css --minify
	@go run github.com/a-h/templ/cmd/templ@latest generate
	@go build -o bin/main cmd/main.go

.PHONY: air
air:
	@go run github.com/cosmtrek/air@latest

.PHONY: templ-fmt
templ-fmt:
	@go run github.com/a-h/templ/cmd/templ@latest fmt view

.PHONY: create-migration
create-migration:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir ./migrations create $(name) sql

.PHONY: apply-migrations
apply-migrations:
	go run github.com/pressly/goose/v3/cmd/goose@latest -dir ./migrations sqlite3 chat.db up

.PHONY: sqlc-gen
sqlc-gen:
	go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate
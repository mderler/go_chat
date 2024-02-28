run: build
	@./bin/main

build:
	@go run github.com/a-h/templ/cmd/templ@latest generate
	@go build -o bin/main cmd/main.go

air:
	@go run github.com/cosmtrek/air@latest
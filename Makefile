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
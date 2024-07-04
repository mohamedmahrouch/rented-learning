.PHONY: run build watch

include .env


run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

# watch:
# 	CompileDaemon --build="go build -o bin/server cmd/server/main.go" --command=./bin/server
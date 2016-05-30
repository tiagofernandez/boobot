PORT := 7000

all: help

help:
	@echo - format: formats the whole source code
	@echo - run: launches the web server listening on port ${PORT}

format:
	gofmt -w **/*.go

run: format
	go build -o bin/boobot src/boobot.go
	./bin/boobot ${PORT}

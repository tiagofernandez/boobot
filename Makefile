PORT := 7000

all: run

help:
	@echo - setup: download and install packages and dependencies
	@echo - format: run code formatting on package sources
	@echo - run (default): launches the web server listening on port ${PORT}

setup:
	go get -u github.com/kataras/iris/iris

format:
	gofmt -w **/*.go

run: format
	go build -o bin/boobot src/boobot.go
	./bin/boobot ${PORT}

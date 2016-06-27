all: run

install:
	go get github.com/constabulary/gb/...

update:
	gb vendor update --all

format:
	gofmt -s -w src

test: format
	gb test -v

run: format
	gb build
	./bin/boobot 7000

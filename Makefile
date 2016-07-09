NAME = boobot
TAG = latest
PORT = 6970

all: run

clean:
	rm -rf bin

clean-all: clean
	rm -rf pkg

install:
	go get github.com/constabulary/gb/...

update: install
	gb vendor update --all

build: install clean
	gb build

build-all: clean-all build

format:
	gofmt -s -w src

test: build format
	gb test -v

run: format build
	./bin/$(NAME) $(PORT)

docker-clean:
	docker ps -a -f status=exited -q | xargs docker rm -f
	docker images | grep '^<none>' | awk '{print $3}' | xargs docker rmi -f

docker-image:
	docker build -t $(NAME):$(TAG) .

docker-shell:
	docker run -v $(shell pwd)/src:/code/src -p $(PORT):$(PORT) -ti $(NAME):$(TAG) /bin/bash

docker-run:
	docker run -p $(PORT):$(PORT) -d $(NAME):$(TAG)

docker-debug:
	docker exec -ti $(shell docker ps | grep $(NAME):$(TAG) | awk '{print $$1}') /bin/bash

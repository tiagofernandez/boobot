NAME = goatbot
TAG = latest
PORT = 6970

all: dev

# Delete binaries
clean:
	rm -rf bin

# Delete package compilations
clean-all: clean
	rm -rf pkg

# Install build tool
install:
	go get github.com/constabulary/gb/...

# Update vendor packages
update: install
	gb vendor update --all

# Build the application
build: install clean
	gb build

# Build the whole project from scratch
build-all: clean-all build

# Format source code
format:
	gofmt -s -w src

# Test source code
test: build format
	gb test -v

# Run the application server
run:
	./bin/$(NAME) $(PORT)

# Rebuild and run
dev: format build run

# Delete unused containers and untagged images
docker-clean:
	docker ps -a -f status=exited -q | xargs docker rm -f
	docker images | grep '^<none>' | awk '{print $3}' | xargs docker rmi -f
	rm -f /tmp/$(NAME).tar

# Build this project's image
docker-image:
	docker build -t $(NAME):$(TAG) .

# Save image to a temporary file
docker-save:
ifeq (, $(shell docker images -q $(NAME):$(TAG)))
	$(MAKE) docker-image
endif
	test -f /tmp/$(NAME).tar || docker save -o /tmp/$(NAME).tar $(NAME):$(TAG)

# Launch a container within the host's source code, and run a shell
docker-shell:
	docker run -v $(shell pwd)/src:/code/src -p $(PORT):$(PORT) -ti $(NAME):$(TAG) /bin/bash

# Launch a detached container, and run the application server
docker-run:
	docker run -p $(PORT):$(PORT) -d $(NAME):$(TAG)

# Debug the running container
docker-debug:
	docker exec -ti $(shell docker ps | grep $(NAME):$(TAG) | awk '{print $$1}') /bin/bash

# Deploy to production
deploy: docker-save
	ansible-playbook -i playbooks/hosts playbooks/deploy.yml

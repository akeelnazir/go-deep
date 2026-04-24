.PHONY: build run test clean docker docker-run help

help:
	@echo "OS Process - Process Management Operating System"
	@echo "Available targets:"
	@echo "  make build       - Build the kernel binary"
	@echo "  make run         - Run the kernel locally"
	@echo "  make test        - Run tests"
	@echo "  make docker      - Build Docker image"
	@echo "  make docker-run  - Run in Docker container"
	@echo "  make clean       - Clean build artifacts"

build:
	mkdir -p bin
	go build -o bin/os-process ./cmd/kernel

run: build
	./bin/os-process

test:
	go test -v ./...

docker:
	docker build -t os-process .

docker-run: docker
	docker run -it os-process

clean:
	rm -rf bin/
	docker rmi os-process 2>/dev/null || true

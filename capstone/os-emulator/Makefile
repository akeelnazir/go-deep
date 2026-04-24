.PHONY: build run test clean docker docker-run help

help:
	@echo "OS Emulator - Go-Based Operating System Simulator"
	@echo "Available targets:"
	@echo "  make build       - Build the emulator binary"
	@echo "  make run         - Run the emulator locally"
	@echo "  make test        - Run tests"
	@echo "  make docker      - Build Docker image"
	@echo "  make docker-run  - Run in Docker container"
	@echo "  make clean       - Clean build artifacts"

build:
	mkdir -p bin
	go build -o bin/os-emulator ./cmd/emulator

run: build
	./bin/os-emulator

test:
	go test -v ./...

docker:
	docker build -t os-emulator .

docker-run: docker
	docker run -it os-emulator

clean:
	rm -rf bin/
	docker rmi os-emulator 2>/dev/null || true

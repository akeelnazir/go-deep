.PHONY: build run test clean docker docker-run help

help:
	@echo "OS Minimal - Bootloader & Kernel"
	@echo "Available targets:"
	@echo "  make build       - Build the kernel binary"
	@echo "  make run         - Run the kernel locally"
	@echo "  make test        - Run tests"
	@echo "  make docker      - Build Docker image"
	@echo "  make docker-run  - Run in Docker container"
	@echo "  make clean       - Clean build artifacts"

build:
	mkdir -p bin
	go build -o bin/os-minimal ./cmd/kernel

run: build
	./bin/os-minimal

test:
	go test -v ./...

docker:
	docker build -t os-minimal .

docker-run: docker
	docker run -it os-minimal

clean:
	rm -rf bin/
	docker rmi os-minimal 2>/dev/null || true

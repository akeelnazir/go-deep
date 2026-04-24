.PHONY: build run test clean docker docker-run disk-image help

help:
	@echo "OS Baremetal - Real Bootable Operating System"
	@echo "Available targets:"
	@echo "  make build       - Build the kernel binary"
	@echo "  make disk-image  - Create bootable disk image"
	@echo "  make run         - Run in QEMU locally"
	@echo "  make test        - Run tests"
	@echo "  make docker      - Build Docker image with QEMU"
	@echo "  make docker-run  - Run in Docker container with QEMU"
	@echo "  make clean       - Clean build artifacts"

build:
	mkdir -p bin
	go build -o bin/os-baremetal ./cmd/kernel

disk-image: build
	mkdir -p disk
	dd if=/dev/zero of=disk/os.img bs=1M count=10
	mkfs.ext2 -F disk/os.img
	mkdir -p disk/mnt
	sudo mount -o loop disk/os.img disk/mnt
	sudo mkdir -p disk/mnt/boot
	sudo cp bin/os-baremetal disk/mnt/boot/kernel
	sudo umount disk/mnt
	@echo "Disk image created at disk/os.img"

run: build
	./bin/os-baremetal

run-qemu: disk-image
	qemu-system-x86_64 -drive format=raw,file=disk/os.img -m 256 -nographic

test:
	go test -v ./...

docker:
	docker build -t os-baremetal .

docker-run: docker
	docker run -it --rm os-baremetal

clean:
	rm -rf bin/ disk/
	docker rmi os-baremetal 2>/dev/null || true

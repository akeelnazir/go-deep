FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum* ./
RUN go mod download || true

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o os-baremetal ./cmd/kernel

FROM qemu/qemu:latest

RUN apt-get update && apt-get install -y \
    e2fsprogs \
    util-linux \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /root/

COPY --from=builder /app/os-baremetal .

RUN dd if=/dev/zero of=os.img bs=1M count=10 && \
    mkfs.ext2 -F os.img && \
    mkdir -p mnt && \
    mount -o loop os.img mnt && \
    mkdir -p mnt/boot && \
    cp os-baremetal mnt/boot/kernel && \
    umount mnt

ENTRYPOINT ["qemu-system-x86_64", "-drive", "format=raw,file=os.img", "-m", "256", "-nographic"]

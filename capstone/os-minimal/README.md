# OS Minimal: Bootloader & Kernel

A minimal operating system implementation featuring a simple bootloader, basic kernel with interrupt handling, and a shell interface.

## Learning Objectives

- Understand bootloader concepts and kernel initialization
- Implement interrupt handling and exception management
- Build a simple command-line shell
- Learn basic memory management
- Understand the boot sequence

## Architecture

### Components

1. **Bootloader**: Initializes hardware and loads the kernel
2. **Kernel**: Core OS functionality with interrupt handling
3. **Memory Manager**: Basic memory allocation and management
4. **Shell**: Command interpreter and execution engine
5. **System Calls**: Interface between user programs and kernel

## Features

- Bootloader simulation with hardware initialization
- Interrupt vector table and exception handling
- Basic memory management (heap allocation)
- Simple shell with built-in commands
- Process creation and termination
- System information display

## Building and Running

### Local Build
```bash
go build -o os-minimal ./cmd/kernel
./os-minimal
```

### Docker Build and Run
```bash
docker build -t os-minimal .
docker run -it os-minimal
```

### Using Makefile
```bash
make build      # Build the kernel
make run        # Run locally
make docker     # Build Docker image
make docker-run # Run in Docker container
make clean      # Clean build artifacts
```

## Shell Commands

- `help` - Display available commands
- `echo <text>` - Print text to console
- `info` - Display system information
- `memory` - Show memory statistics
- `ps` - List running processes
- `exit` - Shutdown the OS

## Implementation Details

### Boot Sequence
1. Bootloader initializes CPU and memory
2. Kernel sets up interrupt handlers
3. Memory manager initializes heap
4. Shell starts and waits for input

### Memory Layout
```
0x0000 - 0x0FFF: Bootloader
0x1000 - 0x7FFF: Kernel
0x8000 - 0xFFFF: User space
```

### Interrupt Handling
- Division by zero (INT 0)
- Page fault (INT 14)
- System call (INT 80)

## Concepts Demonstrated

- Hardware abstraction
- Interrupt handling
- Memory management
- Process management basics
- Shell/REPL implementation
- System call interface

## Testing

```bash
go test -v ./...
```

## Resources

- Day 18: File and IO Operations
- Day 21: Signal Handling and Processes
- Day 25: Memory Management and Unsafe
- Day 28: Advanced Go - Memory, Unsafe Pointers, CGO, Data Structures, and Functional Patterns

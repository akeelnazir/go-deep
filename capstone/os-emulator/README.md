# OS Emulator: Go-Based Operating System Simulator

A practical operating system emulator written entirely in Go that simulates a complete OS environment including virtual CPU, memory management, process scheduling, and filesystem.

## Learning Objectives

- Build a complete OS emulator from scratch in Go
- Understand CPU instruction execution and emulation
- Implement virtual memory with paging
- Create a realistic process scheduler
- Build a functional filesystem
- Learn how operating systems work at a fundamental level
- Practice advanced Go patterns (channels, goroutines, interfaces)

## Architecture

### Components

1. **Virtual CPU**: Simulates x86-like instruction execution
2. **Memory Manager**: Virtual and physical memory management
3. **Process Manager**: Process lifecycle and management
4. **Scheduler**: Multi-level feedback queue scheduling
5. **Filesystem**: Persistent in-memory filesystem
6. **Shell**: Interactive command interpreter
7. **System Call Handler**: Interface between user and kernel space

## Features

- Virtual CPU with instruction set
- Register simulation (EAX, EBX, ECX, EDX, ESP, EBP, ESI, EDI)
- Memory paging and virtual address translation
- Process creation, execution, and termination
- Multi-level feedback queue scheduling
- Filesystem with files and directories
- Shell with built-in commands
- System call interface
- Interrupt handling
- Context switching simulation
- CPU time accounting
- Memory protection

## Building and Running

### Local Build
```bash
go build -o os-emulator ./cmd/emulator
./os-emulator
```

### Docker Build and Run
```bash
docker build -t os-emulator .
docker run -it os-emulator
```

### Using Makefile
```bash
make build      # Build the emulator
make run        # Run locally
make test       # Run tests
make docker     # Build Docker image
make docker-run # Run in Docker container
make clean      # Clean build artifacts
```

## Shell Commands

### Process Management
- `ps` - List all processes
- `run <program>` - Create and run a process
- `kill <pid>` - Terminate process
- `cpu` - Show CPU statistics

### Memory Management
- `memory` - Show memory statistics
- `heap` - Show heap information
- `stack` - Show stack information

### Filesystem
- `ls [path]` - List directory
- `mkdir <path>` - Create directory
- `touch <path>` - Create file
- `rm <path>` - Delete file
- `cat <path>` - Display file contents
- `write <path> <content>` - Write to file

### CPU & Execution
- `cpu-info` - Display CPU information
- `registers` - Show CPU registers
- `execute <instruction>` - Execute CPU instruction
- `trace` - Enable instruction tracing

### System
- `info` - Display system information
- `time` - Show system time
- `exit` - Shutdown emulator

## Implementation Details

### Virtual CPU Instruction Set
- MOV - Move data between registers/memory
- ADD - Add two values
- SUB - Subtract two values
- JMP - Jump to address
- CALL - Call function
- RET - Return from function
- PUSH - Push to stack
- POP - Pop from stack
- CMP - Compare values
- JE/JNE - Conditional jumps

### Memory Layout
```
0x00000000 - 0x00001000: Kernel space
0x00001000 - 0x00100000: Heap
0x00100000 - 0x7FFFFFFF: User space
0x80000000 - 0xFFFFFFFF: I/O and special
```

### Process States
- New: Just created
- Ready: Waiting to run
- Running: Currently executing
- Blocked: Waiting for I/O
- Terminated: Finished execution

### Scheduling Algorithm
- Multi-level feedback queue
- 4 priority levels
- Time slice: 100ms per level
- Aging to prevent starvation

## Concepts Demonstrated

- CPU emulation and instruction execution
- Virtual memory and paging
- Process scheduling algorithms
- Context switching
- Interrupt handling
- System call interface
- Memory protection
- Process synchronization
- Filesystem design
- Shell implementation
- Go concurrency patterns
- Interface design

## Testing

```bash
go test -v ./...
```

## Resources

- All previous day materials (36-53)
- Go standard library documentation
- Operating systems textbooks
- CPU architecture documentation

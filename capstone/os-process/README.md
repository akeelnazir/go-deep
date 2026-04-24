# OS Process: Process Management Operating System

A moderately complex operating system featuring process scheduling, memory management, and a basic filesystem.

## Learning Objectives

- Understand process scheduling algorithms (round-robin)
- Implement memory management with heap allocation
- Build a basic in-memory filesystem
- Learn context switching and process states
- Understand process lifecycle management
- Implement inter-process communication basics

## Architecture

### Components

1. **Process Manager**: Create, schedule, and manage processes
2. **Scheduler**: Round-robin process scheduling
3. **Memory Manager**: Heap allocation and virtual memory simulation
4. **Filesystem**: In-memory filesystem with directories and files
5. **Shell**: Command interpreter with process execution

## Features

- Process creation and termination
- Round-robin scheduling with time slices
- Process states (Ready, Running, Blocked, Terminated)
- Heap memory allocation and deallocation
- In-memory filesystem with directory structure
- File operations (create, read, write, delete)
- Process priority levels
- CPU time tracking per process
- Memory usage tracking

## Building and Running

### Local Build
```bash
go build -o os-process ./cmd/kernel
./os-process
```

### Docker Build and Run
```bash
docker build -t os-process .
docker run -it os-process
```

### Using Makefile
```bash
make build      # Build the kernel
make run        # Run locally
make test       # Run tests
make docker     # Build Docker image
make docker-run # Run in Docker container
make clean      # Clean build artifacts
```

## Shell Commands

- `help` - Display available commands
- `ps` - List all processes
- `run <name> <duration>` - Create and run a process
- `kill <pid>` - Terminate a process
- `info` - Display system information
- `memory` - Show memory statistics
- `ls [path]` - List directory contents
- `mkdir <path>` - Create directory
- `touch <path>` - Create file
- `rm <path>` - Delete file or directory
- `cat <path>` - Display file contents
- `echo <text> > <path>` - Write to file
- `exit` - Shutdown the OS

## Implementation Details

### Process States
```
Ready -> Running -> Blocked -> Ready
         |
         v
      Terminated
```

### Scheduling Algorithm
- Round-robin with 100ms time slices
- Processes cycle through ready queue
- Blocked processes wait for I/O completion

### Memory Layout
```
0x0000 - 0x0FFF: Kernel
0x1000 - 0x7FFF: Filesystem
0x8000 - 0xFFFF: Process heap
```

### Filesystem Structure
```
/
├── home/
│   └── user/
├── tmp/
└── var/
```

## Concepts Demonstrated

- Process management and scheduling
- Context switching simulation
- Memory allocation and deallocation
- Filesystem operations
- Process synchronization basics
- CPU time accounting
- Process priority handling

## Testing

```bash
go test -v ./...
```

## Resources

- Day 36: Unsafe Pointers and Memory Manipulation
- Day 37: System Calls and OS Interaction
- Day 39: Process Management and Signals
- Day 40: File System Operations
- Day 10-12: Concurrency (Goroutines and Channels)

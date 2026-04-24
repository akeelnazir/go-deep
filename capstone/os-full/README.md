# OS Full: Full-Featured Operating System

A comprehensive operating system implementation with advanced process scheduling, virtual memory, filesystem with permissions, device drivers, and inter-process communication.

## Learning Objectives

- Implement advanced process scheduling (priority-based)
- Build virtual memory with paging simulation
- Create a filesystem with permissions and ownership
- Simulate device drivers
- Implement inter-process communication (pipes, signals)
- Understand system call interface design
- Learn advanced memory management techniques

## Architecture

### Components

1. **Process Manager**: Advanced process lifecycle management
2. **Scheduler**: Priority-based scheduling with multiple queues
3. **Memory Manager**: Virtual memory with paging
4. **Filesystem**: Full filesystem with permissions and ownership
5. **Device Manager**: Simulated device drivers
6. **IPC Manager**: Pipes, signals, and message queues
7. **System Call Interface**: User-kernel boundary
8. **Shell**: Advanced command interpreter

## Features

- Priority-based process scheduling
- Virtual memory with page tables
- Filesystem with permissions (rwx)
- File ownership and groups
- Device drivers (disk, console, network)
- Pipes for inter-process communication
- Signal handling
- Message queues
- System call interface
- User and group management
- Process resource limits

## Building and Running

### Local Build
```bash
go build -o os-full ./cmd/kernel
./os-full
```

### Docker Build and Run
```bash
docker build -t os-full .
docker run -it os-full
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

### Process Management
- `ps` - List all processes
- `run <name> <priority> <duration>` - Create process with priority
- `kill <pid>` - Terminate process
- `nice <pid> <priority>` - Change process priority
- `wait <pid>` - Wait for process completion

### Memory Management
- `memory` - Show memory statistics
- `pages` - Show page table information
- `swap` - Show swap statistics

### Filesystem
- `ls [path]` - List directory
- `mkdir <path>` - Create directory
- `touch <path>` - Create file
- `rm <path>` - Delete file
- `cat <path>` - Display file contents
- `chmod <mode> <path>` - Change permissions
- `chown <user> <path>` - Change owner
- `mount <device> <path>` - Mount device

### IPC
- `pipe <name>` - Create named pipe
- `send <pid> <message>` - Send signal
- `queue <name>` - Create message queue

### System
- `info` - Display system information
- `users` - List users
- `devices` - List devices
- `exit` - Shutdown OS

## Implementation Details

### Process Priority Levels
- 0-3: System processes (highest priority)
- 4-7: Interactive processes
- 8-15: Batch processes (lowest priority)

### Virtual Memory
- Page size: 4KB
- Page table entries: 1024
- Virtual address space: 4MB per process

### Filesystem Permissions
- Owner: rwx (7)
- Group: rwx (7)
- Others: rwx (7)
- Total: 9 bits

### System Calls
- `sys_fork()` - Create process
- `sys_exec()` - Execute program
- `sys_exit()` - Terminate process
- `sys_wait()` - Wait for child
- `sys_open()` - Open file
- `sys_read()` - Read from file
- `sys_write()` - Write to file
- `sys_close()` - Close file
- `sys_pipe()` - Create pipe
- `sys_signal()` - Register signal handler

## Concepts Demonstrated

- Advanced process scheduling algorithms
- Virtual memory and paging
- Filesystem design with permissions
- Device driver architecture
- Inter-process communication mechanisms
- System call interface design
- User and group management
- Resource limits and accounting
- Signal handling and delivery
- Memory protection

## Testing

```bash
go test -v ./...
```

## Resources

- Day 36-45: Systems Programming concepts
- Day 46-53: Advanced Go patterns and deployment

# OS Baremetal: Real Bootable Operating System

A real, bootable operating system that runs independently on bare metal hardware (via QEMU emulation). Unlike the Alpine-based OS projects, this OS does not rely on any host operating system—it runs directly on simulated hardware with its own bootloader, kernel, and system components.

## Key Differences from Other OS Projects

| Aspect | OS Minimal/Full | OS Baremetal |
|--------|-----------------|--------------|
| **Base Image** | Alpine Linux | None (bare metal) |
| **Execution** | Userspace program in container | Direct on simulated hardware |
| **Bootloader** | Simulated | Real bootloader sequence |
| **Hardware** | Abstracted | Direct hardware access |
| **Disk Image** | N/A | Bootable disk image |
| **Emulator** | Docker/Alpine | QEMU |

## Learning Objectives

- Understand real bootloader initialization
- Implement bare metal kernel development
- Learn hardware abstraction layers
- Build bootable disk images
- Understand UEFI/BIOS boot sequences
- Implement real memory management
- Learn interrupt handling at hardware level

## Architecture

### Components

1. **Bootloader**: Real boot sequence with hardware detection
2. **Kernel**: Core OS with interrupt handling and memory management
3. **Memory Manager**: Physical and virtual memory allocation
4. **Disk Manager**: Block device I/O operations
5. **Shell**: Command interpreter
6. **System Calls**: Interface between user and kernel

## Features

- Real bootloader with hardware initialization
- UEFI boot mode support
- Physical memory management
- Virtual memory and paging
- Interrupt vector table with exception handling
- Block device abstraction
- Interactive shell with system commands
- Process management basics
- System information display

## Building and Running

### Prerequisites

- Go 1.26 or later
- QEMU (for local execution)
- Docker (optional, for containerized execution)
- Make (for build automation)

### Local Build and Run

```bash
cd capstone/os-baremetal

# Build the kernel binary
make build

# Run locally (interactive shell)
make run

# Create bootable disk image and run in QEMU
make run-qemu
```

### Docker Build and Run

```bash
# Build Docker image with QEMU
make docker

# Run in Docker container
make docker-run
```

### Using Makefile

```bash
make help       # Show all available targets
make build      # Build the kernel binary
make run        # Run the kernel locally
make disk-image # Create bootable disk image
make run-qemu   # Run in QEMU emulator
make test       # Run tests
make docker     # Build Docker image
make docker-run # Run in Docker container
make clean      # Clean build artifacts
```

## Shell Commands

### System Information
- `help` - Display available commands
- `info` - Display system information
- `uname` - Show system name and version
- `uptime` - Show system uptime

### Memory Management
- `memory` - Show memory statistics and usage

### Process Management
- `ps` - List running processes

### Utilities
- `echo <text>` - Print text to console
- `clear` - Clear the screen
- `exit` - Shutdown the OS

## Implementation Details

### Boot Sequence

1. **Stage 1: Hardware Detection**
   - CPU detection and speed measurement
   - Memory detection and mapping
   - Interrupt handler initialization
   - GDT (Global Descriptor Table) setup
   - Protected mode activation

2. **Stage 2: Bootloader Tasks**
   - Bootable device scanning
   - Kernel loading from disk
   - Kernel signature verification
   - Boot parameter setup

3. **Stage 3: Kernel Entry**
   - Jump to kernel entry point
   - Bootloader cleanup

### Memory Layout

```
0x00000000 - 0x000FFFFF: Real mode memory (1 MB)
0x00100000 - 0x7FFFFFFF: Kernel space (2 GB)
0x80000000 - 0xFFFFFFFF: User space (2 GB)
```

### Interrupt Handling

- **INT 0**: Division by zero
- **INT 6**: Invalid opcode
- **INT 14**: Page fault
- **INT 80**: System call

### Memory Management

The OS implements a simple memory allocator with:
- Block allocation and deallocation
- Fragmentation handling
- Memory coalescing
- Statistics tracking

### Disk Management

Block-based disk I/O with:
- 512-byte block size
- Read/write operations
- Block addressing
- Size management

## Concepts Demonstrated

### Hardware Abstraction
- CPU simulation
- Memory management
- Interrupt handling
- Device drivers

### Bootloader Development
- Hardware initialization
- Boot sequence
- Kernel loading
- Protected mode setup

### Kernel Development
- Interrupt handling
- Memory management
- Process basics
- System calls

### System Architecture
- Bootloader → Kernel → Shell hierarchy
- Hardware abstraction layers
- System call interface
- Device driver architecture

## Testing

```bash
go test -v ./...
```

## Running in QEMU

The OS can be run in QEMU with various configurations:

```bash
# Basic QEMU execution
qemu-system-x86_64 -drive format=raw,file=disk/os.img -m 256 -nographic

# With debugging
qemu-system-x86_64 -drive format=raw,file=disk/os.img -m 256 -nographic -s -S

# With network
qemu-system-x86_64 -drive format=raw,file=disk/os.img -m 256 -nographic -net nic -net user
```

## Extending the OS

Possible extensions:
- Filesystem implementation (FAT, ext2, ext4)
- Network stack (TCP/IP)
- Advanced scheduling algorithms
- Virtual memory with demand paging
- Device drivers (keyboard, mouse, graphics)
- Multi-core support
- Security features (SELinux, AppArmor)
- User authentication
- Package management

## Comparison with Other Projects

### vs. OS Minimal
- **OS Minimal**: Runs in Alpine container, simulated hardware
- **OS Baremetal**: Runs on bare metal via QEMU, real bootloader

### vs. OS Process
- **OS Process**: Focus on process scheduling and management
- **OS Baremetal**: Focus on hardware abstraction and bootloader

### vs. OS Full
- **OS Full**: Advanced features (virtual memory, permissions, IPC)
- **OS Baremetal**: Foundation for bare metal development

### vs. OS Emulator
- **OS Emulator**: CPU emulation and instruction execution
- **OS Baremetal**: Real bootable OS with disk images

## Performance Characteristics

- **Build time**: 5-10 seconds
- **Startup time**: 100-200ms
- **Memory usage**: ~50MB (in QEMU)
- **Disk image size**: 10MB (configurable)

## Resources and References

- Day 36: Unsafe Pointers and Memory Manipulation
- Day 37: System Calls and OS Interaction
- Day 39: Process Management and Signals
- Day 40: File System Operations
- Day 41: Memory Profiling and Optimization
- Day 42: Concurrency at Scale
- Day 45: Systems Programming Capstone

## Troubleshooting

### Build Issues

```bash
# Clean and rebuild
make clean
make build

# Check Go version
go version  # Should be 1.26+
```

### QEMU Execution Issues

```bash
# Check if QEMU is installed
qemu-system-x86_64 --version

# Run with verbose output
qemu-system-x86_64 -drive format=raw,file=disk/os.img -m 256 -nographic -d int
```

### Docker Issues

```bash
# Rebuild image without cache
docker build --no-cache -t os-baremetal .

# View logs
docker logs <container-id>
```

## Contributing

To extend or modify this project:
1. Read this README and understand the architecture
2. Study the existing code structure
3. Follow the established patterns
4. Add tests for new features
5. Update documentation

## License

This OS Baremetal project is provided as-is for educational purposes.

---

**Happy bare metal development!** Start by running `make run` to see the OS in action, then explore the bootloader and kernel implementation.

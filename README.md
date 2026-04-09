# Go Deep: A 29-Day Comprehensive Learning Curriculum

A condensed, structured, hands-on learning path for mastering Go from fundamentals to advanced patterns. Each day includes exercises, tests, and detailed explanations. This curriculum consolidates related topics to maintain progressive complexity while reducing redundancy.

## Prerequisites

- **Go 1.26.1** or later ([Install Go](https://golang.org/doc/install))
- A code editor or IDE (VS Code with Go extension, GoLand, or similar)
- Basic command-line familiarity
- No prior Go experience required

## Project Structure

```
go-deep/
├── README.md              # This file
├── exercises/             # Exercises directory
|   ├── day1/                  # Day 1: Go Setup, Syntax, and Basic Types
|   |   ├── README.md
|   |   ├── exercise.go
|   |   ├── exercise_test.go
|   |   └── main.go
|   ├── day2/                  # Day 2: Control Structures and Collections
|   |   ├── README.md
|   |   ├── exercise.go
|   |   ├── exercise_test.go
|   |   └── main.go
|   ├── day3 through day29     # Additional learning modules
└── ...
```

Each day folder contains:
- **README.md**: Learning objectives and topic explanations
- **exercise.go**: Functions to implement (marked with TODO comments)
- **exercise_test.go**: Unit tests to verify your implementations
- **main.go**: Example code and demonstrations

## Getting Started

### 1. Clone or Navigate to the Repository

```bash
cd /Users/<username>/go/src/github.com/<username>/go-deep
```

### 2. Run a Specific Day's Exercises

```bash
# Navigate to a day folder
cd exercises/day1

# Run the main program
go run main.go
# OR
go run .

# Run the tests
go test -v
```

### 3. Run All Tests

```bash
# From the repository root, run all tests
go test ./...
```

### 4. Work Through Exercises

1. Read the `README.md` in each day folder to understand the concepts
2. Look at `main.go` for examples
3. Implement the TODO functions in `exercise.go`
4. Run `go test -v` to verify your implementations
5. Move to the next day when tests pass

## Curriculum Overview

| Day | Topic |
|-----|-------|
| 1 | Go Setup, Syntax, and Basic Types |
| 2 | Control Structures and Collections |
| 3 | Functions and Structs |
| 4 | Pointers and Interfaces |
| 5 | Error Handling and Project Structure |
| 6 | Concurrency Fundamentals |
| 7 | Synchronization Patterns |
| 8 | Context and Advanced Concurrency |
| 9 | Testing and Benchmarking |
| 10 | Reflection and Generics |
| 11 | Web Fundamentals and HTTP |
| 12 | Routing and Middleware |
| 13 | REST API Design |
| 14 | Authentication and Authorization |
| 15 | Databases and ORM |
| 16 | WebSockets and Real-time Communication |
| 17 | Testing Web Applications |
| 18 | File and IO Operations |
| 19 | Serialization and Encoding |
| 20 | Command Line Applications |
| 21 | Signal Handling and Processes |
| 22 | Networking Fundamentals |
| 23 | Parsing and Text Processing |
| 24 | Cryptography and Security |
| 25 | Memory Management and Unsafe |
| 26 | Web Frameworks and MVC Architecture |
| 27 | Embedding and Code Generation |
| 28 | Advanced Go - Memory, Unsafe Pointers, CGO, Data Structures, and Functional Patterns |
| 29 | Final Review and Next Steps |

## Learning Path Breakdown

### Week 1: Fundamentals (Days 1-5)
Master Go basics: syntax, data types, control flow, functions, structs, pointers, interfaces, and error handling. Build the foundation for all subsequent learning.

### Week 2: Core Programming (Days 6-10)
Learn goroutines, channels, synchronization patterns, context management, testing, reflection, and generics—core concurrency and advanced language features.

### Week 3: Web Development (Days 11-17)
Build HTTP servers and clients, implement routing and middleware, design RESTful APIs, add authentication, integrate databases, and test web applications.

### Week 4: Systems Programming (Days 18-24)
Work with files, serialization, command-line applications, process management, networking, text processing, and cryptography for systems-level programming.

### Week 5: Advanced Systems (Days 25-29)
Master memory management, unsafe pointers, web frameworks, code generation, functional patterns, and solidify your Go expertise with a comprehensive review.

## Capstone Projects

The `capstone/` directory contains five comprehensive operating system implementations at increasing complexity levels, demonstrating advanced systems programming concepts from the curriculum:

### 1. **OS Minimal** (`capstone/os-minimal/`)
A minimal bootloader and kernel with interrupt handling and a simple shell. Perfect for understanding OS fundamentals.
- Bootloader initialization
- Interrupt vector table
- Basic memory management
- Simple command shell
- ~500-1000 lines

### 2. **OS Process** (`capstone/os-process/`)
A process management OS with round-robin scheduling, memory allocation, and an in-memory filesystem.
- Process scheduling (round-robin)
- Memory management with heap allocation
- In-memory filesystem with directories
- Process creation and termination
- ~2000-3000 lines

### 3. **OS Full** (`capstone/os-full/`)
A full-featured OS with priority-based scheduling, virtual memory paging, filesystem permissions, device drivers, and IPC.
- Priority-based process scheduling
- Virtual memory with paging
- Filesystem with permissions and ownership
- Device driver simulation
- Inter-process communication (pipes, signals)
- ~5000+ lines

### 4. **OS Emulator** (`capstone/os-emulator/`)
A practical Go-based OS emulator with virtual CPU, memory management, and process scheduling.
- Virtual CPU emulation with registers
- Memory paging and translation
- Process scheduling
- Filesystem simulation
- Shell interface
- ~3000-4000 lines

### 5. **OS Baremetal** (`capstone/os-baremetal/`)
A real bootable operating system that runs independently on bare metal hardware via QEMU emulation. No host OS dependency.
- Real bootloader with hardware detection
- UEFI boot mode support
- Physical memory management
- Bootable disk image creation
- Block device I/O abstraction
- Direct hardware access simulation
- ~2000-3000 lines

### Building and Running Capstone Projects

Each project includes a Dockerfile and Makefile:

```bash
cd capstone/os-minimal
make build      # Build locally
make run        # Run the OS
make docker     # Build Docker image
make docker-run # Run in Docker container
```

## Common Commands

```bash
# Format code
go fmt ./...

# Run static analysis
go vet ./...

# Run tests with coverage
go test -cover ./...

# Run a specific test
go test -run TestName -v

# Build a binary
go build -o myapp ./day1
```

## Tips for Success

1. **Read the READMEs**: Each day's README contains essential context and explanations
2. **Study the examples**: `main.go` files demonstrate the concepts in action
3. **Implement incrementally**: Don't try to complete all TODOs at once
4. **Run tests frequently**: Tests provide immediate feedback on your implementation
5. **Experiment**: Modify examples and try variations to deepen understanding
6. **Take notes**: Jot down key concepts as you progress
7. **Progress at your pace**: The 29-day structure is a guideline; adjust timing based on your learning speed

## Resources

- [Official Go Documentation](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go by Example](https://gobyexample.com/)

## Module Information

- **Module**: `github.com/<username>/go-deep`
- **Go Version**: 1.26.1

## License

This learning curriculum is provided as-is for educational purposes.

---

Happy learning! Start with Day 1 and progress through the curriculum at your own pace.

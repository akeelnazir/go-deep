# Day 1: Go Setup, Syntax, and Basic Types

## Learning Objectives

- Install Go and set up development environment
- Write and run "Hello, World!" programs
- Declare variables using `var`, short declaration (`:=`), and constants
- Master Go's complete type system: integers, floats, strings, booleans, and complex numbers
- Understand zero values and their practical implications
- Perform explicit type conversions safely and idiomatically
- Write idiomatic Go code following Go conventions

## Topics Covered

### 1. Go Installation and Setup
Go is a compiled, statically-typed language developed by Google. Installation is straightforward across all platforms (Windows, macOS, Linux). After installation, verify with `go version`.

### 2. GOPATH and Go Modules
- **GOPATH**: The traditional workspace directory (deprecated in favor of modules)
- **Go Modules**: Modern dependency management using `go.mod` and `go.sum` files
- Initialize a module with `go mod init github.com/username/projectname`

### 3. Development Tools
- **VS Code**: Install the official Go extension for syntax highlighting, debugging, and IntelliSense
- **GoLand**: JetBrains IDE with excellent Go support
- **Vim/Neovim**: Lightweight option with plugins like vim-go

### 4. Running Go Programs
- **`go run`**: Compiles and executes in one step (useful for development)
- **`go build`**: Compiles to a binary executable
- **`go install`**: Builds and installs the binary to `$GOPATH/bin`

### 5. Project Structure
A typical Go project follows this structure:
```
myproject/
├── go.mod
├── go.sum
├── main.go
├── cmd/
│   └── myapp/
│       └── main.go
├── internal/
│   └── mypackage/
│       └── myfile.go
└── pkg/
    └── publicpackage/
        └── publicfile.go
```

### 6. Variables and Constants
Go provides three primary ways to declare variables:
- **`var` keyword**: Explicit, traditional way for package and function scope
- **Short declaration (`:=`)**: Concise way for local variables inside functions
- **`const` keyword**: Immutable values assigned at compile time

See `main.go` lines 19-33 for examples of variable declaration methods, including `var` with explicit type, type inference, short declaration, and multiple variable declarations.

### 7. Basic Data Types
Go's type system includes:
- **Integers**: int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint, byte, rune
- **Floating-point**: float32, float64
- **Strings**: Immutable sequences of bytes
- **Booleans**: true, false
- **Complex numbers**: complex64, complex128

See `main.go` lines 49-110 for examples of all basic data types in action.

### 8. Type Conversions
Go requires explicit type conversions. There is no implicit type coercion between types.

See `main.go` lines 127-151 for examples of type conversions including int to float64, float64 to int, int to rune, and string conversions using the `strconv` package.

### 9. Zero Values
Every type in Go has a zero value - the default value when a variable is declared but not initialized.

See `main.go` lines 113-123 for demonstrations of zero values for different types (int, float64, string, bool).

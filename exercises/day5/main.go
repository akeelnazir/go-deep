package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Section 1: The Error Interface
func demonstrateErrorInterface() {
	fmt.Println("=== The Error Interface ===")

	var err error
	fmt.Printf("Nil error: %v\n", err == nil)

	err = errors.New("something went wrong")
	fmt.Printf("Non-nil error: %v\n", err == nil)
	fmt.Printf("Error message: %v\n", err)
	fmt.Println()
}

// Section 2: Multiple Return Values for Errors
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func demonstrateMultipleReturns() {
	fmt.Println("=== Multiple Return Values for Errors ===")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println()
}

// Section 3: Creating Custom Error Types
func validateAge(age int) error {
	if age < 0 {
		return errors.New("age cannot be negative")
	}
	return nil
}

func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("invalid age: %w", err)
	}
	return age, nil
}

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation error in %s: %s", e.Field, e.Message)
}

func validateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return ValidationError{
			Field:   "email",
			Message: "must contain @",
		}
	}
	return nil
}

func demonstrateCustomErrors() {
	fmt.Println("=== Creating Custom Error Types ===")

	err := validateAge(-5)
	if err != nil {
		fmt.Println("Age validation error:", err)
	}

	age, err := parseAge("25")
	if err != nil {
		fmt.Println("Parse error:", err)
	} else {
		fmt.Printf("Parsed age: %d\n", age)
	}

	err = validateEmail("invalid-email")
	if err != nil {
		if ve, ok := err.(ValidationError); ok {
			fmt.Printf("Field: %s, Message: %s\n", ve.Field, ve.Message)
		}
	}
	fmt.Println()
}

// Section 4: Error Wrapping with Context
func readFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}
	return data, nil
}

func demonstrateErrorWrapping() {
	fmt.Println("=== Error Wrapping with Context ===")

	data, err := readFile("nonexistent.txt")
	if err != nil {
		fmt.Println("Error:", err)

		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("File does not exist")
		}
	} else {
		fmt.Printf("Read %d bytes\n", len(data))
	}
	fmt.Println()
}

// Section 5: Error Handling Patterns - Guard Clause
var (
	ErrNotFound     = errors.New("not found")
	ErrInvalidInput = errors.New("invalid input")
)

type User struct {
	ID    string
	Name  string
	Email string
}

var users = map[string]*User{
	"1": {ID: "1", Name: "Alice", Email: "alice@example.com"},
	"2": {ID: "2", Name: "Bob", Email: "bob@example.com"},
}

func getUser(id string) (*User, error) {
	if id == "" {
		return nil, ErrInvalidInput
	}

	user, exists := users[id]
	if !exists {
		return nil, ErrNotFound
	}

	return user, nil
}

func demonstrateGuardClauses() {
	fmt.Println("=== Guard Clause Pattern ===")

	user, err := getUser("1")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found user: %s (%s)\n", user.Name, user.Email)
	}

	user, err = getUser("")
	if err == ErrInvalidInput {
		fmt.Println("Please provide a valid ID")
	}

	user, err = getUser("999")
	if err == ErrNotFound {
		fmt.Println("User not found")
	}
	fmt.Println()
}

// Section 6: Panic and Recover
func mustParseInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int: %v", err))
	}
	return n
}

func safeDivide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
			result = 0
		}
	}()

	if b == 0 {
		panic("division by zero")
	}
	return a / b
}

func demonstratePanicRecover() {
	fmt.Println("=== Panic and Recover ===")

	result := safeDivide(10, 0)
	fmt.Printf("Result after recovery: %d\n", result)

	result = safeDivide(20, 4)
	fmt.Printf("Normal division: %d\n", result)
	fmt.Println()
}

// Section 7: Project Structure Example
type Config struct {
	AppName string
	Port    int
	Debug   bool
}

func loadConfig(filename string) (*Config, error) {
	if filename == "" {
		return nil, errors.New("filename cannot be empty")
	}

	return &Config{
		AppName: "MyApp",
		Port:    8080,
		Debug:   true,
	}, nil
}

func demonstrateProjectStructure() {
	fmt.Println("=== Project Structure Example ===")

	config, err := loadConfig("config.json")
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}

	fmt.Printf("Config: %s on port %d (Debug: %v)\n", config.AppName, config.Port, config.Debug)
	fmt.Println()
}

// Section 8: Defer for Cleanup
func processFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	_, err = file.WriteString("Hello, World!")
	if err != nil {
		return fmt.Errorf("failed to write to file: %w", err)
	}

	return nil
}

func demonstrateDeferCleanup() {
	fmt.Println("=== Defer for Cleanup ===")

	filename := "/tmp/test_day5.txt"
	err := processFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File processed successfully")
		os.Remove(filename)
	}
	fmt.Println()
}

// Section 9: Common Mistakes
func demonstrateCommonMistakes() {
	fmt.Println("=== Common Mistakes and Best Practices ===")

	fmt.Println("1. Always check errors immediately:")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Result: %d\n", result)
	}

	fmt.Println("\n2. Wrap errors with context:")
	_, err = parseAge("invalid")
	if err != nil {
		fmt.Println("Error with context:", err)
	}

	fmt.Println("\n3. Use sentinel errors for comparison:")
	user, err := getUser("999")
	if err == ErrNotFound {
		fmt.Println("User not found - use sentinel error for comparison")
	}
	_ = user

	fmt.Println()
}

func main() {
	fmt.Println("=== Day 5: Error Handling, Logging and Project Structure ===")

	demonstrateErrorInterface()
	demonstrateMultipleReturns()
	demonstrateCustomErrors()
	demonstrateErrorWrapping()
	demonstrateGuardClauses()
	demonstratePanicRecover()
	demonstrateProjectStructure()
	demonstrateDeferCleanup()
	demonstrateCommonMistakes()

	fmt.Println("=== Day 5 Complete ===")
	fmt.Println("Next: Learn about concurrency fundamentals on Day 6.")
}

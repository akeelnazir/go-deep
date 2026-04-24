package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Command struct {
	Name        string
	Description string
	Handler     func(args []string) error
}

var commands = make(map[string]*Command)

func registerCommand(name, desc string, handler func(args []string) error) {
	commands[name] = &Command{
		Name:        name,
		Description: desc,
		Handler:     handler,
	}
}

func executeCommand(name string, args []string) error {
	cmd, ok := commands[name]
	if !ok {
		return fmt.Errorf("unknown command: %s", name)
	}
	return cmd.Handler(args)
}

func init() {
	registerCommand("greet", "Greet someone", func(args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("name required")
		}
		fmt.Printf("Hello, %s!\n", args[0])
		return nil
	})

	registerCommand("add", "Add two numbers", func(args []string) error {
		if len(args) < 2 {
			return fmt.Errorf("two numbers required")
		}
		a, _ := strconv.Atoi(args[0])
		b, _ := strconv.Atoi(args[1])
		fmt.Printf("%d + %d = %d\n", a, b, a+b)
		return nil
	})

	registerCommand("echo", "Echo arguments", func(args []string) error {
		fmt.Println(strings.Join(args, " "))
		return nil
	})

	registerCommand("help", "Show help", func(args []string) error {
		fmt.Println("Available commands:")
		for name, cmd := range commands {
			fmt.Printf("  %s - %s\n", name, cmd.Description)
		}
		return nil
	})
}

func main() {
	fmt.Println("=== Day 20: Command Line Applications ===")

	fmt.Println("\n--- CLI Command Execution ---")

	executeCommand("greet", []string{"Alice"})
	executeCommand("add", []string{"5", "3"})
	executeCommand("echo", []string{"Hello", "World"})

	fmt.Println("\n--- Help Command ---")
	executeCommand("help", []string{})

	fmt.Println("\n--- Error Handling ---")
	err := executeCommand("greet", []string{})
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println("\n=== Day 20 Complete ===")
	fmt.Println("Next: Learn about signal handling on Day 21.")
}

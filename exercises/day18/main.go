package main

import (
	"fmt"
	"strings"
)

var fileSystem = make(map[string]string)

func writeFile(filename, content string) error {
	fileSystem[filename] = content
	return nil
}

func readFile(filename string) (string, error) {
	content, ok := fileSystem[filename]
	if !ok {
		return "", fmt.Errorf("file not found: %s", filename)
	}
	return content, nil
}

func fileExists(filename string) bool {
	_, ok := fileSystem[filename]
	return ok
}

func deleteFile(filename string) error {
	if !fileExists(filename) {
		return fmt.Errorf("file not found: %s", filename)
	}
	delete(fileSystem, filename)
	return nil
}

func getFileSize(filename string) int {
	content, ok := fileSystem[filename]
	if !ok {
		return -1
	}
	return len(content)
}

func appendToFile(filename, content string) error {
	existing, ok := fileSystem[filename]
	if !ok {
		return fmt.Errorf("file not found: %s", filename)
	}
	fileSystem[filename] = existing + content
	return nil
}

func readLines(filename string) ([]string, error) {
	content, err := readFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(content, "\n"), nil
}

func main() {
	fmt.Println("=== Day 18: File and IO Operations ===")

	fmt.Println("\n--- Writing Files ---")
	writeFile("test.txt", "Hello, World!")
	writeFile("data.txt", "Line 1\nLine 2\nLine 3")
	fmt.Println("Files written")

	fmt.Println("\n--- Reading Files ---")
	content, _ := readFile("test.txt")
	fmt.Printf("Content of test.txt: %s\n", content)

	fmt.Println("\n--- File Information ---")
	if fileExists("test.txt") {
		size := getFileSize("test.txt")
		fmt.Printf("File size: %d bytes\n", size)
	}

	fmt.Println("\n--- Appending to Files ---")
	appendToFile("test.txt", "\nAppended line")
	content, _ = readFile("test.txt")
	fmt.Printf("Updated content: %s\n", content)

	fmt.Println("\n--- Reading Lines ---")
	lines, _ := readLines("data.txt")
	for i, line := range lines {
		if line != "" {
			fmt.Printf("Line %d: %s\n", i+1, line)
		}
	}

	fmt.Println("\n--- Deleting Files ---")
	deleteFile("test.txt")
	if !fileExists("test.txt") {
		fmt.Println("File deleted successfully")
	}

	fmt.Println("\n=== Day 18 Complete ===")
	fmt.Println("Next: Learn about serialization on Day 19.")
}

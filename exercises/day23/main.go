package main

import (
	"fmt"
	"regexp"
	"strings"
)

func matchPattern(pattern, text string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(text, -1)
}

func replacePattern(pattern, text, replacement string) string {
	re := regexp.MustCompile(pattern)
	return re.ReplaceAllString(text, replacement)
}

func splitText(separator, text string) []string {
	return strings.Split(text, separator)
}

func joinText(sep string, parts []string) string {
	return strings.Join(parts, sep)
}

func trimText(text string) string {
	return strings.TrimSpace(text)
}

func containsPattern(pattern, text string) bool {
	re := regexp.MustCompile(pattern)
	return re.MatchString(text)
}

func main() {
	fmt.Println("=== Day 23: Parsing and Text Processing ===")

	fmt.Println("\n--- Pattern Matching ---")
	text := "The numbers are 123 and 456"
	matches := matchPattern(`\d+`, text)
	fmt.Printf("Numbers found: %v\n", matches)

	fmt.Println("\n--- Pattern Replacement ---")
	result := replacePattern(`\d+`, text, "X")
	fmt.Printf("After replacement: %s\n", result)

	fmt.Println("\n--- String Splitting ---")
	csv := "apple,banana,cherry"
	parts := splitText(",", csv)
	fmt.Printf("Parts: %v\n", parts)

	fmt.Println("\n--- String Joining ---")
	joined := joinText("-", []string{"one", "two", "three"})
	fmt.Printf("Joined: %s\n", joined)

	fmt.Println("\n--- Text Trimming ---")
	padded := "  hello world  "
	trimmed := trimText(padded)
	fmt.Printf("Trimmed: '%s'\n", trimmed)

	fmt.Println("\n--- Pattern Matching Check ---")
	email := "user@example.com"
	if containsPattern(`\w+@\w+\.\w+`, email) {
		fmt.Println("Valid email format")
	}

	fmt.Println("\n=== Day 23 Complete ===")
	fmt.Println("Next: Learn about cryptography on Day 24.")
}

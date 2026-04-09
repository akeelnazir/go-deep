package main

import (
	"fmt"
)

func main() {
	fmt.Println("=== Day 2: Control Structures ===")

	// Section 1: If/Else Statements
	fmt.Println("\n--- If/Else Statements ---")

	age := 25
	if age >= 18 {
		fmt.Println("You are an adult")
	} else {
		fmt.Println("You are a minor")
	}

	score := 85
	if score >= 90 {
		fmt.Println("Grade: A")
	} else if score >= 80 {
		fmt.Println("Grade: B")
	} else if score >= 70 {
		fmt.Println("Grade: C")
	} else {
		fmt.Println("Grade: F")
	}

	// Variable declaration in if condition
	if temperature := 28; temperature > 25 {
		fmt.Printf("It's hot! Temperature: %d°C\n", temperature)
	} else {
		fmt.Printf("It's cool. Temperature: %d°C\n", temperature)
	}

	// Logical operators in conditions
	x := 15
	if x > 10 && x < 20 {
		fmt.Printf("%d is between 10 and 20\n", x)
	}

	if x < 5 || x > 20 {
		fmt.Println("x is outside the range [5, 20]")
	} else {
		fmt.Println("x is within the range [5, 20]")
	}

	fmt.Println()

	// Section 2: Switch Statements
	fmt.Println("--- Switch Statements ---")

	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	// Switch with string
	fruit := "apple"
	switch fruit {
	case "apple":
		fmt.Println("Red or green fruit")
	case "banana":
		fmt.Println("Yellow fruit")
	case "orange":
		fmt.Println("Orange fruit")
	default:
		fmt.Println("Unknown fruit")
	}

	// Switch with multiple cases
	letter := 'A'
	switch letter {
	case 'A', 'E', 'I', 'O', 'U':
		fmt.Printf("%c is a vowel\n", letter)
	case 'B', 'C', 'D', 'F', 'G':
		fmt.Printf("%c is a consonant\n", letter)
	default:
		fmt.Printf("%c is not a letter\n", letter)
	}

	fmt.Println()

	// Section 3: Expressionless Switch (switch true)
	fmt.Println("--- Expressionless Switch ---")

	num := 42
	switch {
	case num < 0:
		fmt.Println("Negative number")
	case num == 0:
		fmt.Println("Zero")
	case num > 0 && num < 10:
		fmt.Println("Single digit positive number")
	case num >= 10 && num < 100:
		fmt.Println("Two digit positive number")
	default:
		fmt.Println("Large number")
	}

	// Another expressionless switch example
	temperature := 35
	switch {
	case temperature < 0:
		fmt.Println("Freezing")
	case temperature < 15:
		fmt.Println("Cold")
	case temperature < 25:
		fmt.Println("Mild")
	case temperature < 35:
		fmt.Println("Warm")
	default:
		fmt.Println("Hot")
	}

	fmt.Println()

	// Section 4: Switch with Fallthrough
	fmt.Println("--- Switch with Fallthrough ---")

	value := 2
	switch value {
	case 1:
		fmt.Println("One")
		fallthrough
	case 2:
		fmt.Println("Two")
		fallthrough
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Other")
	}

	fmt.Println()

	// Section 5: Basic For Loop
	fmt.Println("--- Basic For Loop ---")

	for i := 0; i < 5; i++ {
		fmt.Printf("i = %d\n", i)
	}

	// For loop with multiple variables
	for i, j := 0, 10; i < 5; i, j = i+1, j-1 {
		fmt.Printf("i = %d, j = %d\n", i, j)
	}

	fmt.Println()

	// Section 6: While-Style For Loop
	fmt.Println("--- While-Style For Loop ---")

	counter := 0
	for counter < 5 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	fmt.Println()

	// Section 7: Infinite For Loop
	fmt.Println("--- Infinite For Loop (with break) ---")

	count := 0
	for {
		if count >= 3 {
			break
		}
		fmt.Printf("Iteration: %d\n", count)
		count++
	}

	fmt.Println()

	// Section 8: Range Loop with Slices
	fmt.Println("--- Range Loop with Slices ---")

	numbers := []int{10, 20, 30, 40, 50}
	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}

	// Range with only values
	for _, num := range numbers {
		fmt.Printf("Value: %d\n", num)
	}

	// Range with only indices
	for i := range numbers {
		fmt.Printf("Index: %d\n", i)
	}

	fmt.Println()

	// Section 9: Range Loop with Strings
	fmt.Println("--- Range Loop with Strings ---")

	str := "Hello"
	for i, ch := range str {
		fmt.Printf("Index: %d, Character: %c (Unicode: %d)\n", i, ch, ch)
	}

	fmt.Println()

	// Section 10: Range Loop with Maps
	fmt.Println("--- Range Loop with Maps ---")

	person := map[string]string{
		"name": "Alice",
		"city": "New York",
		"job":  "Engineer",
	}

	for key, value := range person {
		fmt.Printf("%s: %s\n", key, value)
	}

	fmt.Println()

	// Section 11: Break and Continue
	fmt.Println("--- Break and Continue ---")

	fmt.Println("Using continue:")
	for i := 0; i < 5; i++ {
		if i == 2 {
			continue
		}
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println("Using break:")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Printf("i = %d\n", i)
	}

	fmt.Println()

	// Section 12: Labeled Loops
	fmt.Println("--- Labeled Loops ---")

OuterLoop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("Breaking out of outer loop")
				break OuterLoop
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
	}

	fmt.Println()

	// Section 13: Practical Example - Nested Loops
	fmt.Println("--- Practical Example: Multiplication Table ---")

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			fmt.Printf("%d*%d=%d  ", i, j, i*j)
		}
		fmt.Println()
	}

	fmt.Println()

	// Section 14: Practical Example - Finding Prime Numbers
	fmt.Println("--- Practical Example: Prime Numbers ---")

	fmt.Print("Prime numbers from 2 to 20: ")
	for num := 2; num <= 20; num++ {
		isPrime := true
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			fmt.Printf("%d ", num)
		}
	}
	fmt.Println()

	fmt.Println()

	// Section 15: Arrays
	fmt.Println("--- Arrays ---")

	// Array declaration and initialization
	var arr [5]int
	fmt.Printf("Zero-valued array: %v\n", arr)

	arr2 := [3]string{"apple", "banana", "cherry"}
	fmt.Printf("String array: %v\n", arr2)

	arr3 := [...]int{10, 20, 30, 40, 50}
	fmt.Printf("Inferred length array: %v, length: %d\n", arr3, len(arr3))

	// Array operations
	arr4 := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("Element at index 2: %d\n", arr4[2])
	arr4[2] = 99
	fmt.Printf("After modification: %v\n", arr4)

	// Array iteration
	fmt.Print("Array iteration: ")
	for i, val := range arr4 {
		fmt.Printf("[%d]=%d ", i, val)
	}
	fmt.Println()

	fmt.Println()

	// Section 16: Slices
	fmt.Println("--- Slices ---")

	// Slice declaration
	slice1 := []int{10, 20, 30, 40, 50}
	fmt.Printf("Slice literal: %v, len=%d, cap=%d\n", slice1, len(slice1), cap(slice1))

	// Slice from array
	arr5 := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := arr5[2:5]
	fmt.Printf("Slice from array [2:5]: %v, len=%d, cap=%d\n", slice2, len(slice2), cap(slice2))

	// Slice variations
	slice3 := arr5[:]
	fmt.Printf("Full slice [:]: %v\n", slice3)

	slice4 := arr5[3:]
	fmt.Printf("Slice from 3 onward [3:]: %v\n", slice4)

	slice5 := arr5[:7]
	fmt.Printf("Slice up to 7 [:7]: %v\n", slice5)

	// Make slices
	slice6 := make([]int, 5)
	fmt.Printf("make([]int, 5): %v, len=%d, cap=%d\n", slice6, len(slice6), cap(slice6))

	slice7 := make([]int, 3, 10)
	fmt.Printf("make([]int, 3, 10): %v, len=%d, cap=%d\n", slice7, len(slice7), cap(slice7))

	// Append operations
	slice8 := []int{1, 2, 3}
	fmt.Printf("Before append: %v, len=%d, cap=%d\n", slice8, len(slice8), cap(slice8))

	slice8 = append(slice8, 4)
	fmt.Printf("After append(4): %v, len=%d, cap=%d\n", slice8, len(slice8), cap(slice8))

	slice8 = append(slice8, 5, 6, 7)
	fmt.Printf("After append(5,6,7): %v, len=%d, cap=%d\n", slice8, len(slice8), cap(slice8))

	// Copy operation
	source := []int{1, 2, 3, 4, 5}
	destination := make([]int, len(source))
	copy(destination, source)
	destination[0] = 999
	fmt.Printf("Source: %v, Destination: %v (independent)\n", source, destination)

	// Slice sharing underlying array
	arr6 := [5]int{1, 2, 3, 4, 5}
	s1 := arr6[1:3]
	s2 := arr6[2:4]
	fmt.Printf("Before modification: s1=%v, s2=%v\n", s1, s2)
	s1[1] = 99
	fmt.Printf("After s1[1]=99: arr6=%v, s1=%v, s2=%v\n", arr6, s1, s2)

	fmt.Println()

	// Section 17: Maps
	fmt.Println("--- Maps ---")

	// Map literal
	personMap := map[string]string{
		"name": "Alice",
		"city": "New York",
		"job":  "Engineer",
	}
	fmt.Printf("Map literal: %v\n", personMap)

	// Empty map
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	scores["Carol"] = 92
	fmt.Printf("Scores map: %v\n", scores)

	// Map access
	fmt.Printf("Alice's score: %d\n", scores["Alice"])
	fmt.Printf("Unknown score (zero value): %d\n", scores["Unknown"])

	// Safe map access with comma-ok
	if value, exists := scores["Bob"]; exists {
		fmt.Printf("Bob's score exists: %d\n", value)
	} else {
		fmt.Println("Bob not found")
	}

	// Map deletion
	delete(scores, "Bob")
	fmt.Printf("After deleting Bob: %v\n", scores)

	// Map iteration (order is random)
	fmt.Print("Map iteration (random order): ")
	for name, score := range scores {
		fmt.Printf("%s=%d ", name, score)
	}
	fmt.Println()

	// Map of slices
	groups := make(map[string][]string)
	groups["team1"] = append(groups["team1"], "Alice")
	groups["team1"] = append(groups["team1"], "Bob")
	groups["team2"] = append(groups["team2"], "Carol")
	fmt.Printf("Map of slices: %v\n", groups)

	// Counting with maps
	words := []string{"apple", "banana", "apple", "cherry", "banana", "apple"}
	frequency := make(map[string]int)
	for _, word := range words {
		frequency[word]++
	}
	fmt.Printf("Word frequency: %v\n", frequency)

	fmt.Println("\n=== Day 2 Complete ===")
	fmt.Println("Next: Learn about functions on Day 3.")
}

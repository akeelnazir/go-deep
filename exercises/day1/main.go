package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("=== Day 1: Go Setup, Syntax, and Basic Types ===")

	// Part 1: Hello World
	fmt.Println("--- Part 1: Hello, World! ---")
	fmt.Println("Hello, World!")
	fmt.Println()

	// Part 2: Variable Declaration Methods
	fmt.Println("--- Part 2: Variable Declaration Methods ---")

	// Using var keyword with explicit type
	var name string = "Alice"
	fmt.Printf("Using var with explicit type: name = %q\n", name)

	// Using var with type inference
	var age = 25
	fmt.Printf("Using var with type inference: age = %d (type: %T)\n", age, age)

	// Using short declaration (preferred for local variables)
	city := "New York"
	fmt.Printf("Using := (short declaration): city = %q\n", city)

	// Multiple variable declarations
	a, b, c := 1, 2, 3
	fmt.Printf("Multiple declarations: a=%d, b=%d, c=%d\n", a, b, c)
	fmt.Println()

	// Part 3: Constants
	fmt.Println("--- Part 3: Constants ---")

	const Pi = 3.14159
	const MaxRetries int = 3
	const Greeting = "Hello, Go!"

	fmt.Printf("Untyped constant Pi: %v\n", Pi)
	fmt.Printf("Typed constant MaxRetries: %d\n", MaxRetries)
	fmt.Printf("String constant Greeting: %q\n", Greeting)
	fmt.Println()

	// Part 4: Integer Types
	fmt.Println("--- Part 4: Integer Types ---")

	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807
	var defaultInt int = 42

	fmt.Printf("int8: %d (range: -128 to 127)\n", i8)
	fmt.Printf("int16: %d (range: -32,768 to 32,767)\n", i16)
	fmt.Printf("int32: %d (range: -2,147,483,648 to 2,147,483,647)\n", i32)
	fmt.Printf("int64: %d (range: -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807)\n", i64)
	fmt.Printf("int (platform-dependent): %d\n", defaultInt)
	fmt.Println()

	// Part 5: Floating-Point Types
	fmt.Println("--- Part 5: Floating-Point Types ---")

	var f32 float32 = 3.14
	var f64 float64 = 3.14159265359

	fmt.Printf("float32: %f\n", f32)
	fmt.Printf("float64: %.10f\n", f64)
	fmt.Println()

	// Part 6: String Type
	fmt.Println("--- Part 6: String Type ---")

	s1 := "Hello, World!"
	s2 := `Raw string literal
with multiple lines`

	fmt.Printf("Interpreted string: %q\n", s1)
	fmt.Printf("Raw string: %q\n", s2)
	fmt.Printf("String length: %d\n", len(s1))
	fmt.Printf("Substring: %q\n", s1[0:5])
	fmt.Println()

	// Part 7: Boolean Type
	fmt.Println("--- Part 7: Boolean Type ---")

	active := true
	inactive := false

	fmt.Printf("active: %v\n", active)
	fmt.Printf("inactive: %v\n", inactive)
	fmt.Printf("!active: %v\n", !active)
	fmt.Printf("active && inactive: %v\n", active && inactive)
	fmt.Printf("active || inactive: %v\n", active || inactive)
	fmt.Println()

	// Part 8: Complex Numbers
	fmt.Println("--- Part 8: Complex Numbers ---")

	c1 := 1 + 2i
	c2 := complex(3, 4)

	fmt.Printf("Complex literal: %v\n", c1)
	fmt.Printf("Using complex(): %v\n", c2)
	fmt.Printf("Real part of c2: %v\n", real(c2))
	fmt.Printf("Imaginary part of c2: %v\n", imag(c2))
	fmt.Println()

	// Part 9: Zero Values
	fmt.Println("--- Part 9: Zero Values ---")

	var zeroInt int
	var zeroFloat float64
	var zeroString string
	var zeroBool bool

	fmt.Printf("Zero int: %d\n", zeroInt)
	fmt.Printf("Zero float64: %f\n", zeroFloat)
	fmt.Printf("Zero string: %q\n", zeroString)
	fmt.Printf("Zero bool: %v\n", zeroBool)
	fmt.Println()

	// Part 10: Type Conversions
	fmt.Println("--- Part 10: Type Conversions ---")

	// int to float64
	intVal := 42
	floatVal := float64(intVal)
	fmt.Printf("int to float64: %d -> %f\n", intVal, floatVal)

	// float64 to int (loses precision)
	floatVal2 := 3.99
	intVal2 := int(floatVal2)
	fmt.Printf("float64 to int: %f -> %d (precision lost)\n", floatVal2, intVal2)

	// int to rune (Unicode code point)
	runeVal := rune(65)
	fmt.Printf("int to rune (Unicode code point): 65 -> %q\n", runeVal)

	// String conversions using strconv package
	strVal := strconv.Itoa(42)
	fmt.Printf("int to string using strconv.Itoa: 42 -> %q\n", strVal)

	intVal3, _ := strconv.Atoi("42")
	fmt.Printf("string to int using strconv.Atoi: \"42\" -> %d\n", intVal3)

	floatVal3, _ := strconv.ParseFloat("3.14", 64)
	fmt.Printf("string to float64 using strconv.ParseFloat: \"3.14\" -> %f\n", floatVal3)
	fmt.Println()

	fmt.Println("=== End of Day 1 Examples ===")
}

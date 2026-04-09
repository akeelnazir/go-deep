package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Counter struct {
	count int
}

func (c Counter) Get() int {
	return c.count
}

func (c *Counter) Increment() {
	c.count++
}

type Reader interface {
	Read() string
}

type File struct {
	Name    string
	Content string
}

func (f File) Read() string {
	return fmt.Sprintf("Reading from %s: %s", f.Name, f.Content)
}

type Database struct {
	Name string
	Data string
}

func (d Database) Read() string {
	return fmt.Sprintf("Query from %s: %s", d.Name, d.Data)
}

func printContent(r Reader) {
	fmt.Println(r.Read())
}

func describe(v interface{}) {
	switch val := v.(type) {
	case string:
		fmt.Printf("String: %s\n", val)
	case int:
		fmt.Printf("Integer: %d\n", val)
	case float64:
		fmt.Printf("Float: %.2f\n", val)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("Validation error in %s: %s", e.Field, e.Message)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, ValidationError{Field: "divisor", Message: "cannot be zero"}
	}
	return a / b, nil
}

type Writer interface {
	Write(data string) error
}

type ReadWriter interface {
	Reader
	Writer
}

type FileRW struct {
	Name    string
	Content string
}

func (f FileRW) Read() string {
	return f.Content
}

func (f *FileRW) Write(data string) error {
	f.Content = data
	fmt.Printf("Written to %s: %s\n", f.Name, data)
	return nil
}

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	CardNumber string
}

func (cc CreditCard) Pay(amount float64) error {
	fmt.Printf("Paying $%.2f with credit card %s\n", amount, cc.CardNumber)
	return nil
}

type PayPal struct {
	Email string
}

func (pp PayPal) Pay(amount float64) error {
	fmt.Printf("Paying $%.2f with PayPal (%s)\n", amount, pp.Email)
	return nil
}

func main() {
	fmt.Println("=== Day 4: Pointers and Interfaces ===")

	// Section 1: Pointer Basics
	fmt.Println("\n--- Pointer Basics ---")

	x := 42
	fmt.Printf("Variable x: %d\n", x)
	fmt.Printf("Address of x: %p\n", &x)

	var p *int
	p = &x
	fmt.Printf("Pointer p: %p\n", p)
	fmt.Printf("Value at p: %d\n", *p)

	*p = 100
	fmt.Printf("After *p = 100, x = %d\n", x)

	fmt.Println()

	// Section 2: Pointers to Structs
	fmt.Println("--- Pointers to Structs ---")

	person := Person{Name: "Alice", Age: 30}
	fmt.Printf("Person: %+v\n", person)

	pPerson := &person
	fmt.Printf("Pointer to person: %p\n", pPerson)
	fmt.Printf("Access Name via pointer: %s\n", pPerson.Name)

	pPerson.Name = "Bob"
	fmt.Printf("After modification: %+v\n", person)

	fmt.Println()

	// Section 3: The new() Function
	fmt.Println("--- The new() Function ---")

	pInt := new(int)
	fmt.Printf("new(int) returns: %p\n", pInt)
	fmt.Printf("Value at pointer: %d\n", *pInt)

	*pInt = 42
	fmt.Printf("After *pInt = 42: %d\n", *pInt)

	pPerson2 := new(Person)
	pPerson2.Name = "Carol"
	pPerson2.Age = 25
	fmt.Printf("Person created with new(): %+v\n", *pPerson2)

	fmt.Println()

	// Section 4: The make() Function
	fmt.Println("--- The make() Function ---")

	slice := make([]int, 5)
	fmt.Printf("Slice created with make([]int, 5): %v\n", slice)
	fmt.Printf("Length: %d, Capacity: %d\n", len(slice), cap(slice))

	m := make(map[string]int)
	m["a"] = 1
	m["b"] = 2
	fmt.Printf("Map created with make: %v\n", m)

	ch := make(chan int, 1)
	fmt.Printf("Channel created with make: %v\n", ch)

	fmt.Println()

	// Section 5: Pointer Receivers and Methods
	fmt.Println("--- Pointer Receivers and Methods ---")

	counter := Counter{count: 0}
	fmt.Printf("Initial count: %d\n", counter.Get())

	counter.Increment()
	fmt.Printf("After Increment(): %d\n", counter.Get())

	counter.Increment()
	fmt.Printf("After second Increment(): %d\n", counter.Get())

	fmt.Println()

	// Section 6: Interfaces - Definition and Implementation
	fmt.Println("--- Interfaces: Definition and Implementation ---")

	file := File{Name: "data.txt", Content: "Hello, World!"}
	db := Database{Name: "UserDB", Data: "SELECT * FROM users"}

	fmt.Println(file.Read())
	fmt.Println(db.Read())

	fmt.Println()

	// Section 7: Using Interfaces for Polymorphism
	fmt.Println("--- Using Interfaces for Polymorphism ---")

	printContent(file)
	printContent(db)

	fmt.Println()

	// Section 8: Empty Interface
	fmt.Println("--- Empty Interface ---")

	var i interface{}
	fmt.Printf("i = 42: %v\n", i)

	i = 42
	fmt.Printf("i = 42: %v\n", i)

	i = "hello"
	fmt.Printf("i = \"hello\": %v\n", i)

	i = []int{1, 2, 3}
	fmt.Printf("i = []int{1, 2, 3}: %v\n", i)

	i = file
	fmt.Printf("i = file: %v\n", i)

	fmt.Println()

	// Section 9: Type Assertions
	fmt.Println("--- Type Assertions ---")

	var val interface{} = "Go"

	str, ok := val.(string)
	if ok {
		fmt.Printf("Assertion to string succeeded: %s\n", str)
	}

	num, ok := val.(int)
	if !ok {
		fmt.Printf("Assertion to int failed (expected)\n")
	}
	fmt.Printf("num value: %d\n", num)

	fmt.Println()

	// Section 10: Type Switches
	fmt.Println("--- Type Switches ---")

	describe("hello")
	describe(42)
	describe(3.14)
	describe(true)

	fmt.Println()

	// Section 11: Custom Error Types
	fmt.Println("--- Custom Error Types ---")

	result, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	fmt.Println()

	// Section 12: Interface Composition
	fmt.Println("--- Interface Composition ---")

	fileRW := FileRW{Name: "test.txt", Content: "initial"}
	fmt.Printf("Read: %s\n", fileRW.Read())
	fileRW.Write("updated content")

	fmt.Println()

	// Section 13: Packages and Visibility
	fmt.Println("--- Packages and Visibility ---")

	fmt.Println("Exported identifiers (start with uppercase):")
	fmt.Println("  - fmt.Println (exported function)")
	fmt.Println("  - Person (exported type)")
	fmt.Println("  - Reader (exported interface)")

	fmt.Println("Unexported identifiers (start with lowercase):")
	fmt.Println("  - printContent (unexported function)")
	fmt.Println("  - describe (unexported function)")

	fmt.Println()

	// Section 14: Practical Example - Payment System
	fmt.Println("--- Practical Example: Payment System ---")

	methods := []PaymentMethod{
		CreditCard{CardNumber: "1234-5678-9012-3456"},
		PayPal{Email: "user@example.com"},
	}

	for _, method := range methods {
		method.Pay(100.00)
	}

	fmt.Println("\n=== Day 4 Complete ===")
	fmt.Println("Next: Learn about error handling and project structure on Day 5.")
}

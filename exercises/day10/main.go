package main

import (
	"fmt"
	"reflect"
)

// Section 1: Type Inspection with Reflection
func inspectType(v interface{}) {
	t := reflect.TypeOf(v)
	fmt.Printf("Type: %v, Kind: %v\n", t, t.Kind())
}

// Section 2: Value Inspection with Reflection
func inspectValue(v interface{}) {
	val := reflect.ValueOf(v)
	fmt.Printf("Value: %v, Type: %v, Kind: %v\n", val, val.Type(), val.Kind())
}

// Section 3: Struct Introspection
type Person struct {
	Name string
	Age  int
}

func inspectStruct(v interface{}) {
	t := reflect.TypeOf(v)
	val := reflect.ValueOf(v)

	fmt.Printf("%s{\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := val.Field(i)
		fmt.Printf("  %s: %v (%v)\n", field.Name, value, field.Type)
	}
	fmt.Println("}")
}

// Section 4: Modifying Values with Reflection
func modifyValue(v interface{}) {
	val := reflect.ValueOf(v).Elem()

	if val.Kind() == reflect.Struct {
		field := val.FieldByName("Name")
		if field.IsValid() && field.CanSet() {
			field.SetString("Modified")
		}
	}
}

// Section 5: Generic Min Function
func Min[T interface {
	int | int64 | float64 | string
}](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// Section 6: Generic Stack Type
type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

// Section 7: Generic Constraint
type Numeric interface {
	int | int64 | float64
}

func Add[T Numeric](a, b T) T {
	return a + b
}

func main() {
	fmt.Println("=== Day 10: Reflection and Generics ===")

	// Section 1: Type Inspection
	fmt.Println("\n--- Type Inspection ---")
	inspectType(42)
	inspectType("hello")
	inspectType([]int{1, 2, 3})

	// Section 2: Value Inspection
	fmt.Println("\n--- Value Inspection ---")
	inspectValue(42)
	inspectValue("world")

	// Section 3: Struct Introspection
	fmt.Println("\n--- Struct Introspection ---")
	p := Person{Name: "Alice", Age: 30}
	inspectStruct(p)

	// Section 4: Modifying Values
	fmt.Println("\n--- Modifying Values ---")
	p2 := &Person{Name: "Bob", Age: 25}
	fmt.Printf("Before: %s\n", p2.Name)
	modifyValue(p2)
	fmt.Printf("After: %s\n", p2.Name)

	// Section 5: Generic Min Function
	fmt.Println("\n--- Generic Min Function ---")
	fmt.Printf("Min(3, 5) = %v\n", Min(3, 5))
	fmt.Printf("Min(1.5, 2.5) = %v\n", Min(1.5, 2.5))
	fmt.Printf("Min(\"apple\", \"banana\") = %v\n", Min("apple", "banana"))

	// Section 6: Generic Stack
	fmt.Println("\n--- Generic Stack ---")
	intStack := &Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	intStack.Push(3)

	fmt.Printf("Stack length: %d\n", intStack.Len())
	for intStack.Len() > 0 {
		val, ok := intStack.Pop()
		fmt.Printf("Popped: %v (ok: %v)\n", val, ok)
	}

	// Section 7: Generic Constraint
	fmt.Println("\n--- Generic Constraint ---")
	fmt.Printf("Add(1, 2) = %v\n", Add(1, 2))
	fmt.Printf("Add(1.5, 2.5) = %v\n", Add(1.5, 2.5))

	fmt.Println("\n=== Day 10 Complete ===")
	fmt.Println("Next: Learn about web fundamentals on Day 11.")
}

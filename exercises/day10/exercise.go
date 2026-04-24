package main

// TODO: Implement ExerciseTypeInspection function
// Should use reflect.TypeOf to inspect the type of a value
// Return the Kind of the value as a string (e.g., "int", "string", "slice")
func ExerciseTypeInspection(v interface{}) string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseStructFields function
// Should use reflection to count the number of fields in a struct
// Given a Person struct with Name and Age fields, return the count
func ExerciseStructFields(v interface{}) int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseModifyStructField function
// Should use reflection to modify a struct field
// Given a pointer to a Person, set the Name field to "Updated"
// Return the new Name value
func ExerciseModifyStructField(v interface{}) string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseGenericMin function
// Should create a generic function that returns the minimum of two values
// Must work with int, float64, and string types
// Return the smaller of the two values
func ExerciseGenericMin[T interface {
	int | float64 | string
}](a, b T) T {
	// TODO: Add logic
	var zero T
	return zero
}

// TODO: Implement ExerciseGenericStack function
// Should create a generic Stack type and perform operations
// Create a Stack[string], push 3 items ("a", "b", "c"), then pop all
// Return the total number of items popped
func ExerciseGenericStack() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseGenericFilter function
// Should create a generic function that filters a slice based on a predicate
// Given a slice of integers and a predicate function, return filtered slice
// Example: filter [1,2,3,4,5] with predicate x > 2 returns [3,4,5]
func ExerciseGenericFilter[T any](items []T, predicate func(T) bool) []T {
	// TODO: Add logic
	return nil
}

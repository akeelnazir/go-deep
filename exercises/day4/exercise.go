package main

// TODO: Implement ExerciseDereferencePointer
// Should dereference the pointer and return the value it points to
func ExerciseDereferencePointer(p *int) int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseModifyThroughPointer
// Should modify the value that the pointer points to
func ExerciseModifyThroughPointer(p *int, newValue int) {
	// TODO: Add logic
}

// TODO: Implement ExerciseGetPointerAddress
// Should return the memory address of the given integer
func ExerciseGetPointerAddress(x int) *int {
	// TODO: Add logic
	return nil
}

// TODO: Implement ExerciseSwapValues
// Should swap the values of two integers using pointers
func ExerciseSwapValues(a, b *int) {
	// TODO: Add logic
}

// TODO: Implement ExerciseStructPointerField
// Should modify a struct field through a pointer
func ExerciseStructPointerField(p *Person, newName string) {
	// TODO: Add logic
}

// TODO: Implement ExerciseNewAllocation
// Should allocate a new Person using new() and return a pointer
func ExerciseNewAllocation(name string, age int) *Person {
	// TODO: Add logic
	return nil
}

// TODO: Implement ExerciseInterfaceImplementation
// Create a type that implements the Reader interface
// The type should have a Read() method that returns a string
type ExerciseReader struct {
	// TODO: Add fields
}

// TODO: Implement the Read() method for ExerciseReader
// Should return a string representation of the reader
func (er ExerciseReader) Read() string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseTypeAssertion
// Should perform a type assertion on the given interface{} value
// Return the string value if it's a string, otherwise return "not a string"
func ExerciseTypeAssertion(v interface{}) string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseCustomError
// Create a custom error type that implements the error interface
type ExerciseError struct {
	// TODO: Add fields
}

// TODO: Implement the Error() method for ExerciseError
// Should return a string describing the error
func (ee ExerciseError) Error() string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseValidateAge
// Should return an ExerciseError if age is negative, nil otherwise
func ExerciseValidateAge(age int) error {
	// TODO: Add logic
	return nil
}

// TODO: Implement ExercisePolymorphism
// Should accept any Reader interface and call its Read() method
// Return the result of calling Read()
func ExercisePolymorphism(r Reader) string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseEmptyInterface
// Should accept any value and return its type as a string
// Use type assertions or type switches to determine the type
func ExerciseEmptyInterface(v interface{}) string {
	// TODO: Add logic
	return ""
}

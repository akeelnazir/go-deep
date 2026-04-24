package main

// TODO: Implement ExerciseDivide function
// Should divide a by b and return the result and an error
// Return an error if b is 0 with message "division by zero"
func ExerciseDivide(a, b int) (int, error) {
	// TODO: Add logic
	return 0, nil
}

// TODO: Implement ExerciseParseInt function
// Should parse a string to an integer and return the result and an error
// Return an error with context using fmt.Errorf with %w if parsing fails
func ExerciseParseInt(s string) (int, error) {
	// TODO: Add logic
	return 0, nil
}

// TODO: Implement ExerciseValidateEmail function
// Should validate an email address and return an error if invalid
// Return a custom ValidationError with Field="email" if email doesn't contain "@"
// Return nil if email is valid
func ExerciseValidateEmail(email string) error {
	// TODO: Add logic
	return nil
}

// TODO: Implement ExerciseGetUserByID function
// Should retrieve a user by ID from the users map
// Return ErrInvalidInput if id is empty
// Return ErrNotFound if user doesn't exist
// Return the user and nil if found
// Use the users map and sentinel errors defined in main.go
func ExerciseGetUserByID(id string) (*User, error) {
	// TODO: Add logic
	return nil, nil
}

// TODO: Implement ExerciseSafeOperation function
// Should perform a division operation that recovers from panic
// If b is 0, the function will panic
// Use defer and recover to catch the panic
// Return the result of a/b if successful, or 0 if panic is recovered
func ExerciseSafeOperation(a, b int) (result int) {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseWrapError function
// Should wrap an error with additional context
// Take an error and a context message
// Return a new error that wraps the original error using fmt.Errorf with %w
// Example: if err is "file not found", return "failed to process: file not found"
func ExerciseWrapError(err error, context string) error {
	// TODO: Add logic
	return nil
}

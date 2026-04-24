package main

// TODO: Implement ExerciseCircuitBreakerCall function
// Should call a function through a circuit breaker
// Return the error from the function or circuit breaker error
func ExerciseCircuitBreakerCall(cb *CircuitBreaker, fn func() error) error {
	// TODO: Add logic
	return nil
}

// TODO: Implement ExerciseRetryOperation function
// Should retry an operation up to maxRetries times
// Return the error if all retries fail, nil if successful
func ExerciseRetryOperation(fn func() error, maxRetries int) error {
	// TODO: Add logic
	return nil
}

// TODO: Implement ExerciseGetCircuitBreakerState function
// Should return the current state of a circuit breaker
// Return "closed", "open", or "half-open"
func ExerciseGetCircuitBreakerState(cb *CircuitBreaker) string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseResetCircuitBreaker function
// Should reset a circuit breaker to closed state
func ExerciseResetCircuitBreaker(cb *CircuitBreaker) {
	// TODO: Add logic
}

// TODO: Implement ExerciseCountFailures function
// Should count the number of failures in a circuit breaker
// Return the failure count
func ExerciseCountFailures(cb *CircuitBreaker) int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseIsCircuitBreakerOpen function
// Should check if a circuit breaker is open
// Return true if open, false otherwise
func ExerciseIsCircuitBreakerOpen(cb *CircuitBreaker) bool {
	// TODO: Add logic
	return false
}

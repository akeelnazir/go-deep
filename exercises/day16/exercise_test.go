package main

import (
	"fmt"
	"testing"
	"time"
)

func TestExerciseCircuitBreakerCall(t *testing.T) {
	cb := NewCircuitBreaker(2, time.Second)
	
	callCount := 0
	fn := func() error {
		callCount++
		if callCount <= 2 {
			return fmt.Errorf("error")
		}
		return nil
	}

	ExerciseCircuitBreakerCall(cb, fn)
	ExerciseCircuitBreakerCall(cb, fn)
	err := ExerciseCircuitBreakerCall(cb, fn)
	
	if err == nil {
		t.Errorf("Expected circuit breaker to be open")
	}
}

func TestExerciseRetryOperation(t *testing.T) {
	callCount := 0
	fn := func() error {
		callCount++
		if callCount < 3 {
			return fmt.Errorf("error")
		}
		return nil
	}

	err := ExerciseRetryOperation(fn, 5)
	if err != nil {
		t.Errorf("ExerciseRetryOperation() = %v, want nil", err)
	}

	if callCount != 3 {
		t.Errorf("Expected 3 calls, got %d", callCount)
	}
}

func TestExerciseGetCircuitBreakerState(t *testing.T) {
	cb := NewCircuitBreaker(2, time.Second)
	
	state := ExerciseGetCircuitBreakerState(cb)
	if state != "closed" {
		t.Errorf("ExerciseGetCircuitBreakerState() = %q, want %q", state, "closed")
	}
}

func TestExerciseResetCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(1, time.Second)
	
	fn := func() error {
		return fmt.Errorf("error")
	}
	
	ExerciseCircuitBreakerCall(cb, fn)
	
	if cb.State() != "open" {
		t.Errorf("Expected circuit breaker to be open")
	}
	
	ExerciseResetCircuitBreaker(cb)
	
	if cb.State() != "closed" {
		t.Errorf("Expected circuit breaker to be closed after reset")
	}
}

func TestExerciseCountFailures(t *testing.T) {
	cb := NewCircuitBreaker(5, time.Second)
	
	fn := func() error {
		return fmt.Errorf("error")
	}
	
	ExerciseCircuitBreakerCall(cb, fn)
	ExerciseCircuitBreakerCall(cb, fn)
	
	count := ExerciseCountFailures(cb)
	if count != 2 {
		t.Errorf("ExerciseCountFailures() = %d, want 2", count)
	}
}

func TestExerciseIsCircuitBreakerOpen(t *testing.T) {
	cb := NewCircuitBreaker(1, time.Second)
	
	if ExerciseIsCircuitBreakerOpen(cb) {
		t.Errorf("Expected circuit breaker to be closed initially")
	}
	
	fn := func() error {
		return fmt.Errorf("error")
	}
	
	ExerciseCircuitBreakerCall(cb, fn)
	
	if !ExerciseIsCircuitBreakerOpen(cb) {
		t.Errorf("Expected circuit breaker to be open after failure")
	}
}

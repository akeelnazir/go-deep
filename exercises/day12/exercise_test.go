package main

import (
	"testing"
)

func TestExerciseBasicRouter(t *testing.T) {
	result := ExerciseBasicRouter()
	expected := 3
	if result != expected {
		t.Errorf("ExerciseBasicRouter() = %d, want %d", result, expected)
	}
}

func TestExerciseMiddlewareChain(t *testing.T) {
	result := ExerciseMiddlewareChain()
	expected := 1
	if result != expected {
		t.Errorf("ExerciseMiddlewareChain() = %d, want %d", result, expected)
	}
}

func TestExerciseAuthMiddleware(t *testing.T) {
	result := ExerciseAuthMiddleware()
	expected := 1
	if result != expected {
		t.Errorf("ExerciseAuthMiddleware() = %d, want %d", result, expected)
	}
}

func TestExerciseCORSMiddleware(t *testing.T) {
	result := ExerciseCORSMiddleware()
	expected := 1
	if result != expected {
		t.Errorf("ExerciseCORSMiddleware() = %d, want %d", result, expected)
	}
}

func TestExerciseContextPropagation(t *testing.T) {
	result := ExerciseContextPropagation()
	expected := 1
	if result != expected {
		t.Errorf("ExerciseContextPropagation() = %d, want %d", result, expected)
	}
}

func TestExerciseRecoveryMiddleware(t *testing.T) {
	result := ExerciseRecoveryMiddleware()
	expected := 1
	if result != expected {
		t.Errorf("ExerciseRecoveryMiddleware() = %d, want %d", result, expected)
	}
}

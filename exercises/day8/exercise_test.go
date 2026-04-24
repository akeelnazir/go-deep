package main

import (
	"testing"
)

func TestExerciseContextCancellation(t *testing.T) {
	result := ExerciseContextCancellation()
	if result <= 0 {
		t.Errorf("ExerciseContextCancellation() = %d, want > 0", result)
	}
}

func TestExerciseContextTimeout(t *testing.T) {
	result := ExerciseContextTimeout()
	if !result {
		t.Errorf("ExerciseContextTimeout() = %v, want true (timeout should occur)", result)
	}
}

func TestExerciseContextDeadline(t *testing.T) {
	err := ExerciseContextDeadline()
	if err != nil {
		t.Errorf("ExerciseContextDeadline() = %v, want nil (operation should complete before deadline)", err)
	}
}

func TestExerciseWorkerPoolWithContext(t *testing.T) {
	result := ExerciseWorkerPoolWithContext()
	expected := 4
	if result != expected {
		t.Errorf("ExerciseWorkerPoolWithContext() = %d, want %d", result, expected)
	}
}

func TestExercisePipelineWithContext(t *testing.T) {
	result := ExercisePipelineWithContext()
	expected := 36
	if result != expected {
		t.Errorf("ExercisePipelineWithContext() = %d, want %d", result, expected)
	}
}

func TestExerciseFanOutFanIn(t *testing.T) {
	result := ExerciseFanOutFanIn()
	expected := 4
	if result != expected {
		t.Errorf("ExerciseFanOutFanIn() = %d, want %d", result, expected)
	}
}

func TestExerciseGracefulShutdown(t *testing.T) {
	result := ExerciseGracefulShutdown()
	expected := 2
	if result != expected {
		t.Errorf("ExerciseGracefulShutdown() = %d, want %d", result, expected)
	}
}

func TestExerciseContextWithValue(t *testing.T) {
	result := ExerciseContextWithValue()
	expected := "alice"
	if result != expected {
		t.Errorf("ExerciseContextWithValue() = %q, want %q", result, expected)
	}
}

func TestExerciseMultipleContexts(t *testing.T) {
	result := ExerciseMultipleContexts()
	if !result {
		t.Errorf("ExerciseMultipleContexts() = %v, want true (both contexts should be cancelled)", result)
	}
}

func TestExerciseContextPropagation(t *testing.T) {
	result := ExerciseContextPropagation()
	if result < 1 || result > 3 {
		t.Errorf("ExerciseContextPropagation() = %d, want between 1 and 3", result)
	}
}

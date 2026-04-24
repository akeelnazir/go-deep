package main

import (
	"testing"
)

func TestExerciseMutexCounter(t *testing.T) {
	result := ExerciseMutexCounter()
	expected := 500
	if result != expected {
		t.Errorf("ExerciseMutexCounter() = %d, want %d", result, expected)
	}
}

func TestExerciseRWMutexCache(t *testing.T) {
	result := ExerciseRWMutexCache()
	expected := "value2"
	if result != expected {
		t.Errorf("ExerciseRWMutexCache() = %q, want %q", result, expected)
	}
}

func TestExerciseAtomicCounter(t *testing.T) {
	result := ExerciseAtomicCounter()
	expected := int64(500)
	if result != expected {
		t.Errorf("ExerciseAtomicCounter() = %d, want %d", result, expected)
	}
}

func TestExerciseWorkerPool(t *testing.T) {
	result := ExerciseWorkerPool()
	expected := 5
	if result != expected {
		t.Errorf("ExerciseWorkerPool() = %d, want %d", result, expected)
	}
}

func TestExerciseChannelWithMutex(t *testing.T) {
	result := ExerciseChannelWithMutex()
	expected := 10
	if result != expected {
		t.Errorf("ExerciseChannelWithMutex() = %d, want %d", result, expected)
	}
}

func TestExerciseWaitGroupWithChannel(t *testing.T) {
	result := ExerciseWaitGroupWithChannel()
	expected := 6
	if result != expected {
		t.Errorf("ExerciseWaitGroupWithChannel() = %d, want %d", result, expected)
	}
}

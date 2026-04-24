package main

import (
	"testing"
	"time"
)

func TestExerciseSimpleGoroutine(t *testing.T) {
	result := ExerciseSimpleGoroutine()
	expected := "Hello from goroutine"
	if result != expected {
		t.Errorf("ExerciseSimpleGoroutine() = %q, want %q", result, expected)
	}
}

func TestExerciseWaitGroupCounter(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"count 1", 1, 1},
		{"count 3", 3, 3},
		{"count 5", 5, 5},
		{"count 10", 10, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseWaitGroupCounter(tt.n)
			if got != tt.want {
				t.Errorf("ExerciseWaitGroupCounter(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestExerciseBufferedChannelSum(t *testing.T) {
	result := ExerciseBufferedChannelSum()
	expected := 6
	if result != expected {
		t.Errorf("ExerciseBufferedChannelSum() = %d, want %d", result, expected)
	}
}

func TestExerciseChannelRange(t *testing.T) {
	result := ExerciseChannelRange()
	expected := 15
	if result != expected {
		t.Errorf("ExerciseChannelRange() = %d, want %d", result, expected)
	}
}

func TestExerciseDirectionalChannels(t *testing.T) {
	result := ExerciseDirectionalChannels()
	expected := 6
	if result != expected {
		t.Errorf("ExerciseDirectionalChannels() = %d, want %d", result, expected)
	}
}

func TestExerciseDetectChannelClosure(t *testing.T) {
	result := ExerciseDetectChannelClosure()
	expected := 3
	if result != expected {
		t.Errorf("ExerciseDetectChannelClosure() = %d, want %d", result, expected)
	}
}

func TestExerciseSimpleGoroutineTimeout(t *testing.T) {
	done := make(chan bool)
	go func() {
		result := ExerciseSimpleGoroutine()
		done <- result == "Hello from goroutine"
	}()

	select {
	case success := <-done:
		if !success {
			t.Error("ExerciseSimpleGoroutine() did not return expected message")
		}
	case <-time.After(2 * time.Second):
		t.Error("ExerciseSimpleGoroutine() timed out")
	}
}

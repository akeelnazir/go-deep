package main

import (
	"testing"
)

func TestExerciseAllocateMemory(t *testing.T) {
	ExerciseResetMemoryStats()

	ExerciseAllocateMemory(1024)
	usage := ExerciseGetMemoryUsage()
	if usage != 1024 {
		t.Errorf("ExerciseGetMemoryUsage() = %d, want 1024", usage)
	}

	count := ExerciseGetAllocationCount()
	if count != 1 {
		t.Errorf("ExerciseGetAllocationCount() = %d, want 1", count)
	}
}

func TestExerciseDeallocateMemory(t *testing.T) {
	ExerciseResetMemoryStats()

	ExerciseAllocateMemory(1024)
	ExerciseDeallocateMemory(512)

	usage := ExerciseGetMemoryUsage()
	if usage != 512 {
		t.Errorf("ExerciseGetMemoryUsage() = %d, want 512", usage)
	}

	count := ExerciseGetDeallocationCount()
	if count != 1 {
		t.Errorf("ExerciseGetDeallocationCount() = %d, want 1", count)
	}
}

func TestExerciseGetMemoryUsage(t *testing.T) {
	ExerciseResetMemoryStats()

	ExerciseAllocateMemory(2048)
	usage := ExerciseGetMemoryUsage()
	if usage != 2048 {
		t.Errorf("ExerciseGetMemoryUsage() = %d, want 2048", usage)
	}
}

func TestExerciseGetAllocationCount(t *testing.T) {
	ExerciseResetMemoryStats()

	ExerciseAllocateMemory(1024)
	ExerciseAllocateMemory(512)

	count := ExerciseGetAllocationCount()
	if count != 2 {
		t.Errorf("ExerciseGetAllocationCount() = %d, want 2", count)
	}
}

func TestExerciseResetMemoryStats(t *testing.T) {
	ExerciseAllocateMemory(1024)
	ExerciseResetMemoryStats()

	usage := ExerciseGetMemoryUsage()
	if usage != 0 {
		t.Errorf("ExerciseGetMemoryUsage() = %d, want 0 after reset", usage)
	}

	count := ExerciseGetAllocationCount()
	if count != 0 {
		t.Errorf("ExerciseGetAllocationCount() = %d, want 0 after reset", count)
	}
}

package main

import (
	"testing"
)

func TestExercisePush(t *testing.T) {
	ll := &LinkedList{}
	ExercisePush(ll, 1)
	ExercisePush(ll, 2)

	if ll.Size() != 2 {
		t.Errorf("Expected size 2 after 2 pushes, got %d", ll.Size())
	}
}

func TestExercisePop(t *testing.T) {
	ll := &LinkedList{}
	ExercisePush(ll, 1)
	ExercisePush(ll, 2)

	value := ExercisePop(ll)
	if value != 2 {
		t.Errorf("ExercisePop() = %d, want 2", value)
	}

	if ll.Size() != 1 {
		t.Errorf("Expected size 1 after pop, got %d", ll.Size())
	}
}

func TestExerciseListSize(t *testing.T) {
	ll := &LinkedList{}

	size := ExerciseListSize(ll)
	if size != 0 {
		t.Errorf("ExerciseListSize() = %d, want 0", size)
	}

	ExercisePush(ll, 1)
	ExercisePush(ll, 2)

	size = ExerciseListSize(ll)
	if size != 2 {
		t.Errorf("ExerciseListSize() = %d, want 2", size)
	}
}

func TestExerciseMapValues(t *testing.T) {
	nums := []int{1, 2, 3}
	count := ExerciseMapValues(func(x int) int { return x * 2 }, nums)
	if count != 3 {
		t.Errorf("ExerciseMapValues() = %d, want 3", count)
	}
}

func TestExerciseFilterValues(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	count := ExerciseFilterValues(func(x int) bool { return x%2 == 0 }, nums)
	if count != 2 {
		t.Errorf("ExerciseFilterValues() = %d, want 2", count)
	}
}

func TestExerciseReduceValues(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	sum := ExerciseReduceValues(func(acc, x int) int { return acc + x }, nums, 0)
	if sum != 15 {
		t.Errorf("ExerciseReduceValues() = %d, want 15", sum)
	}

	product := ExerciseReduceValues(func(acc, x int) int { return acc * x }, nums, 1)
	if product != 120 {
		t.Errorf("ExerciseReduceValues() = %d, want 120", product)
	}
}

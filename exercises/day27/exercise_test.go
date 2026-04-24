package main

import (
	"testing"
)

func TestExerciseRegisterResource(t *testing.T) {
	resources = make(map[string]Resource)

	ExerciseRegisterResource("test.txt", "test content")
	if !ExerciseResourceExists("test.txt") {
		t.Errorf("Resource should exist after registration")
	}
}

func TestExerciseGetResource(t *testing.T) {
	resources = make(map[string]Resource)

	ExerciseRegisterResource("test.txt", "test content")

	result := ExerciseGetResource("test.txt")
	if result != "test content" {
		t.Errorf("ExerciseGetResource() = %q, want %q", result, "test content")
	}

	result = ExerciseGetResource("nonexistent.txt")
	if result != "" {
		t.Errorf("ExerciseGetResource() = %q, want empty string", result)
	}
}

func TestExerciseResourceExists(t *testing.T) {
	resources = make(map[string]Resource)

	ExerciseRegisterResource("test.txt", "content")

	if !ExerciseResourceExists("test.txt") {
		t.Errorf("ExerciseResourceExists() = false, want true")
	}

	if ExerciseResourceExists("nonexistent.txt") {
		t.Errorf("ExerciseResourceExists() = true, want false")
	}
}

func TestExerciseCountResources(t *testing.T) {
	resources = make(map[string]Resource)

	count := ExerciseCountResources()
	if count != 0 {
		t.Errorf("ExerciseCountResources() = %d, want 0", count)
	}

	ExerciseRegisterResource("test1.txt", "content1")
	ExerciseRegisterResource("test2.txt", "content2")

	count = ExerciseCountResources()
	if count != 2 {
		t.Errorf("ExerciseCountResources() = %d, want 2", count)
	}
}

func TestExerciseDeleteResource(t *testing.T) {
	resources = make(map[string]Resource)

	ExerciseRegisterResource("test.txt", "content")

	result := ExerciseDeleteResource("test.txt")
	if !result {
		t.Errorf("ExerciseDeleteResource() = %v, want true", result)
	}

	if ExerciseResourceExists("test.txt") {
		t.Errorf("Resource should not exist after deletion")
	}

	result = ExerciseDeleteResource("nonexistent.txt")
	if result {
		t.Errorf("ExerciseDeleteResource() = %v, want false for nonexistent", result)
	}
}

func TestExerciseGetResourceSize(t *testing.T) {
	resources = make(map[string]Resource)

	ExerciseRegisterResource("test.txt", "hello")

	size := ExerciseGetResourceSize("test.txt")
	if size != 5 {
		t.Errorf("ExerciseGetResourceSize() = %d, want 5", size)
	}

	size = ExerciseGetResourceSize("nonexistent.txt")
	if size != -1 {
		t.Errorf("ExerciseGetResourceSize() = %d, want -1", size)
	}
}

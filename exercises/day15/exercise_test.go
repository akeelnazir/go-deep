package main

import (
	"testing"
)

func TestExerciseInsertUser(t *testing.T) {
	db = make(map[int]User)
	nextID = 1

	id := ExerciseInsertUser("Alice", "alice@example.com")
	if id != 1 {
		t.Errorf("ExerciseInsertUser() = %d, want 1", id)
	}

	id = ExerciseInsertUser("Bob", "bob@example.com")
	if id != 2 {
		t.Errorf("ExerciseInsertUser() = %d, want 2", id)
	}
}

func TestExerciseQueryUser(t *testing.T) {
	db = make(map[int]User)
	nextID = 1

	ExerciseInsertUser("Alice", "alice@example.com")

	result := ExerciseQueryUser(1)
	if result != "Alice" {
		t.Errorf("ExerciseQueryUser(1) = %q, want %q", result, "Alice")
	}

	result = ExerciseQueryUser(999)
	if result != "" {
		t.Errorf("ExerciseQueryUser(999) = %q, want empty string", result)
	}
}

func TestExerciseUpdateUser(t *testing.T) {
	db = make(map[int]User)
	nextID = 1

	ExerciseInsertUser("Alice", "alice@example.com")

	result := ExerciseUpdateUser(1, "Alice Updated", "alice.updated@example.com")
	if !result {
		t.Errorf("ExerciseUpdateUser(1) = %v, want true", result)
	}

	result = ExerciseUpdateUser(999, "Bob", "bob@example.com")
	if result {
		t.Errorf("ExerciseUpdateUser(999) = %v, want false", result)
	}
}

func TestExerciseDeleteUser(t *testing.T) {
	db = make(map[int]User)
	nextID = 1

	ExerciseInsertUser("Alice", "alice@example.com")

	result := ExerciseDeleteUser(1)
	if !result {
		t.Errorf("ExerciseDeleteUser(1) = %v, want true", result)
	}

	result = ExerciseDeleteUser(999)
	if result {
		t.Errorf("ExerciseDeleteUser(999) = %v, want false", result)
	}
}

func TestExerciseCountUsers(t *testing.T) {
	db = make(map[int]User)
	nextID = 1

	count := ExerciseCountUsers()
	if count != 0 {
		t.Errorf("ExerciseCountUsers() = %d, want 0", count)
	}

	ExerciseInsertUser("Alice", "alice@example.com")
	ExerciseInsertUser("Bob", "bob@example.com")

	count = ExerciseCountUsers()
	if count != 2 {
		t.Errorf("ExerciseCountUsers() = %d, want 2", count)
	}
}

func TestExerciseQueryUserByEmail(t *testing.T) {
	db = make(map[int]User)
	nextID = 1

	id := ExerciseInsertUser("Alice", "alice@example.com")

	result := ExerciseQueryUserByEmail("alice@example.com")
	if result != id {
		t.Errorf("ExerciseQueryUserByEmail() = %d, want %d", result, id)
	}

	result = ExerciseQueryUserByEmail("nonexistent@example.com")
	if result != -1 {
		t.Errorf("ExerciseQueryUserByEmail() = %d, want -1", result)
	}
}

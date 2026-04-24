package main

import (
	"testing"
)

func TestExerciseCreateUser(t *testing.T) {
	userService.users = make(map[int]User)
	userService.nextID = 1

	id := ExerciseCreateUser("Alice", "alice@example.com")
	if id != 1 {
		t.Errorf("ExerciseCreateUser() = %d, want 1", id)
	}

	id = ExerciseCreateUser("Bob", "bob@example.com")
	if id != 2 {
		t.Errorf("ExerciseCreateUser() = %d, want 2", id)
	}
}

func TestExerciseGetUser(t *testing.T) {
	userService.users = make(map[int]User)
	userService.nextID = 1

	ExerciseCreateUser("Alice", "alice@example.com")

	result := ExerciseGetUser(1)
	if result != "Alice" {
		t.Errorf("ExerciseGetUser(1) = %q, want %q", result, "Alice")
	}

	result = ExerciseGetUser(999)
	if result != "" {
		t.Errorf("ExerciseGetUser(999) = %q, want empty string", result)
	}
}

func TestExerciseUpdateUser(t *testing.T) {
	userService.users = make(map[int]User)
	userService.nextID = 1

	ExerciseCreateUser("Alice", "alice@example.com")

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
	userService.users = make(map[int]User)
	userService.nextID = 1

	ExerciseCreateUser("Alice", "alice@example.com")

	result := ExerciseDeleteUser(1)
	if !result {
		t.Errorf("ExerciseDeleteUser(1) = %v, want true", result)
	}

	result = ExerciseDeleteUser(999)
	if result {
		t.Errorf("ExerciseDeleteUser(999) = %v, want false", result)
	}
}

func TestExerciseListUsers(t *testing.T) {
	userService.users = make(map[int]User)
	userService.nextID = 1

	count := ExerciseListUsers()
	if count != 0 {
		t.Errorf("ExerciseListUsers() = %d, want 0", count)
	}

	ExerciseCreateUser("Alice", "alice@example.com")
	ExerciseCreateUser("Bob", "bob@example.com")

	count = ExerciseListUsers()
	if count != 2 {
		t.Errorf("ExerciseListUsers() = %d, want 2", count)
	}
}

func TestExerciseUserExists(t *testing.T) {
	userService.users = make(map[int]User)
	userService.nextID = 1

	ExerciseCreateUser("Alice", "alice@example.com")

	if !ExerciseUserExists(1) {
		t.Errorf("ExerciseUserExists(1) = false, want true")
	}

	if ExerciseUserExists(999) {
		t.Errorf("ExerciseUserExists(999) = true, want false")
	}
}

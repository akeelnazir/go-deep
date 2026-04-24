package main

import (
	"testing"
)

func TestExerciseListUsers(t *testing.T) {
	users = make(map[int]User)
	nextID = 1

	result := ExerciseListUsers()
	if result != 0 {
		t.Errorf("ExerciseListUsers() = %d, want 0", result)
	}

	users[1] = User{ID: 1, Name: "Alice", Email: "alice@example.com"}
	users[2] = User{ID: 2, Name: "Bob", Email: "bob@example.com"}

	result = ExerciseListUsers()
	if result != 2 {
		t.Errorf("ExerciseListUsers() = %d, want 2", result)
	}
}

func TestExerciseCreateUser(t *testing.T) {
	users = make(map[int]User)
	nextID = 1

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
	users = make(map[int]User)
	nextID = 1

	users[1] = User{ID: 1, Name: "Alice", Email: "alice@example.com"}

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
	users = make(map[int]User)
	nextID = 1

	users[1] = User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	result := ExerciseUpdateUser(1, "Alice Updated", "alice.updated@example.com")
	if !result {
		t.Errorf("ExerciseUpdateUser(1) = %v, want true", result)
	}

	user := users[1]
	if user.Name != "Alice Updated" {
		t.Errorf("User name = %q, want %q", user.Name, "Alice Updated")
	}

	result = ExerciseUpdateUser(999, "Bob", "bob@example.com")
	if result {
		t.Errorf("ExerciseUpdateUser(999) = %v, want false", result)
	}
}

func TestExerciseDeleteUser(t *testing.T) {
	users = make(map[int]User)
	nextID = 1

	users[1] = User{ID: 1, Name: "Alice", Email: "alice@example.com"}

	result := ExerciseDeleteUser(1)
	if !result {
		t.Errorf("ExerciseDeleteUser(1) = %v, want true", result)
	}

	_, ok := users[1]
	if ok {
		t.Errorf("User 1 should be deleted")
	}

	result = ExerciseDeleteUser(999)
	if result {
		t.Errorf("ExerciseDeleteUser(999) = %v, want false", result)
	}
}

func TestExerciseValidateEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"alice@example.com", true},
		{"bob@test.org", true},
		{"invalid.email", false},
		{"@example.com", true},
		{"", false},
	}

	for _, tt := range tests {
		result := ExerciseValidateEmail(tt.email)
		if result != tt.valid {
			t.Errorf("ExerciseValidateEmail(%q) = %v, want %v", tt.email, result, tt.valid)
		}
	}
}

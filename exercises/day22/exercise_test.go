package main

import (
	"testing"
)

func TestExerciseCreateConnection(t *testing.T) {
	pool.connections = make(map[int]*Connection)
	pool.nextID = 1

	id := ExerciseCreateConnection("localhost:8080")
	if id == 0 {
		t.Errorf("ExerciseCreateConnection() = 0, want non-zero")
	}

	if !ExerciseConnectionExists(id) {
		t.Errorf("Connection should exist after creation")
	}
}

func TestExerciseCloseConnection(t *testing.T) {
	pool.connections = make(map[int]*Connection)
	pool.nextID = 1

	id := ExerciseCreateConnection("localhost:8080")
	result := ExerciseCloseConnection(id)
	if !result {
		t.Errorf("ExerciseCloseConnection() = %v, want true", result)
	}

	if ExerciseIsConnectionActive(id) {
		t.Errorf("Connection should not be active after close")
	}
}

func TestExerciseGetConnectionAddress(t *testing.T) {
	pool.connections = make(map[int]*Connection)
	pool.nextID = 1

	id := ExerciseCreateConnection("localhost:8080")
	address := ExerciseGetConnectionAddress(id)
	if address != "localhost:8080" {
		t.Errorf("ExerciseGetConnectionAddress() = %q, want %q", address, "localhost:8080")
	}
}

func TestExerciseCountActiveConnections(t *testing.T) {
	pool.connections = make(map[int]*Connection)
	pool.nextID = 1

	count := ExerciseCountActiveConnections()
	if count != 0 {
		t.Errorf("ExerciseCountActiveConnections() = %d, want 0", count)
	}

	id1 := ExerciseCreateConnection("localhost:8080")
	ExerciseCreateConnection("localhost:8081")

	count = ExerciseCountActiveConnections()
	if count != 2 {
		t.Errorf("ExerciseCountActiveConnections() = %d, want 2", count)
	}

	ExerciseCloseConnection(id1)
	count = ExerciseCountActiveConnections()
	if count != 1 {
		t.Errorf("ExerciseCountActiveConnections() = %d, want 1", count)
	}
}

func TestExerciseConnectionExists(t *testing.T) {
	pool.connections = make(map[int]*Connection)
	pool.nextID = 1

	id := ExerciseCreateConnection("localhost:8080")

	if !ExerciseConnectionExists(id) {
		t.Errorf("ExerciseConnectionExists() = false, want true")
	}

	if ExerciseConnectionExists(9999) {
		t.Errorf("ExerciseConnectionExists() = true, want false for nonexistent")
	}
}

func TestExerciseIsConnectionActive(t *testing.T) {
	pool.connections = make(map[int]*Connection)
	pool.nextID = 1

	id := ExerciseCreateConnection("localhost:8080")

	if !ExerciseIsConnectionActive(id) {
		t.Errorf("ExerciseIsConnectionActive() = false, want true")
	}

	ExerciseCloseConnection(id)

	if ExerciseIsConnectionActive(id) {
		t.Errorf("ExerciseIsConnectionActive() = true, want false after close")
	}
}

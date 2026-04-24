package main

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestExerciseTestHandler(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	status := ExerciseTestHandler(handler, "GET", "/test")
	if status != http.StatusOK {
		t.Errorf("ExerciseTestHandler() = %d, want %d", status, http.StatusOK)
	}
}

func TestExerciseTestHandlerBody(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test body"))
	}

	body := ExerciseTestHandlerBody(handler, "GET", "/test")
	if body != "test body" {
		t.Errorf("ExerciseTestHandlerBody() = %q, want %q", body, "test body")
	}
}

func TestExerciseTestMockServer(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}

	status := ExerciseTestMockServer(handler)
	if status != http.StatusOK {
		t.Errorf("ExerciseTestMockServer() = %d, want %d", status, http.StatusOK)
	}
}

func TestExerciseTestHandlerHeader(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom-Header", "test-value")
	}

	header := ExerciseTestHandlerHeader(handler, "GET", "/test", "X-Custom-Header")
	if header != "test-value" {
		t.Errorf("ExerciseTestHandlerHeader() = %q, want %q", header, "test-value")
	}
}

func TestExerciseTestHandlerJSON(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"key": "value"})
	}

	isJSON := ExerciseTestHandlerJSON(handler, "GET", "/test")
	if !isJSON {
		t.Errorf("ExerciseTestHandlerJSON() = %v, want true", isJSON)
	}
}

package main

import (
	"net/http"
	"testing"
)

func TestExerciseSimpleHandler(t *testing.T) {
	result := ExerciseSimpleHandler()
	expected := "Hello, Exercise!"
	if result != expected {
		t.Errorf("ExerciseSimpleHandler() = %q, want %q", result, expected)
	}
}

func TestExerciseMethodHandler(t *testing.T) {
	result := ExerciseMethodHandler()
	expected := "GET method"
	if result != expected {
		t.Errorf("ExerciseMethodHandler() = %q, want %q", result, expected)
	}
}

func TestExerciseQueryParameters(t *testing.T) {
	result := ExerciseQueryParameters()
	expected := "q=golang, limit=10"
	if result != expected {
		t.Errorf("ExerciseQueryParameters() = %q, want %q", result, expected)
	}
}

func TestExerciseJSONRequest(t *testing.T) {
	result := ExerciseJSONRequest()
	expected := "Bob"
	if result != expected {
		t.Errorf("ExerciseJSONRequest() = %q, want %q", result, expected)
	}
}

func TestExerciseJSONResponse(t *testing.T) {
	result := ExerciseJSONResponse()
	expected := true
	if result != expected {
		t.Errorf("ExerciseJSONResponse() = %v, want %v", result, expected)
	}
}

func TestExerciseStatusCodes(t *testing.T) {
	result := ExerciseStatusCodes()
	expected := http.StatusNotFound
	if result != expected {
		t.Errorf("ExerciseStatusCodes() = %d, want %d", result, expected)
	}
}

package main

import (
	"testing"
)

func TestExerciseMatchPattern(t *testing.T) {
	count := ExerciseMatchPattern(`\d+`, "abc123def456")
	if count != 2 {
		t.Errorf("ExerciseMatchPattern() = %d, want 2", count)
	}
}

func TestExerciseReplacePattern(t *testing.T) {
	result := ExerciseReplacePattern(`\d+`, "abc123def456", "X")
	if result != "abcXdefX" {
		t.Errorf("ExerciseReplacePattern() = %q, want %q", result, "abcXdefX")
	}
}

func TestExerciseSplitText(t *testing.T) {
	count := ExerciseSplitText(",", "a,b,c,d")
	if count != 4 {
		t.Errorf("ExerciseSplitText() = %d, want 4", count)
	}
}

func TestExerciseContainsPattern(t *testing.T) {
	result := ExerciseContainsPattern(`\d+`, "abc123")
	if !result {
		t.Errorf("ExerciseContainsPattern() = %v, want true", result)
	}

	result = ExerciseContainsPattern(`\d+`, "abc")
	if result {
		t.Errorf("ExerciseContainsPattern() = %v, want false", result)
	}
}

func TestExerciseExtractNumbers(t *testing.T) {
	count := ExerciseExtractNumbers("a1b2c3d4")
	if count != 4 {
		t.Errorf("ExerciseExtractNumbers() = %d, want 4", count)
	}
}

func TestExerciseValidateEmail(t *testing.T) {
	tests := []struct {
		email string
		valid bool
	}{
		{"user@example.com", true},
		{"test@test.org", true},
		{"invalid.email", false},
		{"@example.com", false},
	}

	for _, tt := range tests {
		result := ExerciseValidateEmail(tt.email)
		if result != tt.valid {
			t.Errorf("ExerciseValidateEmail(%q) = %v, want %v", tt.email, result, tt.valid)
		}
	}
}

package main

import (
	"errors"
	"testing"
)

func TestExerciseDivide(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
		errMsg  string
	}{
		{"divide 10 by 2", 10, 2, 5, false, ""},
		{"divide 20 by 4", 20, 4, 5, false, ""},
		{"divide by zero", 10, 0, 0, true, "division by zero"},
		{"divide 0 by 5", 0, 5, 0, false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExerciseDivide(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseDivide(%d, %d) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if err != nil && err.Error() != tt.errMsg {
				t.Errorf("ExerciseDivide(%d, %d) error message = %q, want %q", tt.a, tt.b, err.Error(), tt.errMsg)
			}
			if got != tt.want {
				t.Errorf("ExerciseDivide(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestExerciseParseInt(t *testing.T) {
	tests := []struct {
		name    string
		s       string
		want    int
		wantErr bool
	}{
		{"parse 42", "42", 42, false},
		{"parse 0", "0", 0, false},
		{"parse -10", "-10", -10, false},
		{"parse invalid", "abc", 0, true},
		{"parse empty", "", 0, true},
		{"parse with spaces", "  123  ", 123, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExerciseParseInt(tt.s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseParseInt(%q) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExerciseParseInt(%q) = %d, want %d", tt.s, got, tt.want)
			}
			if err != nil && !errors.Is(err, errors.New("strconv.Atoi: parsing")) {
				// Check that error is wrapped (contains context)
				if err.Error() == "" {
					t.Errorf("ExerciseParseInt(%q) error should be wrapped with context", tt.s)
				}
			}
		})
	}
}

func TestExerciseValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		email   string
		wantErr bool
		errType string
	}{
		{"valid email", "user@example.com", false, ""},
		{"valid email 2", "test@domain.co.uk", false, ""},
		{"missing @", "invalid.email", true, "ValidationError"},
		{"empty email", "", true, "ValidationError"},
		{"only @", "@", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ExerciseValidateEmail(tt.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseValidateEmail(%q) error = %v, wantErr %v", tt.email, err, tt.wantErr)
				return
			}
			if err != nil && tt.errType == "ValidationError" {
				if _, ok := err.(ValidationError); !ok {
					t.Errorf("ExerciseValidateEmail(%q) error type = %T, want ValidationError", tt.email, err)
				}
			}
		})
	}
}

func TestExerciseGetUserByID(t *testing.T) {
	tests := []struct {
		name    string
		id      string
		wantErr error
		wantNil bool
	}{
		{"get user 1", "1", nil, false},
		{"get user 2", "2", nil, false},
		{"empty id", "", ErrInvalidInput, true},
		{"nonexistent user", "999", ErrNotFound, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExerciseGetUserByID(tt.id)
			if tt.wantErr != nil {
				if err != tt.wantErr {
					t.Errorf("ExerciseGetUserByID(%q) error = %v, want %v", tt.id, err, tt.wantErr)
				}
			} else if err != nil {
				t.Errorf("ExerciseGetUserByID(%q) unexpected error = %v", tt.id, err)
			}

			if tt.wantNil && got != nil {
				t.Errorf("ExerciseGetUserByID(%q) = %v, want nil", tt.id, got)
			}
			if !tt.wantNil && got == nil {
				t.Errorf("ExerciseGetUserByID(%q) = nil, want non-nil", tt.id)
			}
		})
	}
}

func TestExerciseSafeOperation(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"normal division", 10, 2, 5},
		{"normal division 2", 20, 4, 5},
		{"division by zero (recovered)", 10, 0, 0},
		{"zero divided by 5", 0, 5, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseSafeOperation(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("ExerciseSafeOperation(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestExerciseWrapError(t *testing.T) {
	tests := []struct {
		name       string
		err        error
		context    string
		wantErr    bool
		wantString string
	}{
		{"wrap simple error", errors.New("file not found"), "failed to process", true, "failed to process"},
		{"wrap with context", errors.New("connection refused"), "database", true, "database"},
		{"nil error", nil, "context", false, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseWrapError(tt.err, tt.context)
			if (got != nil) != tt.wantErr {
				t.Errorf("ExerciseWrapError() error = %v, wantErr %v", got, tt.wantErr)
				return
			}
			if got != nil && tt.wantString != "" {
				if !errors.Is(got, tt.err) {
					t.Errorf("ExerciseWrapError() should wrap original error")
				}
				if !contains(got.Error(), tt.wantString) {
					t.Errorf("ExerciseWrapError() error message = %q, should contain %q", got.Error(), tt.wantString)
				}
			}
		})
	}
}

func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

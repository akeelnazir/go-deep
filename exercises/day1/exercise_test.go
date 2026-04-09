package main

import (
	"testing"
)

func TestExerciseDeclareVariable(t *testing.T) {
	got := ExerciseDeclareVariable()
	want := "Hello, Go!"
	if got != want {
		t.Errorf("ExerciseDeclareVariable() = %q, want %q", got, want)
	}
}

func TestExerciseShortDeclaration(t *testing.T) {
	got := ExerciseShortDeclaration()
	want := 42
	if got != want {
		t.Errorf("ExerciseShortDeclaration() = %d, want %d", got, want)
	}
}

func TestExerciseMultipleDeclarations(t *testing.T) {
	got := ExerciseMultipleDeclarations()
	want := 60
	if got != want {
		t.Errorf("ExerciseMultipleDeclarations() = %d, want %d", got, want)
	}
}

func TestExerciseConstantDeclaration(t *testing.T) {
	got := ExerciseConstantDeclaration()
	want := 10
	if got != want {
		t.Errorf("ExerciseConstantDeclaration() = %d, want %d", got, want)
	}
}

func TestExerciseZeroValues(t *testing.T) {
	got := ExerciseZeroValues()
	want := 0
	if got != want {
		t.Errorf("ExerciseZeroValues() = %d, want %d", got, want)
	}
}

func TestExerciseIntToFloat(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want float64
	}{
		{"positive", 42, 42.0},
		{"zero", 0, 0.0},
		{"negative", -10, -10.0},
		{"large", 1000, 1000.0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIntToFloat(tt.arg); got != tt.want {
				t.Errorf("ExerciseIntToFloat(%d) = %f, want %f", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseFloatToInt(t *testing.T) {
	tests := []struct {
		name string
		arg  float64
		want int
	}{
		{"positive", 42.7, 42},
		{"zero", 0.0, 0},
		{"negative", -10.9, -10},
		{"truncates", 3.99, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseFloatToInt(tt.arg); got != tt.want {
				t.Errorf("ExerciseFloatToInt(%f) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseIntToString(t *testing.T) {
	tests := []struct {
		name string
		arg  int
		want string
	}{
		{"positive", 42, "42"},
		{"zero", 0, "0"},
		{"negative", -10, "-10"},
		{"large", 1000, "1000"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIntToString(tt.arg); got != tt.want {
				t.Errorf("ExerciseIntToString(%d) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseStringToInt(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{"positive", "42", 42},
		{"zero", "0", 0},
		{"negative", "-10", -10},
		{"invalid", "abc", 0},
		{"empty", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseStringToInt(tt.arg); got != tt.want {
				t.Errorf("ExerciseStringToInt(%q) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseTypeInference(t *testing.T) {
	got := ExerciseTypeInference()
	want := "float64"
	if got != want {
		t.Errorf("ExerciseTypeInference() = %q, want %q", got, want)
	}
}

func TestExerciseComplexNumber(t *testing.T) {
	got := ExerciseComplexNumber()
	want := 3.0
	if got != want {
		t.Errorf("ExerciseComplexNumber() = %f, want %f", got, want)
	}
}

func TestExerciseStringLength(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want int
	}{
		{"empty", "", 0},
		{"single char", "a", 1},
		{"word", "hello", 5},
		{"sentence", "hello world", 11},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseStringLength(tt.arg); got != tt.want {
				t.Errorf("ExerciseStringLength(%q) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseStringConcatenation(t *testing.T) {
	tests := []struct {
		name   string
		first  string
		second string
		want   string
	}{
		{"simple", "Hello", "World", "Hello World"},
		{"names", "John", "Doe", "John Doe"},
		{"empty first", "", "World", " World"},
		{"empty second", "Hello", "", "Hello "},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseStringConcatenation(tt.first, tt.second); got != tt.want {
				t.Errorf("ExerciseStringConcatenation(%q, %q) = %q, want %q", tt.first, tt.second, got, tt.want)
			}
		})
	}
}

func TestExerciseBooleanLogic(t *testing.T) {
	tests := []struct {
		name string
		a    bool
		b    bool
		want bool
	}{
		{"both true", true, true, true},
		{"first true", true, false, false},
		{"second true", false, true, false},
		{"both false", false, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseBooleanLogic(tt.a, tt.b); got != tt.want {
				t.Errorf("ExerciseBooleanLogic(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestExerciseIsPalindrome(t *testing.T) {
	tests := []struct {
		input string
		want  bool
		name  string
	}{
		{"", true, "empty"},
		{"a", true, "single"},
		{"racecar", true, "palindrome odd"},
		{"level", true, "palindrome even"},
		{"hello", false, "not palindrome"},
		{"A man a plan a canal Panama", true, "phrase with spaces"},
		{"Was it a car or a cat I saw", true, "phrase with spaces and punctuation"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIsPalindrome(tt.input); got != tt.want {
				t.Errorf("ExerciseIsPalindrome(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestExerciseCalculateAverage(t *testing.T) {
	tests := []struct {
		input []float64
		want  float64
		name  string
	}{
		{[]float64{}, 0, "empty slice"},
		{[]float64{5}, 5, "single element"},
		{[]float64{1, 2, 3, 4, 5}, 3, "simple average"},
		{[]float64{1.5, 2.5, 3.0}, 7.0 / 3.0, "float average"},
		{[]float64{-1, 0, 1}, 0, "average with negatives"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseCalculateAverage(tt.input); got != tt.want {
				t.Errorf("ExerciseCalculateAverage(%v) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestExerciseFormatName(t *testing.T) {
	tests := []struct {
		input string
		want  string
		name  string
	}{
		{"John Doe", "Doe, John", "simple name"},
		{"Jane Marie Smith", "Smith, Jane Marie", "middle name"},
		{"Prince", "Prince", "single name"},
		{"", "", "empty string"},
		{"  John   Doe  ", "Doe, John", "extra spaces"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseFormatName(tt.input); got != tt.want {
				t.Errorf("ExerciseFormatName(%q) = %q, want %q", tt.input, got, tt.want)
			}
		})
	}
}

func TestExerciseUserService_GetUser(t *testing.T) {
	service := ExerciseUserService{}

	tests := []struct {
		userID   string
		wantName string
		wantErr  bool
		name     string
	}{
		{"1", "John Doe", false, "valid user 1"},
		{"2", "Jane Smith", false, "valid user 2"},
		{"3", "", true, "invalid user"},
		{"", "", true, "empty user ID"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotName, err := service.GetUser(tt.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseUserService.GetUser(%q) error = %v, wantErr %v", tt.userID, err, tt.wantErr)
				return
			}
			if gotName != tt.wantName {
				t.Errorf("ExerciseUserService.GetUser(%q) = %q, want %q", tt.userID, gotName, tt.wantName)
			}
		})
	}
}

func TestExerciseCounter(t *testing.T) {
	counter := ExerciseCounter{}
	if got := counter.Value(); got != 0 {
		t.Errorf("ExerciseCounter.Value() = %d, want 0", got)
	}

	counter.Increment()
	if got := counter.Value(); got != 1 {
		t.Errorf("ExerciseCounter.Value() after Increment = %d, want 1", got)
	}

	counter.Increment()
	counter.Increment()
	if got := counter.Value(); got != 3 {
		t.Errorf("ExerciseCounter.Value() after 3 increments = %d, want 3", got)
	}
}

func BenchmarkExerciseStringConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExerciseStringConcat("hello", "world")
	}
}

func BenchmarkExerciseStringBuilderConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ExerciseStringBuilderConcat("hello", "world")
	}
}

func BenchmarkExerciseCalculateSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < b.N; i++ {
		ExerciseCalculateSum(numbers)
	}
}

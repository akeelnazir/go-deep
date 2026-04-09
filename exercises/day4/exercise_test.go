package main

import (
	"testing"
)

func TestExerciseDereferencePointer(t *testing.T) {
	x := 42
	p := &x
	got := ExerciseDereferencePointer(p)
	if got != 42 {
		t.Errorf("ExerciseDereferencePointer(&42) = %d, want 42", got)
	}
}

func TestExerciseModifyThroughPointer(t *testing.T) {
	x := 10
	p := &x
	ExerciseModifyThroughPointer(p, 20)
	if x != 20 {
		t.Errorf("After ExerciseModifyThroughPointer(&10, 20), x = %d, want 20", x)
	}
}

func TestExerciseGetPointerAddress(t *testing.T) {
	x := 42
	p := ExerciseGetPointerAddress(x)
	if p == nil {
		t.Errorf("ExerciseGetPointerAddress(42) returned nil")
	}
	if *p != 42 {
		t.Errorf("ExerciseGetPointerAddress(42) points to %d, want 42", *p)
	}
}

func TestExerciseSwapValues(t *testing.T) {
	a := 10
	b := 20
	ExerciseSwapValues(&a, &b)
	if a != 20 || b != 10 {
		t.Errorf("After ExerciseSwapValues(&10, &20), a = %d, b = %d, want a = 20, b = 10", a, b)
	}
}

func TestExerciseStructPointerField(t *testing.T) {
	person := Person{Name: "Alice", Age: 30}
	ExerciseStructPointerField(&person, "Bob")
	if person.Name != "Bob" {
		t.Errorf("After ExerciseStructPointerField(&person, \"Bob\"), person.Name = %s, want Bob", person.Name)
	}
}

func TestExerciseNewAllocation(t *testing.T) {
	p := ExerciseNewAllocation("Carol", 25)
	if p == nil {
		t.Errorf("ExerciseNewAllocation(\"Carol\", 25) returned nil")
	}
	if p.Name != "Carol" || p.Age != 25 {
		t.Errorf("ExerciseNewAllocation(\"Carol\", 25) = {%s, %d}, want {Carol, 25}", p.Name, p.Age)
	}
}

func TestExerciseInterfaceImplementation(t *testing.T) {
	er := ExerciseReader{}
	result := er.Read()
	if result == "" {
		t.Errorf("ExerciseReader.Read() returned empty string")
	}
}

func TestExerciseTypeAssertion(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"string", "hello", "hello"},
		{"int", 42, "not a string"},
		{"float", 3.14, "not a string"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseTypeAssertion(tt.value)
			if got != tt.want {
				t.Errorf("ExerciseTypeAssertion(%v) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

func TestExerciseCustomError(t *testing.T) {
	err := ExerciseValidateAge(-5)
	if err == nil {
		t.Errorf("ExerciseValidateAge(-5) returned nil, want error")
	}

	err = ExerciseValidateAge(25)
	if err != nil {
		t.Errorf("ExerciseValidateAge(25) returned %v, want nil", err)
	}
}

func TestExerciseValidateAge(t *testing.T) {
	tests := []struct {
		name    string
		age     int
		wantErr bool
	}{
		{"negative age", -5, true},
		{"zero age", 0, false},
		{"positive age", 25, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ExerciseValidateAge(tt.age)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseValidateAge(%d) error = %v, wantErr %v", tt.age, err, tt.wantErr)
			}
		})
	}
}

func TestExercisePolymorphism(t *testing.T) {
	file := File{Name: "test.txt", Content: "Hello"}
	result := ExercisePolymorphism(file)
	if result == "" {
		t.Errorf("ExercisePolymorphism(file) returned empty string")
	}
}

func TestExerciseEmptyInterface(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		want  string
	}{
		{"string", "hello", "string"},
		{"int", 42, "int"},
		{"float64", 3.14, "float64"},
		{"bool", true, "bool"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseEmptyInterface(tt.value)
			if got != tt.want {
				t.Errorf("ExerciseEmptyInterface(%v) = %q, want %q", tt.value, got, tt.want)
			}
		})
	}
}

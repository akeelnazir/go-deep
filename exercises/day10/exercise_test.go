package main

import (
	"testing"
)

func TestExerciseTypeInspection(t *testing.T) {
	result := ExerciseTypeInspection(42)
	expected := "int"
	if result != expected {
		t.Errorf("ExerciseTypeInspection(42) = %q, want %q", result, expected)
	}

	result = ExerciseTypeInspection("hello")
	expected = "string"
	if result != expected {
		t.Errorf("ExerciseTypeInspection(\"hello\") = %q, want %q", result, expected)
	}

	result = ExerciseTypeInspection([]int{1, 2, 3})
	expected = "slice"
	if result != expected {
		t.Errorf("ExerciseTypeInspection([]int{1, 2, 3}) = %q, want %q", result, expected)
	}
}

func TestExerciseStructFields(t *testing.T) {
	p := Person{Name: "Alice", Age: 30}
	result := ExerciseStructFields(p)
	expected := 2
	if result != expected {
		t.Errorf("ExerciseStructFields(Person) = %d, want %d", result, expected)
	}
}

func TestExerciseModifyStructField(t *testing.T) {
	p := &Person{Name: "Original", Age: 25}
	result := ExerciseModifyStructField(p)
	expected := "Updated"
	if result != expected {
		t.Errorf("ExerciseModifyStructField() = %q, want %q", result, expected)
	}
	if p.Name != "Updated" {
		t.Errorf("Person.Name = %q, want %q", p.Name, "Updated")
	}
}

func TestExerciseGenericMin(t *testing.T) {
	result := ExerciseGenericMin(3, 5)
	expected := 3
	if result != expected {
		t.Errorf("ExerciseGenericMin(3, 5) = %d, want %d", result, expected)
	}

	resultFloat := ExerciseGenericMin(1.5, 2.5)
	expectedFloat := 1.5
	if resultFloat != expectedFloat {
		t.Errorf("ExerciseGenericMin(1.5, 2.5) = %f, want %f", resultFloat, expectedFloat)
	}

	resultStr := ExerciseGenericMin("apple", "banana")
	expectedStr := "apple"
	if resultStr != expectedStr {
		t.Errorf("ExerciseGenericMin(\"apple\", \"banana\") = %q, want %q", resultStr, expectedStr)
	}
}

func TestExerciseGenericStack(t *testing.T) {
	result := ExerciseGenericStack()
	expected := 3
	if result != expected {
		t.Errorf("ExerciseGenericStack() = %d, want %d", result, expected)
	}
}

func TestExerciseGenericFilter(t *testing.T) {
	items := []int{1, 2, 3, 4, 5}
	predicate := func(x int) bool { return x > 2 }
	result := ExerciseGenericFilter(items, predicate)

	expected := []int{3, 4, 5}
	if len(result) != len(expected) {
		t.Errorf("ExerciseGenericFilter() length = %d, want %d", len(result), len(expected))
	}

	for i, v := range result {
		if i >= len(expected) || v != expected[i] {
			t.Errorf("ExerciseGenericFilter() = %v, want %v", result, expected)
			break
		}
	}
}

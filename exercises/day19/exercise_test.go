package main

import (
	"encoding/json"
	"testing"
)

func TestExerciseMarshalJSON(t *testing.T) {
	data := map[string]interface{}{"name": "Alice", "age": 30}
	result := ExerciseMarshalJSON(data)
	
	if result == "" {
		t.Errorf("ExerciseMarshalJSON() returned empty string")
	}
	
	var parsed map[string]interface{}
	json.Unmarshal([]byte(result), &parsed)
	if parsed["name"] != "Alice" {
		t.Errorf("ExerciseMarshalJSON() did not preserve data")
	}
}

func TestExerciseUnmarshalJSON(t *testing.T) {
	jsonStr := `{"name":"Bob","age":25}`
	result := ExerciseUnmarshalJSON(jsonStr)
	
	if result == nil {
		t.Errorf("ExerciseUnmarshalJSON() returned nil")
	}
	
	if result["name"] != "Bob" {
		t.Errorf("ExerciseUnmarshalJSON() = %v, want Bob", result["name"])
	}
}

func TestExerciseMarshalPerson(t *testing.T) {
	result := ExerciseMarshalPerson("Alice", 30, "alice@example.com")
	
	if result == "" {
		t.Errorf("ExerciseMarshalPerson() returned empty string")
	}
	
	var p Person
	json.Unmarshal([]byte(result), &p)
	if p.Name != "Alice" || p.Age != 30 {
		t.Errorf("ExerciseMarshalPerson() did not marshal correctly")
	}
}

func TestExerciseUnmarshalPerson(t *testing.T) {
	jsonStr := `{"name":"Bob","age":25,"email":"bob@example.com"}`
	result := ExerciseUnmarshalPerson(jsonStr)
	
	if result != "Bob" {
		t.Errorf("ExerciseUnmarshalPerson() = %q, want %q", result, "Bob")
	}
}

func TestExerciseValidateJSON(t *testing.T) {
	validJSON := `{"name":"Alice"}`
	if !ExerciseValidateJSON(validJSON) {
		t.Errorf("ExerciseValidateJSON() = false, want true for valid JSON")
	}
	
	invalidJSON := `{invalid}`
	if ExerciseValidateJSON(invalidJSON) {
		t.Errorf("ExerciseValidateJSON() = true, want false for invalid JSON")
	}
}

func TestExercisePrettyPrintJSON(t *testing.T) {
	jsonStr := `{"name":"Alice","age":30}`
	result := ExercisePrettyPrintJSON(jsonStr)
	
	if result == "" {
		t.Errorf("ExercisePrettyPrintJSON() returned empty string")
	}
	
	if result == jsonStr {
		t.Errorf("ExercisePrettyPrintJSON() did not format the JSON")
	}
}

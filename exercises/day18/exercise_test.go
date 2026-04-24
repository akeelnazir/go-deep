package main

import (
	"testing"
)

func TestExerciseWriteFile(t *testing.T) {
	fileSystem = make(map[string]string)

	result := ExerciseWriteFile("test.txt", "Hello")
	if !result {
		t.Errorf("ExerciseWriteFile() = %v, want true", result)
	}

	if !fileExists("test.txt") {
		t.Errorf("File should exist after write")
	}
}

func TestExerciseReadFile(t *testing.T) {
	fileSystem = make(map[string]string)
	writeFile("test.txt", "Hello")

	result := ExerciseReadFile("test.txt")
	if result != "Hello" {
		t.Errorf("ExerciseReadFile() = %q, want %q", result, "Hello")
	}

	result = ExerciseReadFile("nonexistent.txt")
	if result != "" {
		t.Errorf("ExerciseReadFile() = %q, want empty string", result)
	}
}

func TestExerciseFileExists(t *testing.T) {
	fileSystem = make(map[string]string)
	writeFile("test.txt", "Hello")

	result := ExerciseFileExists("test.txt")
	if !result {
		t.Errorf("ExerciseFileExists() = %v, want true", result)
	}

	result = ExerciseFileExists("nonexistent.txt")
	if result {
		t.Errorf("ExerciseFileExists() = %v, want false", result)
	}
}

func TestExerciseDeleteFile(t *testing.T) {
	fileSystem = make(map[string]string)
	writeFile("test.txt", "Hello")

	result := ExerciseDeleteFile("test.txt")
	if !result {
		t.Errorf("ExerciseDeleteFile() = %v, want true", result)
	}

	if fileExists("test.txt") {
		t.Errorf("File should not exist after delete")
	}
}

func TestExerciseGetFileSize(t *testing.T) {
	fileSystem = make(map[string]string)
	writeFile("test.txt", "Hello")

	result := ExerciseGetFileSize("test.txt")
	if result != 5 {
		t.Errorf("ExerciseGetFileSize() = %d, want 5", result)
	}

	result = ExerciseGetFileSize("nonexistent.txt")
	if result != -1 {
		t.Errorf("ExerciseGetFileSize() = %d, want -1", result)
	}
}

func TestExerciseAppendToFile(t *testing.T) {
	fileSystem = make(map[string]string)
	writeFile("test.txt", "Hello")

	result := ExerciseAppendToFile("test.txt", " World")
	if !result {
		t.Errorf("ExerciseAppendToFile() = %v, want true", result)
	}

	content, _ := readFile("test.txt")
	if content != "Hello World" {
		t.Errorf("File content = %q, want %q", content, "Hello World")
	}
}

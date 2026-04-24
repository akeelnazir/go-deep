package main

import (
	"testing"
)

func TestExerciseHashData(t *testing.T) {
	hash := ExerciseHashData("hello")
	if hash == "" {
		t.Errorf("ExerciseHashData() returned empty string")
	}

	hash2 := ExerciseHashData("hello")
	if hash != hash2 {
		t.Errorf("ExerciseHashData() should return same hash for same input")
	}

	hash3 := ExerciseHashData("world")
	if hash == hash3 {
		t.Errorf("ExerciseHashData() should return different hash for different input")
	}
}

func TestExerciseVerifyHash(t *testing.T) {
	data := "test data"
	hash := ExerciseHashData(data)

	result := ExerciseVerifyHash(data, hash)
	if !result {
		t.Errorf("ExerciseVerifyHash() = %v, want true", result)
	}

	result = ExerciseVerifyHash("wrong data", hash)
	if result {
		t.Errorf("ExerciseVerifyHash() = %v, want false for wrong data", result)
	}
}

func TestExerciseGenerateSecureKey(t *testing.T) {
	key := ExerciseGenerateSecureKey(32)
	if len(key) == 0 {
		t.Errorf("ExerciseGenerateSecureKey() returned empty string")
	}

	key2 := ExerciseGenerateSecureKey(32)
	if key == key2 {
		t.Errorf("ExerciseGenerateSecureKey() should generate different keys")
	}
}

func TestExerciseHashPassword(t *testing.T) {
	hash := ExerciseHashPassword("mypassword")
	if hash == "" {
		t.Errorf("ExerciseHashPassword() returned empty string")
	}

	if hash == "mypassword" {
		t.Errorf("ExerciseHashPassword() should not return plain password")
	}
}

func TestExerciseVerifyPassword(t *testing.T) {
	password := "mypassword"
	hash := ExerciseHashPassword(password)

	result := ExerciseVerifyPassword(password, hash)
	if !result {
		t.Errorf("ExerciseVerifyPassword() = %v, want true", result)
	}

	result = ExerciseVerifyPassword("wrongpassword", hash)
	if result {
		t.Errorf("ExerciseVerifyPassword() = %v, want false for wrong password", result)
	}
}

func TestExerciseComputeChecksum(t *testing.T) {
	checksum := ExerciseComputeChecksum("data")
	if checksum == "" {
		t.Errorf("ExerciseComputeChecksum() returned empty string")
	}
}

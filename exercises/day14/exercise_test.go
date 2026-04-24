package main

import (
	"testing"
)

func TestExerciseHashPassword(t *testing.T) {
	password := "testPassword"
	hash := ExerciseHashPassword(password)
	if hash == "" {
		t.Errorf("ExerciseHashPassword() returned empty string")
	}
	if hash == password {
		t.Errorf("ExerciseHashPassword() should not return plain password")
	}
}

func TestExerciseVerifyPassword(t *testing.T) {
	password := "testPassword"
	hash := ExerciseHashPassword(password)
	
	result := ExerciseVerifyPassword(hash, password)
	if !result {
		t.Errorf("ExerciseVerifyPassword() = %v, want true", result)
	}

	result = ExerciseVerifyPassword(hash, "wrongPassword")
	if result {
		t.Errorf("ExerciseVerifyPassword() = %v, want false", result)
	}
}

func TestExerciseGenerateToken(t *testing.T) {
	token := ExerciseGenerateToken(1, "alice")
	if token == "" {
		t.Errorf("ExerciseGenerateToken() returned empty string")
	}
}

func TestExerciseValidateToken(t *testing.T) {
	token := ExerciseGenerateToken(1, "alice")
	
	username := ExerciseValidateToken(token)
	if username != "alice" {
		t.Errorf("ExerciseValidateToken() = %q, want %q", username, "alice")
	}

	invalidUsername := ExerciseValidateToken("invalid_token")
	if invalidUsername != "" {
		t.Errorf("ExerciseValidateToken() = %q, want empty string", invalidUsername)
	}
}

func TestExerciseCheckRole(t *testing.T) {
	result := ExerciseCheckRole("alice", "admin")
	if !result {
		t.Errorf("ExerciseCheckRole(alice, admin) = %v, want true", result)
	}

	result = ExerciseCheckRole("bob", "admin")
	if result {
		t.Errorf("ExerciseCheckRole(bob, admin) = %v, want false", result)
	}
}

func TestExerciseAuthenticateUser(t *testing.T) {
	userID := ExerciseAuthenticateUser("alice", "hashed_password_1")
	if userID != 1 {
		t.Errorf("ExerciseAuthenticateUser(alice, hashed_password_1) = %d, want 1", userID)
	}

	userID = ExerciseAuthenticateUser("alice", "wrongPassword")
	if userID != -1 {
		t.Errorf("ExerciseAuthenticateUser(alice, wrongPassword) = %d, want -1", userID)
	}
}

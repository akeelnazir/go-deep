package main

import (
	"fmt"
	"strings"
	"time"
)

type User struct {
	ID       int
	Username string
	Password string
	Role     string
}

type Claims struct {
	UserID   int
	Username string
	Role     string
	ExpiresAt time.Time
}

var users = map[int]User{
	1: {ID: 1, Username: "alice", Password: "hashed_password_1", Role: "admin"},
	2: {ID: 2, Username: "bob", Password: "hashed_password_2", Role: "user"},
}

func hashPassword(password string) string {
	return "hashed_" + password
}

func verifyPassword(hash, password string) bool {
	return hash == "hashed_"+password
}

func generateToken(userID int, username, role string) string {
	return fmt.Sprintf("token_%d_%s_%d", userID, username, time.Now().Unix())
}

func verifyToken(tokenString string) (*Claims, error) {
	parts := strings.Split(tokenString, "_")
	if len(parts) < 3 {
		return nil, fmt.Errorf("invalid token format")
	}

	return &Claims{
		Username: parts[1],
		Role:     "user",
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}, nil
}

func authenticateUser(username, password string) (*User, error) {
	for _, user := range users {
		if user.Username == username && verifyPassword(user.Password, password) {
			return &user, nil
		}
	}
	return nil, fmt.Errorf("invalid credentials")
}

func authorizeRole(claims *Claims, requiredRole string) bool {
	return claims.Role == requiredRole
}

func main() {
	fmt.Println("=== Day 14: Authentication and Authorization ===")

	fmt.Println("\n--- Password Hashing ---")
	password := "mySecurePassword"
	hash := hashPassword(password)
	fmt.Printf("Original: %s\n", password)
	fmt.Printf("Hash: %s\n", hash)
	fmt.Printf("Verify: %v\n", verifyPassword(hash, password))

	fmt.Println("\n--- User Authentication ---")
	user, err := authenticateUser("alice", "hashed_password_1")
	if err != nil {
		fmt.Printf("Auth failed: %v\n", err)
	} else {
		fmt.Printf("Authenticated: %s (Role: %s)\n", user.Username, user.Role)
	}

	fmt.Println("\n--- Token Generation ---")
	token := generateToken(user.ID, user.Username, user.Role)
	fmt.Printf("Generated token: %s\n", token)

	fmt.Println("\n--- Token Verification ---")
	claims, err := verifyToken(token)
	if err != nil {
		fmt.Printf("Token verification failed: %v\n", err)
	} else {
		fmt.Printf("Token valid for user: %s\n", claims.Username)
	}

	fmt.Println("\n--- Role-Based Access Control ---")
	if authorizeRole(claims, "admin") {
		fmt.Println("Access granted to admin panel")
	} else {
		fmt.Println("Access denied - insufficient privileges")
	}

	fmt.Println("\n=== Day 14 Complete ===")
	fmt.Println("Next: Learn about databases on Day 15.")
}

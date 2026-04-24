package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type APIError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

var (
	users = make(map[int]User)
	mu    sync.RWMutex
	nextID = 1
)

func respondError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(APIError{
		Code:    code,
		Message: message,
	})
}

func respondJSON(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func extractID(path string) int {
	parts := strings.Split(strings.TrimPrefix(path, "/users/"), "/")
	if len(parts) > 0 {
		if id, err := strconv.Atoi(parts[0]); err == nil {
			return id
		}
	}
	return -1
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	mu.RLock()
	userList := make([]User, 0, len(users))
	for _, user := range users {
		userList = append(userList, user)
	}
	mu.RUnlock()

	respondJSON(w, http.StatusOK, userList)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := extractID(r.URL.Path)
	if id == -1 {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	mu.RLock()
	user, ok := users[id]
	mu.RUnlock()

	if !ok {
		respondError(w, http.StatusNotFound, "User not found")
		return
	}

	respondJSON(w, http.StatusOK, user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	mu.Lock()
	user.ID = nextID
	nextID++
	users[user.ID] = user
	mu.Unlock()

	respondJSON(w, http.StatusCreated, user)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := extractID(r.URL.Path)
	if id == -1 {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		respondError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	mu.Lock()
	_, ok := users[id]
	if !ok {
		mu.Unlock()
		respondError(w, http.StatusNotFound, "User not found")
		return
	}
	user.ID = id
	users[id] = user
	mu.Unlock()

	respondJSON(w, http.StatusOK, user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	id := extractID(r.URL.Path)
	if id == -1 {
		respondError(w, http.StatusBadRequest, "Invalid user ID")
		return
	}

	mu.Lock()
	delete(users, id)
	mu.Unlock()

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	fmt.Println("=== Day 13: REST API Design ===")

	mux := http.NewServeMux()
	mux.HandleFunc("/users", listUsers)
	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUser(w, r)
		case http.MethodPut:
			updateUser(w, r)
		case http.MethodDelete:
			deleteUser(w, r)
		default:
			respondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		}
	})

	fmt.Println("Server running on :8080")
	fmt.Println("Endpoints:")
	fmt.Println("  GET    /users     - List all users")
	fmt.Println("  POST   /users     - Create user")
	fmt.Println("  GET    /users/{id} - Get user")
	fmt.Println("  PUT    /users/{id} - Update user")
	fmt.Println("  DELETE /users/{id} - Delete user")
	fmt.Println("\n=== Day 13 Complete ===")
	fmt.Println("Next: Learn about authentication on Day 14.")

	http.ListenAndServe(":8080", mux)
}

package main

import (
	"fmt"
	"sync"
)

type User struct {
	ID    int
	Name  string
	Email string
}

var (
	db    = make(map[int]User)
	dbMu  sync.RWMutex
	nextID = 1
)

func queryUser(id int) (*User, error) {
	dbMu.RLock()
	defer dbMu.RUnlock()

	user, ok := db[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func queryAllUsers() ([]User, error) {
	dbMu.RLock()
	defer dbMu.RUnlock()

	users := make([]User, 0, len(db))
	for _, user := range db {
		users = append(users, user)
	}
	return users, nil
}

func insertUser(name, email string) (int, error) {
	dbMu.Lock()
	defer dbMu.Unlock()

	id := nextID
	nextID++
	db[id] = User{ID: id, Name: name, Email: email}
	return id, nil
}

func updateUser(id int, name, email string) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	if _, ok := db[id]; !ok {
		return fmt.Errorf("user not found")
	}
	db[id] = User{ID: id, Name: name, Email: email}
	return nil
}

func deleteUser(id int) error {
	dbMu.Lock()
	defer dbMu.Unlock()

	if _, ok := db[id]; !ok {
		return fmt.Errorf("user not found")
	}
	delete(db, id)
	return nil
}

func beginTransaction() {
	fmt.Println("Transaction started")
}

func commitTransaction() {
	fmt.Println("Transaction committed")
}

func rollbackTransaction() {
	fmt.Println("Transaction rolled back")
}

func main() {
	fmt.Println("=== Day 15: Databases and ORM ===")

	fmt.Println("\n--- Insert Users ---")
	id1, _ := insertUser("Alice", "alice@example.com")
	id2, _ := insertUser("Bob", "bob@example.com")
	fmt.Printf("Inserted user 1: %d\n", id1)
	fmt.Printf("Inserted user 2: %d\n", id2)

	fmt.Println("\n--- Query Single User ---")
	user, err := queryUser(id1)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("User: %+v\n", user)
	}

	fmt.Println("\n--- Query All Users ---")
	users, _ := queryAllUsers()
	for _, u := range users {
		fmt.Printf("  %+v\n", u)
	}

	fmt.Println("\n--- Update User ---")
	updateUser(id1, "Alice Updated", "alice.updated@example.com")
	user, _ = queryUser(id1)
	fmt.Printf("Updated user: %+v\n", user)

	fmt.Println("\n--- Transaction Example ---")
	beginTransaction()
	insertUser("Charlie", "charlie@example.com")
	commitTransaction()

	fmt.Println("\n--- Delete User ---")
	deleteUser(id2)
	_, err = queryUser(id2)
	if err != nil {
		fmt.Printf("User deleted: %v\n", err)
	}

	fmt.Println("\n=== Day 15 Complete ===")
	fmt.Println("Next: Learn about microservices on Day 16.")
}

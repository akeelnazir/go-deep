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

type UserService struct {
	users map[int]User
	mu    sync.RWMutex
	nextID int
}

var userService = &UserService{
	users:  make(map[int]User),
	nextID: 1,
}

func (us *UserService) CreateUser(name, email string) int {
	us.mu.Lock()
	defer us.mu.Unlock()

	id := us.nextID
	us.nextID++
	us.users[id] = User{ID: id, Name: name, Email: email}
	return id
}

func (us *UserService) GetUser(id int) *User {
	us.mu.RLock()
	defer us.mu.RUnlock()

	if user, ok := us.users[id]; ok {
		return &user
	}
	return nil
}

func (us *UserService) UpdateUser(id int, name, email string) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	if _, ok := us.users[id]; ok {
		us.users[id] = User{ID: id, Name: name, Email: email}
		return true
	}
	return false
}

func (us *UserService) DeleteUser(id int) bool {
	us.mu.Lock()
	defer us.mu.Unlock()

	if _, ok := us.users[id]; ok {
		delete(us.users, id)
		return true
	}
	return false
}

func (us *UserService) ListUsers() []User {
	us.mu.RLock()
	defer us.mu.RUnlock()

	users := make([]User, 0, len(us.users))
	for _, user := range us.users {
		users = append(users, user)
	}
	return users
}

func main() {
	fmt.Println("=== Day 26: Web Frameworks and MVC Architecture ===")

	fmt.Println("\n--- Creating Users ---")
	id1 := userService.CreateUser("Alice", "alice@example.com")
	id2 := userService.CreateUser("Bob", "bob@example.com")
	fmt.Printf("Created users: %d, %d\n", id1, id2)

	fmt.Println("\n--- Getting User ---")
	user := userService.GetUser(id1)
	if user != nil {
		fmt.Printf("User: %+v\n", user)
	}

	fmt.Println("\n--- Listing Users ---")
	users := userService.ListUsers()
	for _, u := range users {
		fmt.Printf("  %+v\n", u)
	}

	fmt.Println("\n--- Updating User ---")
	userService.UpdateUser(id1, "Alice Updated", "alice.updated@example.com")
	user = userService.GetUser(id1)
	fmt.Printf("Updated user: %+v\n", user)

	fmt.Println("\n--- Deleting User ---")
	userService.DeleteUser(id2)
	fmt.Printf("Users after delete: %d\n", len(userService.ListUsers()))

	fmt.Println("\n=== Day 26 Complete ===")
	fmt.Println("Next: Learn about embedding and code generation on Day 27.")
}

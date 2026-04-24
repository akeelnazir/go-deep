package main

import (
	"fmt"
	"sync"
	"time"
)

type Connection struct {
	ID        int
	Address   string
	Connected bool
	CreatedAt time.Time
}

type ConnectionPool struct {
	connections map[int]*Connection
	mu          sync.RWMutex
	nextID      int
}

var pool = &ConnectionPool{
	connections: make(map[int]*Connection),
	nextID:      1,
}

func (cp *ConnectionPool) CreateConnection(address string) int {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	id := cp.nextID
	cp.nextID++

	cp.connections[id] = &Connection{
		ID:        id,
		Address:   address,
		Connected: true,
		CreatedAt: time.Now(),
	}

	return id
}

func (cp *ConnectionPool) CloseConnection(id int) error {
	cp.mu.Lock()
	defer cp.mu.Unlock()

	if conn, ok := cp.connections[id]; ok {
		conn.Connected = false
		return nil
	}
	return fmt.Errorf("connection not found")
}

func (cp *ConnectionPool) GetConnection(id int) *Connection {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	return cp.connections[id]
}

func (cp *ConnectionPool) ListConnections() []*Connection {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	conns := make([]*Connection, 0, len(cp.connections))
	for _, c := range cp.connections {
		conns = append(conns, c)
	}
	return conns
}

func (cp *ConnectionPool) ActiveConnections() int {
	cp.mu.RLock()
	defer cp.mu.RUnlock()

	count := 0
	for _, c := range cp.connections {
		if c.Connected {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println("=== Day 22: Networking Fundamentals ===")

	fmt.Println("\n--- Creating Connections ---")
	id1 := pool.CreateConnection("localhost:8080")
	id2 := pool.CreateConnection("localhost:8081")
	id3 := pool.CreateConnection("localhost:8082")
	fmt.Printf("Created connections: %d, %d, %d\n", id1, id2, id3)

	fmt.Println("\n--- Listing Connections ---")
	for _, conn := range pool.ListConnections() {
		fmt.Printf("ID: %d, Address: %s, Connected: %v\n", conn.ID, conn.Address, conn.Connected)
	}

	fmt.Println("\n--- Active Connections ---")
	fmt.Printf("Active connections: %d\n", pool.ActiveConnections())

	fmt.Println("\n--- Closing Connection ---")
	pool.CloseConnection(id1)
	fmt.Printf("Active connections after close: %d\n", pool.ActiveConnections())

	fmt.Println("\n--- Connection Details ---")
	conn := pool.GetConnection(id2)
	if conn != nil {
		uptime := time.Since(conn.CreatedAt)
		fmt.Printf("Connection %d uptime: %v\n", id2, uptime)
	}

	fmt.Println("\n=== Day 22 Complete ===")
	fmt.Println("Next: Learn about parsing and text processing on Day 23.")
}

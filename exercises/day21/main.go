package main

import (
	"fmt"
	"sync"
	"time"
)

type ProcessManager struct {
	processes map[int]*ProcessInfo
	mu        sync.RWMutex
	nextPID   int
}

type ProcessInfo struct {
	PID       int
	Name      string
	StartTime time.Time
	Status    string
}

var pm = &ProcessManager{
	processes: make(map[int]*ProcessInfo),
	nextPID:   1000,
}

func (pm *ProcessManager) StartProcess(name string) int {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pid := pm.nextPID
	pm.nextPID++

	pm.processes[pid] = &ProcessInfo{
		PID:       pid,
		Name:      name,
		StartTime: time.Now(),
		Status:    "running",
	}

	return pid
}

func (pm *ProcessManager) StopProcess(pid int) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	if proc, ok := pm.processes[pid]; ok {
		proc.Status = "stopped"
		return nil
	}
	return fmt.Errorf("process not found")
}

func (pm *ProcessManager) GetProcess(pid int) *ProcessInfo {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	return pm.processes[pid]
}

func (pm *ProcessManager) ListProcesses() []*ProcessInfo {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	procs := make([]*ProcessInfo, 0, len(pm.processes))
	for _, p := range pm.processes {
		procs = append(procs, p)
	}
	return procs
}

func main() {
	fmt.Println("=== Day 21: Signal Handling and Processes ===")

	fmt.Println("\n--- Starting Processes ---")
	pid1 := pm.StartProcess("server")
	pid2 := pm.StartProcess("worker")
	fmt.Printf("Started processes: %d, %d\n", pid1, pid2)

	fmt.Println("\n--- Listing Processes ---")
	for _, proc := range pm.ListProcesses() {
		fmt.Printf("PID: %d, Name: %s, Status: %s\n", proc.PID, proc.Name, proc.Status)
	}

	fmt.Println("\n--- Stopping Process ---")
	pm.StopProcess(pid1)
	proc := pm.GetProcess(pid1)
	fmt.Printf("Process %d status: %s\n", pid1, proc.Status)

	fmt.Println("\n--- Process Uptime ---")
	proc = pm.GetProcess(pid2)
	uptime := time.Since(proc.StartTime)
	fmt.Printf("Process %d uptime: %v\n", pid2, uptime)

	fmt.Println("\n=== Day 21 Complete ===")
	fmt.Println("Next: Learn about networking on Day 22.")
}

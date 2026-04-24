package process

import (
	"fmt"
	"sync"
	"time"
)

type ProcessState int

const (
	Ready ProcessState = iota
	Running
	Blocked
	Terminated
)

type Process struct {
	PID          uint32
	Name         string
	State        ProcessState
	Priority     int
	CreatedAt    time.Time
	StartedAt    time.Time
	TerminatedAt time.Time
	CPUTime      time.Duration
	MemoryUsed   uint64
	Duration     time.Duration
	Mu           sync.RWMutex
}

func NewProcess(pid uint32, name string, duration time.Duration) *Process {
	return &Process{
		PID:       pid,
		Name:      name,
		State:     Ready,
		Priority:  5,
		CreatedAt: time.Now(),
		Duration:  duration,
	}
}

func (p *Process) Start() {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.State = Running
	p.StartedAt = time.Now()
}

func (p *Process) Terminate() {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.State = Terminated
	p.TerminatedAt = time.Now()
}

func (p *Process) AddCPUTime(duration time.Duration) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.CPUTime += duration
}

func (p *Process) GetState() ProcessState {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.State
}

func (p *Process) GetCPUTime() time.Duration {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return p.CPUTime
}

func (p *Process) String() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()
	return fmt.Sprintf("PID: %d, Name: %s, State: %s, CPU: %v, Memory: %d KB",
		p.PID, p.Name, stateString(p.State), p.CPUTime, p.MemoryUsed/1024)
}

func stateString(state ProcessState) string {
	switch state {
	case Ready:
		return "Ready"
	case Running:
		return "Running"
	case Blocked:
		return "Blocked"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

type ProcessManager struct {
	processes map[uint32]*Process
	nextPID   uint32
	mu        sync.RWMutex
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[uint32]*Process),
		nextPID:   1,
	}
}

func (pm *ProcessManager) CreateProcess(name string, duration time.Duration) *Process {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	pid := pm.nextPID
	pm.nextPID++

	p := NewProcess(pid, name, duration)
	pm.processes[pid] = p
	return p
}

func (pm *ProcessManager) GetProcess(pid uint32) *Process {
	pm.mu.RLock()
	defer pm.mu.RUnlock()
	return pm.processes[pid]
}

func (pm *ProcessManager) TerminateProcess(pid uint32) error {
	pm.mu.Lock()
	defer pm.mu.Unlock()

	p, exists := pm.processes[pid]
	if !exists {
		return fmt.Errorf("process %d not found", pid)
	}

	p.Terminate()
	return nil
}

func (pm *ProcessManager) GetAllProcesses() []*Process {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	processes := make([]*Process, 0, len(pm.processes))
	for _, p := range pm.processes {
		processes = append(processes, p)
	}
	return processes
}

func (pm *ProcessManager) GetReadyProcesses() []*Process {
	pm.mu.RLock()
	defer pm.mu.RUnlock()

	var ready []*Process
	for _, p := range pm.processes {
		p.Mu.RLock()
		if p.State == Ready {
			ready = append(ready, p)
		}
		p.Mu.RUnlock()
	}
	return ready
}

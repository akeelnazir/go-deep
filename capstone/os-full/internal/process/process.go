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
	Waiting
	Terminated
)

type Process struct {
	PID           uint32
	PPID          uint32
	Name          string
	State         ProcessState
	Priority      int
	CreatedAt     time.Time
	StartedAt     time.Time
	TerminatedAt  time.Time
	CPUTime       time.Duration
	MemoryUsed    uint64
	MemoryLimit   uint64
	Duration      time.Duration
	Owner         string
	ExitCode      int
	Mu            sync.RWMutex
}

func NewProcess(pid uint32, ppid uint32, name string, priority int, duration time.Duration) *Process {
	return &Process{
		PID:         pid,
		PPID:        ppid,
		Name:        name,
		State:       Ready,
		Priority:    priority,
		CreatedAt:   time.Now(),
		Duration:    duration,
		MemoryLimit: 10 * 1024 * 1024,
		Owner:       "root",
	}
}

func (p *Process) Start() {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.State = Running
	p.StartedAt = time.Now()
}

func (p *Process) Terminate(exitCode int) {
	p.Mu.Lock()
	defer p.Mu.Unlock()
	p.State = Terminated
	p.TerminatedAt = time.Now()
	p.ExitCode = exitCode
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
	return fmt.Sprintf("PID: %d, Name: %s, State: %s, Priority: %d, CPU: %v, Memory: %d KB",
		p.PID, p.Name, stateString(p.State), p.Priority, p.CPUTime, p.MemoryUsed/1024)
}

func stateString(state ProcessState) string {
	switch state {
	case Ready:
		return "Ready"
	case Running:
		return "Running"
	case Blocked:
		return "Blocked"
	case Waiting:
		return "Waiting"
	case Terminated:
		return "Terminated"
	default:
		return "Unknown"
	}
}

type ProcessManager struct {
	processes map[uint32]*Process
	nextPID   uint32
	Mu        sync.RWMutex
}

func NewProcessManager() *ProcessManager {
	return &ProcessManager{
		processes: make(map[uint32]*Process),
		nextPID:   1,
	}
}

func (pm *ProcessManager) CreateProcess(name string, priority int, duration time.Duration) *Process {
	pm.Mu.Lock()
	defer pm.Mu.Unlock()

	pid := pm.nextPID
	pm.nextPID++

	p := NewProcess(pid, 0, name, priority, duration)
	pm.processes[pid] = p
	return p
}

func (pm *ProcessManager) GetProcess(pid uint32) *Process {
	pm.Mu.RLock()
	defer pm.Mu.RUnlock()
	return pm.processes[pid]
}

func (pm *ProcessManager) TerminateProcess(pid uint32) error {
	pm.Mu.Lock()
	defer pm.Mu.Unlock()

	p, exists := pm.processes[pid]
	if !exists {
		return fmt.Errorf("process %d not found", pid)
	}

	p.Terminate(0)
	return nil
}

func (pm *ProcessManager) GetAllProcesses() []*Process {
	pm.Mu.RLock()
	defer pm.Mu.RUnlock()

	processes := make([]*Process, 0, len(pm.processes))
	for _, p := range pm.processes {
		processes = append(processes, p)
	}
	return processes
}

func (pm *ProcessManager) GetProcessesByPriority(priority int) []*Process {
	pm.Mu.RLock()
	defer pm.Mu.RUnlock()

	var procs []*Process
	for _, p := range pm.processes {
		p.Mu.RLock()
		if p.Priority == priority && p.State == Ready {
			procs = append(procs, p)
		}
		p.Mu.RUnlock()
	}
	return procs
}

func (pm *ProcessManager) ChangePriority(pid uint32, newPriority int) error {
	pm.Mu.RLock()
	defer pm.Mu.RUnlock()

	p, exists := pm.processes[pid]
	if !exists {
		return fmt.Errorf("process %d not found", pid)
	}

	p.Mu.Lock()
	p.Priority = newPriority
	p.Mu.Unlock()

	return nil
}

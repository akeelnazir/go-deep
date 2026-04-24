package process

import (
	"fmt"
	"sync"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/cpu"
)

type ProcessState int

const (
	New ProcessState = iota
	Ready
	Running
	Blocked
	Terminated
)

type Process struct {
	PID        uint32
	Name       string
	State      ProcessState
	CPU        *cpu.CPU
	MemoryBase uint32
	MemorySize uint32
	CreatedAt  time.Time
	StartedAt  time.Time
	TerminatedAt time.Time
	CPUTime    time.Duration
	ExitCode   int
	Mu         sync.RWMutex
}

func NewProcess(pid uint32, name string) *Process {
	return &Process{
		PID:       pid,
		Name:      name,
		State:     New,
		CPU:       cpu.NewCPU(),
		CreatedAt: time.Now(),
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

func (p *Process) String() string {
	p.Mu.RLock()
	defer p.Mu.RUnlock()

	stateStr := "Unknown"
	switch p.State {
	case New:
		stateStr = "New"
	case Ready:
		stateStr = "Ready"
	case Running:
		stateStr = "Running"
	case Blocked:
		stateStr = "Blocked"
	case Terminated:
		stateStr = "Terminated"
	}

	return fmt.Sprintf("PID: %d, Name: %s, State: %s, CPU Time: %v, Memory: %d KB",
		p.PID, p.Name, stateStr, p.CPUTime, p.MemorySize/1024)
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

func (pm *ProcessManager) CreateProcess(name string) *Process {
	pm.Mu.Lock()
	defer pm.Mu.Unlock()

	pid := pm.nextPID
	pm.nextPID++

	p := NewProcess(pid, name)
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

func (pm *ProcessManager) GetReadyProcesses() []*Process {
	pm.Mu.RLock()
	defer pm.Mu.RUnlock()

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

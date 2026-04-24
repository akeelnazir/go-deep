package scheduler

import (
	"sync"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-process/internal/process"
)

const TimeSlice = 100 * time.Millisecond

type Scheduler struct {
	pm          *process.ProcessManager
	readyQueue  []*process.Process
	currentProc *process.Process
	mu          sync.RWMutex
	running     bool
	tickChan    chan struct{}
}

func NewScheduler(pm *process.ProcessManager) *Scheduler {
	return &Scheduler{
		pm:         pm,
		readyQueue: make([]*process.Process, 0),
		tickChan:   make(chan struct{}, 1),
		running:    false,
	}
}

func (s *Scheduler) Start() {
	s.mu.Lock()
	s.running = true
	s.mu.Unlock()

	ticker := time.NewTicker(TimeSlice)
	defer ticker.Stop()

	for {
		s.mu.RLock()
		if !s.running {
			s.mu.RUnlock()
			break
		}
		s.mu.RUnlock()

		select {
		case <-ticker.C:
			s.schedule()
		}
	}
}

func (s *Scheduler) Stop() {
	s.mu.Lock()
	s.running = false
	s.mu.Unlock()
}

func (s *Scheduler) schedule() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.currentProc != nil {
		s.currentProc.Mu.Lock()
		if s.currentProc.State == process.Running {
			s.currentProc.State = process.Ready
			s.currentProc.AddCPUTime(TimeSlice)
		}
		s.currentProc.Mu.Unlock()
	}

	allProcs := s.pm.GetAllProcesses()
	var nextProc *process.Process

	for _, p := range allProcs {
		p.Mu.RLock()
		if p.State == process.Ready {
			nextProc = p
			p.Mu.RUnlock()
			break
		}
		p.Mu.RUnlock()
	}

	if nextProc != nil {
		nextProc.Mu.Lock()
		if nextProc.State == process.Ready {
			nextProc.State = process.Running
			if nextProc.StartedAt.IsZero() {
				nextProc.StartedAt = time.Now()
			}

			if nextProc.Duration > 0 && nextProc.CPUTime >= nextProc.Duration {
				nextProc.State = process.Terminated
				nextProc.TerminatedAt = time.Now()
			}
		}
		nextProc.Mu.Unlock()
		s.currentProc = nextProc
	} else {
		s.currentProc = nil
	}
}

func (s *Scheduler) GetCurrentProcess() *process.Process {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.currentProc
}

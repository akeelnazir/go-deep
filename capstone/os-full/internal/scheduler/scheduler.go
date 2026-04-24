package scheduler

import (
	"sync"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-full/internal/process"
)

const TimeSlice = 100 * time.Millisecond

type Scheduler struct {
	pm          *process.ProcessManager
	currentProc *process.Process
	Mu          sync.RWMutex
	running     bool
}

func NewScheduler(pm *process.ProcessManager) *Scheduler {
	return &Scheduler{
		pm:      pm,
		running: false,
	}
}

func (s *Scheduler) Start() {
	s.Mu.Lock()
	s.running = true
	s.Mu.Unlock()

	ticker := time.NewTicker(TimeSlice)
	defer ticker.Stop()

	for {
		s.Mu.RLock()
		if !s.running {
			s.Mu.RUnlock()
			break
		}
		s.Mu.RUnlock()

		select {
		case <-ticker.C:
			s.schedule()
		}
	}
}

func (s *Scheduler) Stop() {
	s.Mu.Lock()
	s.running = false
	s.Mu.Unlock()
}

func (s *Scheduler) schedule() {
	s.Mu.Lock()
	defer s.Mu.Unlock()

	if s.currentProc != nil {
		s.currentProc.Mu.Lock()
		if s.currentProc.State == process.Running {
			s.currentProc.State = process.Ready
			s.currentProc.AddCPUTime(TimeSlice)

			if s.currentProc.Duration > 0 && s.currentProc.CPUTime >= s.currentProc.Duration {
				s.currentProc.State = process.Terminated
				s.currentProc.TerminatedAt = time.Now()
			}
		}
		s.currentProc.Mu.Unlock()
	}

	var nextProc *process.Process
	allProcs := s.pm.GetAllProcesses()

	for priority := 0; priority < 16; priority++ {
		for _, p := range allProcs {
			p.Mu.RLock()
			if p.Priority == priority && p.State == process.Ready {
				nextProc = p
				p.Mu.RUnlock()
				break
			}
			p.Mu.RUnlock()
		}
		if nextProc != nil {
			break
		}
	}

	if nextProc != nil {
		nextProc.Mu.Lock()
		if nextProc.State == process.Ready {
			nextProc.State = process.Running
			if nextProc.StartedAt.IsZero() {
				nextProc.StartedAt = time.Now()
			}
		}
		nextProc.Mu.Unlock()
		s.currentProc = nextProc
	} else {
		s.currentProc = nil
	}
}

func (s *Scheduler) GetCurrentProcess() *process.Process {
	s.Mu.RLock()
	defer s.Mu.RUnlock()
	return s.currentProc
}

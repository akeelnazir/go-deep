package main

import (
	"testing"
	"time"
)

func TestExerciseStartProcess(t *testing.T) {
	pm.processes = make(map[int]*ProcessInfo)
	pm.nextPID = 1000

	pid := ExerciseStartProcess("test")
	if pid == 0 {
		t.Errorf("ExerciseStartProcess() = 0, want non-zero PID")
	}

	if !ExerciseProcessExists(pid) {
		t.Errorf("Process should exist after start")
	}
}

func TestExerciseStopProcess(t *testing.T) {
	pm.processes = make(map[int]*ProcessInfo)
	pm.nextPID = 1000

	pid := ExerciseStartProcess("test")
	result := ExerciseStopProcess(pid)
	if !result {
		t.Errorf("ExerciseStopProcess() = %v, want true", result)
	}

	status := ExerciseGetProcessStatus(pid)
	if status != "stopped" {
		t.Errorf("Process status = %q, want %q", status, "stopped")
	}
}

func TestExerciseGetProcessStatus(t *testing.T) {
	pm.processes = make(map[int]*ProcessInfo)
	pm.nextPID = 1000

	pid := ExerciseStartProcess("test")
	status := ExerciseGetProcessStatus(pid)
	if status != "running" {
		t.Errorf("ExerciseGetProcessStatus() = %q, want %q", status, "running")
	}
}

func TestExerciseListAllProcesses(t *testing.T) {
	pm.processes = make(map[int]*ProcessInfo)
	pm.nextPID = 1000

	count := ExerciseListAllProcesses()
	if count != 0 {
		t.Errorf("ExerciseListAllProcesses() = %d, want 0", count)
	}

	ExerciseStartProcess("test1")
	ExerciseStartProcess("test2")

	count = ExerciseListAllProcesses()
	if count != 2 {
		t.Errorf("ExerciseListAllProcesses() = %d, want 2", count)
	}
}

func TestExerciseGetProcessUptime(t *testing.T) {
	pm.processes = make(map[int]*ProcessInfo)
	pm.nextPID = 1000

	pid := ExerciseStartProcess("test")
	time.Sleep(100 * time.Millisecond)

	uptime := ExerciseGetProcessUptime(pid)
	if uptime < 0 {
		t.Errorf("ExerciseGetProcessUptime() = %d, want >= 0", uptime)
	}
}

func TestExerciseProcessExists(t *testing.T) {
	pm.processes = make(map[int]*ProcessInfo)
	pm.nextPID = 1000

	pid := ExerciseStartProcess("test")

	if !ExerciseProcessExists(pid) {
		t.Errorf("ExerciseProcessExists() = false, want true")
	}

	if ExerciseProcessExists(9999) {
		t.Errorf("ExerciseProcessExists() = true, want false for nonexistent")
	}
}

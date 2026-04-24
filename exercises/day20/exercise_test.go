package main

import (
	"testing"
)

func TestExerciseRegisterCommand(t *testing.T) {
	commands = make(map[string]*Command)

	ExerciseRegisterCommand("test", "Test command", func(args []string) error {
		return nil
	})

	if !ExerciseCommandExists("test") {
		t.Errorf("Command should be registered")
	}
}

func TestExerciseExecuteCommand(t *testing.T) {
	commands = make(map[string]*Command)

	ExerciseRegisterCommand("greet", "Greet", func(args []string) error {
		return nil
	})

	err := ExerciseExecuteCommand("greet", []string{})
	if err != nil {
		t.Errorf("ExerciseExecuteCommand() = %v, want nil", err)
	}

	err = ExerciseExecuteCommand("nonexistent", []string{})
	if err == nil {
		t.Errorf("ExerciseExecuteCommand() = nil, want error for nonexistent command")
	}
}

func TestExerciseGetCommand(t *testing.T) {
	commands = make(map[string]*Command)

	ExerciseRegisterCommand("test", "Test", func(args []string) error {
		return nil
	})

	cmd := ExerciseGetCommand("test")
	if cmd == nil {
		t.Errorf("ExerciseGetCommand() = nil, want command")
	}

	cmd = ExerciseGetCommand("nonexistent")
	if cmd != nil {
		t.Errorf("ExerciseGetCommand() = %v, want nil", cmd)
	}
}

func TestExerciseListCommands(t *testing.T) {
	commands = make(map[string]*Command)

	ExerciseRegisterCommand("cmd1", "Command 1", func(args []string) error { return nil })
	ExerciseRegisterCommand("cmd2", "Command 2", func(args []string) error { return nil })

	list := ExerciseListCommands()
	if len(list) != 2 {
		t.Errorf("ExerciseListCommands() = %d commands, want 2", len(list))
	}
}

func TestExerciseCommandExists(t *testing.T) {
	commands = make(map[string]*Command)

	ExerciseRegisterCommand("test", "Test", func(args []string) error { return nil })

	if !ExerciseCommandExists("test") {
		t.Errorf("ExerciseCommandExists() = false, want true")
	}

	if ExerciseCommandExists("nonexistent") {
		t.Errorf("ExerciseCommandExists() = true, want false")
	}
}

func TestExerciseRemoveCommand(t *testing.T) {
	commands = make(map[string]*Command)

	ExerciseRegisterCommand("test", "Test", func(args []string) error { return nil })

	result := ExerciseRemoveCommand("test")
	if !result {
		t.Errorf("ExerciseRemoveCommand() = %v, want true", result)
	}

	if ExerciseCommandExists("test") {
		t.Errorf("Command should be removed")
	}

	result = ExerciseRemoveCommand("nonexistent")
	if result {
		t.Errorf("ExerciseRemoveCommand() = %v, want false for nonexistent", result)
	}
}

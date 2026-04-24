package main

import (
	"fmt"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/filesystem"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/memory"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/process"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/scheduler"
	"github.com/akeelnazir/go-deep/capstone/os-emulator/internal/shell"
)

func main() {
	fmt.Println("Starting OS Emulator...")
	fmt.Println()

	pm := process.NewProcessManager()
	mm := memory.NewMemoryManager()
	fs := filesystem.NewFileSystem()
	fs.Initialize()

	sched := scheduler.NewScheduler(pm)

	go func() {
		sched.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	sh := shell.NewShell(pm, mm, fs, sched)
	sh.Start()

	sched.Stop()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("OS Emulator shutdown complete.")
}

package main

import (
	"fmt"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-process/internal/filesystem"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/memory"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/process"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/scheduler"
	"github.com/akeelnazir/go-deep/capstone/os-process/internal/shell"
)

func main() {
	fmt.Println("Starting OS Process...")
	fmt.Println()

	pm := process.NewProcessManager()
	mm := memory.NewMemoryManager(64 * 1024 * 1024)
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

	fmt.Println("OS Process shutdown complete.")
}

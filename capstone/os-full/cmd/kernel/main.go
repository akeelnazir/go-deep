package main

import (
	"fmt"
	"time"

	"github.com/akeelnazir/go-deep/capstone/os-full/internal/device"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/filesystem"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/ipc"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/memory"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/process"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/scheduler"
	"github.com/akeelnazir/go-deep/capstone/os-full/internal/shell"
)

func main() {
	fmt.Println("Starting OS Full...")
	fmt.Println()

	pm := process.NewProcessManager()
	mm := memory.NewMemoryManager(256 * 1024 * 1024)
	fs := filesystem.NewFileSystem()
	fs.Initialize()
	dm := device.NewDeviceManager()
	im := ipc.NewIPCManager()

	dm.RegisterDevice("disk0", device.TypeDisk, 8, 0)
	dm.RegisterDevice("console", device.TypeConsole, 5, 1)
	dm.RegisterDevice("eth0", device.TypeNetwork, 10, 0)

	sched := scheduler.NewScheduler(pm)

	go func() {
		sched.Start()
	}()

	time.Sleep(100 * time.Millisecond)

	sh := shell.NewShell(pm, mm, fs, dm, im, sched)
	sh.Start()

	sched.Stop()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("OS Full shutdown complete.")
}

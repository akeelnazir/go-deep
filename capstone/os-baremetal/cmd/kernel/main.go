package main

import (
	"fmt"
	"os"

	"github.com/akeelnazir/go-deep/capstone/os-baremetal/internal/bootloader"
	"github.com/akeelnazir/go-deep/capstone/os-baremetal/internal/kernel"
	"github.com/akeelnazir/go-deep/capstone/os-baremetal/internal/shell"
)

func main() {
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     OS Baremetal - Real Bootable OS    ║")
	fmt.Println("║   Running independently on bare metal  ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	bl := bootloader.NewBootloader()
	if err := bl.Initialize(); err != nil {
		fmt.Fprintf(os.Stderr, "Bootloader error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()

	k := kernel.NewKernel(bl.GetMemorySize(), bl.GetCPUSpeed())
	if err := k.Initialize(); err != nil {
		fmt.Fprintf(os.Stderr, "Kernel error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()

	sh := shell.NewShell(k)
	sh.Start()

	fmt.Println()
	fmt.Println("OS Baremetal shutdown complete.")
}

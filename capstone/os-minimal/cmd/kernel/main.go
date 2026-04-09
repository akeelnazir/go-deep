package main

import (
	"fmt"
	"os"

	"github.com/akeelnazir/go-deep/capstone/os-minimal/internal/bootloader"
	"github.com/akeelnazir/go-deep/capstone/os-minimal/internal/kernel"
	"github.com/akeelnazir/go-deep/capstone/os-minimal/internal/shell"
)

func main() {
	fmt.Println("Starting OS Minimal...")
	fmt.Println()

	bl := bootloader.NewBootloader()
	if err := bl.Initialize(); err != nil {
		fmt.Fprintf(os.Stderr, "Bootloader error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()

	k := kernel.NewKernel(bl.GetMemorySize())
	if err := k.Initialize(); err != nil {
		fmt.Fprintf(os.Stderr, "Kernel error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println()

	sh := shell.NewShell(k)
	sh.Start()

	fmt.Println("OS Minimal shutdown complete.")
}

package bootloader

import "fmt"

type Bootloader struct {
	memorySize uint64
	cpuSpeed   uint32
}

func NewBootloader() *Bootloader {
	return &Bootloader{
		memorySize: 64 * 1024 * 1024,
		cpuSpeed:   3000,
	}
}

func (b *Bootloader) Initialize() error {
	fmt.Println("=== Bootloader Initialization ===")
	fmt.Println("Initializing CPU...")
	fmt.Println("Checking memory...")
	fmt.Printf("Memory available: %d MB\n", b.memorySize/(1024*1024))
	fmt.Printf("CPU speed: %d MHz\n", b.cpuSpeed)
	fmt.Println("Loading kernel...")
	fmt.Println("Bootloader complete. Jumping to kernel entry point.")
	return nil
}

func (b *Bootloader) GetMemorySize() uint64 {
	return b.memorySize
}

func (b *Bootloader) GetCPUSpeed() uint32 {
	return b.cpuSpeed
}

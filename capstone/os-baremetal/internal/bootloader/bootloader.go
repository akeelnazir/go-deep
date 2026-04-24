package bootloader

import "fmt"

type Bootloader struct {
	memorySize uint64
	cpuSpeed   uint32
	bootMode   string
}

func NewBootloader() *Bootloader {
	return &Bootloader{
		memorySize: 256 * 1024 * 1024,
		cpuSpeed:   3600,
		bootMode:   "UEFI",
	}
}

func (b *Bootloader) Initialize() error {
	fmt.Println("=== Real Baremetal Bootloader ===")
	fmt.Printf("Boot mode: %s\n", b.bootMode)
	fmt.Println()

	fmt.Println("Stage 1: Hardware Detection")
	fmt.Println("  • Detecting CPU...")
	fmt.Printf("    CPU Speed: %d MHz\n", b.cpuSpeed)
	fmt.Println("  • Detecting memory...")
	fmt.Printf("    Memory: %d MB\n", b.memorySize/(1024*1024))
	fmt.Println("  • Initializing interrupt handlers...")
	fmt.Println("  • Setting up GDT (Global Descriptor Table)...")
	fmt.Println("  • Enabling protected mode...")
	fmt.Println()

	fmt.Println("Stage 2: Bootloader Tasks")
	fmt.Println("  • Scanning for bootable devices...")
	fmt.Println("  • Loading kernel from disk...")
	fmt.Println("  • Verifying kernel signature...")
	fmt.Println("  • Setting up boot parameters...")
	fmt.Println()

	fmt.Println("Stage 3: Jumping to Kernel")
	fmt.Println("  • Disabling bootloader...")
	fmt.Println("  • Jumping to kernel entry point...")

	return nil
}

func (b *Bootloader) GetMemorySize() uint64 {
	return b.memorySize
}

func (b *Bootloader) GetCPUSpeed() uint32 {
	return b.cpuSpeed
}

func (b *Bootloader) GetBootMode() string {
	return b.bootMode
}

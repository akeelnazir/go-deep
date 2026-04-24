package kernel

import (
	"fmt"
	"time"
)

type Kernel struct {
	bootTime      time.Time
	interruptMap  map[int]InterruptHandler
	memoryManager MemoryManager
	processID     uint32
	cpuSpeed      uint32
}

type InterruptHandler func(code int) error

type MemoryManager struct {
	totalMemory uint64
	usedMemory  uint64
	heapStart   uint64
	heapEnd     uint64
}

func NewKernel(totalMemory uint64, cpuSpeed uint32) *Kernel {
	return &Kernel{
		bootTime:     time.Now(),
		interruptMap: make(map[int]InterruptHandler),
		memoryManager: MemoryManager{
			totalMemory: totalMemory,
			usedMemory:  0,
			heapStart:   0x100000,
			heapEnd:     totalMemory,
		},
		processID: 1,
		cpuSpeed:  cpuSpeed,
	}
}

func (k *Kernel) Initialize() error {
	fmt.Println("=== Baremetal Kernel Initialization ===")
	fmt.Println()

	fmt.Println("Phase 1: CPU Setup")
	fmt.Println("  • Initializing CPU features...")
	fmt.Printf("    CPU Speed: %d MHz\n", k.cpuSpeed)
	fmt.Println("  • Setting up control registers...")
	fmt.Println("  • Enabling caching...")
	fmt.Println()

	fmt.Println("Phase 2: Memory Management")
	fmt.Println("  • Setting up paging...")
	fmt.Printf("    Total Memory: %d MB\n", k.memoryManager.totalMemory/(1024*1024))
	fmt.Printf("    Heap Start: 0x%x\n", k.memoryManager.heapStart)
	fmt.Println("  • Initializing memory allocator...")
	fmt.Println("  • Setting up page tables...")
	fmt.Println()

	fmt.Println("Phase 3: Interrupt Handling")
	fmt.Println("  • Setting up interrupt vector table...")
	k.setupInterruptHandlers()
	fmt.Println("  • Registering exception handlers...")
	fmt.Println("  • Enabling interrupts...")
	fmt.Println()

	fmt.Println("Phase 4: Device Initialization")
	fmt.Println("  • Initializing timer...")
	fmt.Println("  • Initializing console...")
	fmt.Println("  • Initializing disk controller...")
	fmt.Println()

	fmt.Println("Kernel initialization complete.")
	return nil
}

func (k *Kernel) setupInterruptHandlers() {
	k.interruptMap[0] = k.handleDivisionByZero
	k.interruptMap[6] = k.handleInvalidOpcode
	k.interruptMap[14] = k.handlePageFault
	k.interruptMap[80] = k.handleSystemCall
}

func (k *Kernel) handleDivisionByZero(code int) error {
	fmt.Println("EXCEPTION: Division by zero (INT 0)")
	return fmt.Errorf("division by zero exception")
}

func (k *Kernel) handleInvalidOpcode(code int) error {
	fmt.Println("EXCEPTION: Invalid opcode (INT 6)")
	return fmt.Errorf("invalid opcode exception")
}

func (k *Kernel) handlePageFault(code int) error {
	fmt.Println("EXCEPTION: Page fault (INT 14)")
	return fmt.Errorf("page fault exception")
}

func (k *Kernel) handleSystemCall(code int) error {
	return nil
}

func (k *Kernel) RaiseInterrupt(code int) error {
	if handler, exists := k.interruptMap[code]; exists {
		return handler(code)
	}
	return fmt.Errorf("unknown interrupt: %d", code)
}

func (k *Kernel) AllocateMemory(size uint64) (uint64, error) {
	if k.memoryManager.usedMemory+size > k.memoryManager.heapEnd {
		return 0, fmt.Errorf("out of memory")
	}
	addr := k.memoryManager.heapStart + k.memoryManager.usedMemory
	k.memoryManager.usedMemory += size
	return addr, nil
}

func (k *Kernel) GetMemoryStats() (total, used uint64) {
	return k.memoryManager.totalMemory, k.memoryManager.usedMemory
}

func (k *Kernel) GetUptime() time.Duration {
	return time.Since(k.bootTime)
}

func (k *Kernel) CreateProcess() uint32 {
	k.processID++
	return k.processID
}

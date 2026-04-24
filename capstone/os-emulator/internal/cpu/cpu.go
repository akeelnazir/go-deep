package cpu

import (
	"fmt"
	"sync"
	"time"
)

type Register uint32

type CPU struct {
	Registers map[string]Register
	PC        uint32
	SP        uint32
	Flags     uint32
	Cycles    uint64
	StartTime time.Time
	Mu        sync.RWMutex
}

func NewCPU() *CPU {
	return &CPU{
		Registers: map[string]Register{
			"EAX": 0,
			"EBX": 0,
			"ECX": 0,
			"EDX": 0,
			"ESP": 0xFFFF,
			"EBP": 0xFFFF,
			"ESI": 0,
			"EDI": 0,
		},
		PC:        0,
		SP:        0xFFFF,
		Flags:     0,
		Cycles:    0,
		StartTime: time.Now(),
	}
}

func (c *CPU) GetRegister(name string) (Register, error) {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	val, exists := c.Registers[name]
	if !exists {
		return 0, fmt.Errorf("unknown register: %s", name)
	}
	return val, nil
}

func (c *CPU) SetRegister(name string, value Register) error {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	if _, exists := c.Registers[name]; !exists {
		return fmt.Errorf("unknown register: %s", name)
	}
	c.Registers[name] = value
	return nil
}

func (c *CPU) IncrementPC(amount uint32) {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.PC += amount
}

func (c *CPU) IncrementCycles() {
	c.Mu.Lock()
	defer c.Mu.Unlock()
	c.Cycles++
}

func (c *CPU) GetCycles() uint64 {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	return c.Cycles
}

func (c *CPU) GetUptime() time.Duration {
	c.Mu.RLock()
	defer c.Mu.RUnlock()
	return time.Since(c.StartTime)
}

func (c *CPU) Reset() {
	c.Mu.Lock()
	defer c.Mu.Unlock()

	for key := range c.Registers {
		c.Registers[key] = 0
	}
	c.PC = 0
	c.SP = 0xFFFF
	c.Flags = 0
	c.Cycles = 0
}

func (c *CPU) String() string {
	c.Mu.RLock()
	defer c.Mu.RUnlock()

	return fmt.Sprintf("CPU: PC=0x%x, SP=0x%x, Cycles=%d, EAX=0x%x, EBX=0x%x",
		c.PC, c.SP, c.Cycles, c.Registers["EAX"], c.Registers["EBX"])
}

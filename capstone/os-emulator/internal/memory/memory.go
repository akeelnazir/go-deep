package memory

import (
	"fmt"
	"sync"
)

const PageSize = 4096
const TotalMemory = 256 * 1024 * 1024

type MemoryManager struct {
	memory    [TotalMemory]byte
	pageTable map[uint32]map[uint32]uint32
	Mu        sync.RWMutex
}

func NewMemoryManager() *MemoryManager {
	return &MemoryManager{
		pageTable: make(map[uint32]map[uint32]uint32),
	}
}

func (mm *MemoryManager) Write(address uint32, data []byte) error {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	if uint64(address)+uint64(len(data)) > uint64(len(mm.memory)) {
		return fmt.Errorf("memory write out of bounds: 0x%x", address)
	}

	copy(mm.memory[address:], data)
	return nil
}

func (mm *MemoryManager) Read(address uint32, size uint32) ([]byte, error) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	if uint64(address)+uint64(size) > uint64(len(mm.memory)) {
		return nil, fmt.Errorf("memory read out of bounds: 0x%x", address)
	}

	data := make([]byte, size)
	copy(data, mm.memory[address:address+size])
	return data, nil
}

func (mm *MemoryManager) ReadByte(address uint32) (byte, error) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	if address >= uint32(len(mm.memory)) {
		return 0, fmt.Errorf("memory read out of bounds: 0x%x", address)
	}

	return mm.memory[address], nil
}

func (mm *MemoryManager) WriteByte(address uint32, value byte) error {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	if address >= uint32(len(mm.memory)) {
		return fmt.Errorf("memory write out of bounds: 0x%x", address)
	}

	mm.memory[address] = value
	return nil
}

func (mm *MemoryManager) AllocatePages(pid uint32, numPages uint32) (uint32, error) {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	if _, exists := mm.pageTable[pid]; !exists {
		mm.pageTable[pid] = make(map[uint32]uint32)
	}

	physicalPage := uint32(len(mm.pageTable[pid]))
	for i := uint32(0); i < numPages; i++ {
		mm.pageTable[pid][i] = physicalPage + i
	}

	return physicalPage * PageSize, nil
}

func (mm *MemoryManager) Translate(pid uint32, virtualAddr uint32) (uint32, error) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	pageTable, exists := mm.pageTable[pid]
	if !exists {
		return 0, fmt.Errorf("no page table for process %d", pid)
	}

	pageNum := virtualAddr / PageSize
	physicalPage, exists := pageTable[pageNum]
	if !exists {
		return 0, fmt.Errorf("page not found for virtual address 0x%x", virtualAddr)
	}

	offset := virtualAddr % PageSize
	return physicalPage*PageSize + offset, nil
}

func (mm *MemoryManager) GetMemoryStats() (used, free uint64) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	used = uint64(len(mm.pageTable)) * uint64(PageSize)
	free = uint64(TotalMemory) - used
	return
}

func (mm *MemoryManager) Clear(address uint32, size uint32) error {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	if uint64(address)+uint64(size) > uint64(len(mm.memory)) {
		return fmt.Errorf("clear out of bounds: 0x%x", address)
	}

	for i := uint32(0); i < size; i++ {
		mm.memory[address+i] = 0
	}
	return nil
}

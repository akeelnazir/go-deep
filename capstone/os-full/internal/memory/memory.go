package memory

import (
	"fmt"
	"sync"
)

const PageSize = 4096

type PageTableEntry struct {
	VirtualAddr uint64
	PhysicalAddr uint64
	Present     bool
	Writable    bool
	Owner       uint32
}

type PageTable struct {
	entries map[uint64]*PageTableEntry
	Mu      sync.RWMutex
}

type MemoryManager struct {
	totalMemory uint64
	usedMemory  uint64
	pageTables  map[uint32]*PageTable
	Mu          sync.RWMutex
}

func NewMemoryManager(totalMemory uint64) *MemoryManager {
	return &MemoryManager{
		totalMemory: totalMemory,
		usedMemory:  0,
		pageTables:  make(map[uint32]*PageTable),
	}
}

func NewPageTable() *PageTable {
	return &PageTable{
		entries: make(map[uint64]*PageTableEntry),
	}
}

func (mm *MemoryManager) CreatePageTable(pid uint32) *PageTable {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	pt := NewPageTable()
	mm.pageTables[pid] = pt
	return pt
}

func (mm *MemoryManager) AllocatePages(pid uint32, numPages uint64) (uint64, error) {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	size := numPages * PageSize
	if mm.usedMemory+size > mm.totalMemory {
		return 0, fmt.Errorf("not enough memory for %d pages", numPages)
	}

	physicalAddr := mm.usedMemory
	mm.usedMemory += size

	pt, exists := mm.pageTables[pid]
	if !exists {
		pt = NewPageTable()
		mm.pageTables[pid] = pt
	}

	pt.Mu.Lock()
	for i := uint64(0); i < numPages; i++ {
		vAddr := i * PageSize
		pAddr := physicalAddr + (i * PageSize)
		pt.entries[vAddr] = &PageTableEntry{
			VirtualAddr:  vAddr,
			PhysicalAddr: pAddr,
			Present:      true,
			Writable:     true,
			Owner:        pid,
		}
	}
	pt.Mu.Unlock()

	return physicalAddr, nil
}

func (mm *MemoryManager) DeallocatePages(pid uint32) error {
	mm.Mu.Lock()
	defer mm.Mu.Unlock()

	pt, exists := mm.pageTables[pid]
	if !exists {
		return fmt.Errorf("page table for process %d not found", pid)
	}

	pt.Mu.Lock()
	for _, entry := range pt.entries {
		if entry.Present {
			mm.usedMemory -= PageSize
		}
	}
	pt.entries = make(map[uint64]*PageTableEntry)
	pt.Mu.Unlock()

	delete(mm.pageTables, pid)
	return nil
}

func (mm *MemoryManager) Translate(pid uint32, virtualAddr uint64) (uint64, error) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	pt, exists := mm.pageTables[pid]
	if !exists {
		return 0, fmt.Errorf("page table for process %d not found", pid)
	}

	pt.Mu.RLock()
	defer pt.Mu.RUnlock()

	pageNum := virtualAddr / PageSize
	entry, exists := pt.entries[pageNum*PageSize]
	if !exists || !entry.Present {
		return 0, fmt.Errorf("page fault at address 0x%x", virtualAddr)
	}

	offset := virtualAddr % PageSize
	return entry.PhysicalAddr + offset, nil
}

func (mm *MemoryManager) GetMemoryStats() (total, used, free uint64) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	total = mm.totalMemory
	used = mm.usedMemory
	free = total - used
	return
}

func (mm *MemoryManager) GetPageCount() (total, used uint64) {
	mm.Mu.RLock()
	defer mm.Mu.RUnlock()

	total = mm.totalMemory / PageSize
	used = mm.usedMemory / PageSize
	return
}

package memory

import (
	"fmt"
	"sync"
)

type MemoryBlock struct {
	Address uint64
	Size    uint64
	Free    bool
	Owner   uint32
}

type MemoryManager struct {
	totalMemory uint64
	blocks      []*MemoryBlock
	mu          sync.RWMutex
}

func NewMemoryManager(totalMemory uint64) *MemoryManager {
	return &MemoryManager{
		totalMemory: totalMemory,
		blocks: []*MemoryBlock{
			{
				Address: 0,
				Size:    totalMemory,
				Free:    true,
				Owner:   0,
			},
		},
	}
}

func (mm *MemoryManager) Allocate(pid uint32, size uint64) (uint64, error) {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	for _, block := range mm.blocks {
		if block.Free && block.Size >= size {
			block.Free = false
			block.Owner = pid

			if block.Size > size {
				newBlock := &MemoryBlock{
					Address: block.Address + size,
					Size:    block.Size - size,
					Free:    true,
					Owner:   0,
				}
				mm.blocks = append(mm.blocks, newBlock)
				block.Size = size
			}

			return block.Address, nil
		}
	}

	return 0, fmt.Errorf("not enough memory for allocation of %d bytes", size)
}

func (mm *MemoryManager) Deallocate(address uint64) error {
	mm.mu.Lock()
	defer mm.mu.Unlock()

	for i, block := range mm.blocks {
		if block.Address == address {
			block.Free = true
			block.Owner = 0

			if i > 0 && mm.blocks[i-1].Free {
				mm.blocks[i-1].Size += block.Size
				mm.blocks = append(mm.blocks[:i], mm.blocks[i+1:]...)
			}

			if i < len(mm.blocks)-1 && mm.blocks[i+1].Free {
				block.Size += mm.blocks[i+1].Size
				mm.blocks = append(mm.blocks[:i+1], mm.blocks[i+2:]...)
			}

			return nil
		}
	}

	return fmt.Errorf("memory block at address %d not found", address)
}

func (mm *MemoryManager) GetUsedMemory() uint64 {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	var used uint64
	for _, block := range mm.blocks {
		if !block.Free {
			used += block.Size
		}
	}
	return used
}

func (mm *MemoryManager) GetFreeMemory() uint64 {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	var free uint64
	for _, block := range mm.blocks {
		if block.Free {
			free += block.Size
		}
	}
	return free
}

func (mm *MemoryManager) GetMemoryStats() (total, used, free uint64) {
	mm.mu.RLock()
	defer mm.mu.RUnlock()

	total = mm.totalMemory
	for _, block := range mm.blocks {
		if !block.Free {
			used += block.Size
		} else {
			free += block.Size
		}
	}
	return
}

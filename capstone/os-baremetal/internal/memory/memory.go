package memory

import "fmt"

type MemoryBlock struct {
	Address uint64
	Size    uint64
	Free    bool
}

type MemoryAllocator struct {
	blocks []MemoryBlock
	total  uint64
}

func NewMemoryAllocator(totalSize uint64) *MemoryAllocator {
	return &MemoryAllocator{
		blocks: []MemoryBlock{
			{
				Address: 0,
				Size:    totalSize,
				Free:    true,
			},
		},
		total: totalSize,
	}
}

func (ma *MemoryAllocator) Allocate(size uint64) (uint64, error) {
	for i, block := range ma.blocks {
		if block.Free && block.Size >= size {
			if block.Size == size {
				ma.blocks[i].Free = false
				return block.Address, nil
			}

			newBlock := MemoryBlock{
				Address: block.Address + size,
				Size:    block.Size - size,
				Free:    true,
			}

			ma.blocks[i].Size = size
			ma.blocks[i].Free = false

			ma.blocks = append(ma.blocks[:i+1], append([]MemoryBlock{newBlock}, ma.blocks[i+1:]...)...)
			return block.Address, nil
		}
	}
	return 0, fmt.Errorf("not enough contiguous memory")
}

func (ma *MemoryAllocator) Free(address uint64) error {
	for i, block := range ma.blocks {
		if block.Address == address && !block.Free {
			ma.blocks[i].Free = true

			if i > 0 && ma.blocks[i-1].Free {
				ma.blocks[i-1].Size += block.Size
				ma.blocks = append(ma.blocks[:i], ma.blocks[i+1:]...)
				i--
			}

			if i < len(ma.blocks)-1 && ma.blocks[i+1].Free {
				ma.blocks[i].Size += ma.blocks[i+1].Size
				ma.blocks = append(ma.blocks[:i+1], ma.blocks[i+2:]...)
			}

			return nil
		}
	}
	return fmt.Errorf("address not found or already free")
}

func (ma *MemoryAllocator) GetStats() (total, used, free uint64) {
	total = ma.total
	for _, block := range ma.blocks {
		if block.Free {
			free += block.Size
		} else {
			used += block.Size
		}
	}
	return
}

func (ma *MemoryAllocator) PrintLayout() {
	fmt.Println("=== Memory Layout ===")
	for i, block := range ma.blocks {
		status := "USED"
		if block.Free {
			status = "FREE"
		}
		fmt.Printf("[%d] 0x%x - 0x%x (%d bytes) [%s]\n",
			i, block.Address, block.Address+block.Size, block.Size, status)
	}
}

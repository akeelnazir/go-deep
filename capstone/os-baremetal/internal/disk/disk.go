package disk

import (
	"fmt"
	"os"
)

type Disk struct {
	filename string
	size     uint64
	file     *os.File
}

type DiskBlock struct {
	BlockID uint64
	Data    [512]byte
	Used    bool
}

func NewDisk(filename string, sizeInMB uint64) *Disk {
	return &Disk{
		filename: filename,
		size:     sizeInMB * 1024 * 1024,
	}
}

func (d *Disk) Initialize() error {
	file, err := os.Create(d.filename)
	if err != nil {
		return fmt.Errorf("failed to create disk file: %v", err)
	}
	d.file = file

	if err := file.Truncate(int64(d.size)); err != nil {
		return fmt.Errorf("failed to allocate disk space: %v", err)
	}

	fmt.Printf("Disk initialized: %s (%d MB)\n", d.filename, d.size/(1024*1024))
	return nil
}

func (d *Disk) ReadBlock(blockID uint64) ([]byte, error) {
	if d.file == nil {
		return nil, fmt.Errorf("disk not initialized")
	}

	offset := blockID * 512
	if offset >= d.size {
		return nil, fmt.Errorf("block out of range")
	}

	buffer := make([]byte, 512)
	n, err := d.file.ReadAt(buffer, int64(offset))
	if err != nil {
		return nil, fmt.Errorf("failed to read block: %v", err)
	}

	return buffer[:n], nil
}

func (d *Disk) WriteBlock(blockID uint64, data []byte) error {
	if d.file == nil {
		return fmt.Errorf("disk not initialized")
	}

	if len(data) > 512 {
		return fmt.Errorf("data too large for block")
	}

	offset := blockID * 512
	if offset >= d.size {
		return fmt.Errorf("block out of range")
	}

	_, err := d.file.WriteAt(data, int64(offset))
	if err != nil {
		return fmt.Errorf("failed to write block: %v", err)
	}

	return nil
}

func (d *Disk) Close() error {
	if d.file != nil {
		return d.file.Close()
	}
	return nil
}

func (d *Disk) GetSize() uint64 {
	return d.size
}

func (d *Disk) GetBlockCount() uint64 {
	return d.size / 512
}

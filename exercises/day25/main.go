package main

import (
	"fmt"
	"unsafe"
)

type MemoryStats struct {
	Allocations int64
	Deallocations int64
	CurrentUsage int64
}

var memStats = &MemoryStats{}

func allocateMemory(size int64) {
	memStats.Allocations++
	memStats.CurrentUsage += size
}

func deallocateMemory(size int64) {
	memStats.Deallocations++
	memStats.CurrentUsage -= size
}

func getMemoryUsage() int64 {
	return memStats.CurrentUsage
}

func getSizeOf(v interface{}) uintptr {
	switch v.(type) {
	case int:
		return unsafe.Sizeof(int(0))
	case int32:
		return unsafe.Sizeof(int32(0))
	case int64:
		return unsafe.Sizeof(int64(0))
	case string:
		return unsafe.Sizeof("")
	case []int:
		return unsafe.Sizeof([]int{})
	default:
		return 0
	}
}

func main() {
	fmt.Println("=== Day 25: Memory Management and Unsafe ===")

	fmt.Println("\n--- Memory Allocation ---")
	allocateMemory(1024)
	allocateMemory(2048)
	fmt.Printf("Current memory usage: %d bytes\n", getMemoryUsage())
	fmt.Printf("Total allocations: %d\n", memStats.Allocations)

	fmt.Println("\n--- Memory Deallocation ---")
	deallocateMemory(1024)
	fmt.Printf("Current memory usage after deallocation: %d bytes\n", getMemoryUsage())
	fmt.Printf("Total deallocations: %d\n", memStats.Deallocations)

	fmt.Println("\n--- Size Information ---")
	fmt.Printf("Size of int: %d bytes\n", getSizeOf(int(0)))
	fmt.Printf("Size of int32: %d bytes\n", getSizeOf(int32(0)))
	fmt.Printf("Size of int64: %d bytes\n", getSizeOf(int64(0)))
	fmt.Printf("Size of string: %d bytes\n", getSizeOf(""))

	fmt.Println("\n--- Pointer Operations ---")
	x := 42
	ptr := unsafe.Pointer(&x)
	fmt.Printf("Pointer address: %v\n", ptr)
	fmt.Printf("Dereferenced value: %d\n", *(*int)(ptr))

	fmt.Println("\n--- Array Memory ---")
	arr := [5]int{1, 2, 3, 4, 5}
	arrPtr := unsafe.Pointer(&arr[0])
	fmt.Printf("Array pointer: %v\n", arrPtr)
	fmt.Printf("Array size: %d bytes\n", unsafe.Sizeof(arr))

	fmt.Println("\n=== Day 25 Complete ===")
	fmt.Println("Next: Learn about web frameworks on Day 26.")
}

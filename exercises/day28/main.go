package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Push(value int) {
	ll.Head = &Node{Value: value, Next: ll.Head}
}

func (ll *LinkedList) Pop() int {
	if ll.Head == nil {
		return 0
	}
	value := ll.Head.Value
	ll.Head = ll.Head.Next
	return value
}

func (ll *LinkedList) Peek() int {
	if ll.Head == nil {
		return 0
	}
	return ll.Head.Value
}

func (ll *LinkedList) Size() int {
	count := 0
	current := ll.Head
	for current != nil {
		count++
		current = current.Next
	}
	return count
}

func MapInt(fn func(int) int, nums []int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

func FilterInt(fn func(int) bool, nums []int) []int {
	var result []int
	for _, n := range nums {
		if fn(n) {
			result = append(result, n)
		}
	}
	return result
}

func ReduceInt(fn func(int, int) int, nums []int, initial int) int {
	result := initial
	for _, n := range nums {
		result = fn(result, n)
	}
	return result
}

func main() {
	fmt.Println("=== Day 28: Advanced Go ===")

	fmt.Println("\n--- Linked List ---")
	ll := &LinkedList{}
	ll.Push(1)
	ll.Push(2)
	ll.Push(3)
	fmt.Printf("List size: %d\n", ll.Size())
	fmt.Printf("Peek: %d\n", ll.Peek())
	fmt.Printf("Pop: %d\n", ll.Pop())
	fmt.Printf("Size after pop: %d\n", ll.Size())

	fmt.Println("\n--- Map Function ---")
	nums := []int{1, 2, 3, 4, 5}
	doubled := MapInt(func(x int) int { return x * 2 }, nums)
	fmt.Printf("Original: %v\n", nums)
	fmt.Printf("Doubled: %v\n", doubled)

	fmt.Println("\n--- Filter Function ---")
	evens := FilterInt(func(x int) bool { return x%2 == 0 }, nums)
	fmt.Printf("Evens: %v\n", evens)

	fmt.Println("\n--- Reduce Function ---")
	sum := ReduceInt(func(acc, x int) int { return acc + x }, nums, 0)
	product := ReduceInt(func(acc, x int) int { return acc * x }, nums, 1)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Product: %d\n", product)

	fmt.Println("\n=== Day 28 Complete ===")
	fmt.Println("Next: Final review on Day 29.")
}

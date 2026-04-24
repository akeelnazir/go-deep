package main

import (
	"fmt"
	"sync"
	"time"
)

// Section 1: Simple Goroutine
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Section 2: Worker function for WaitGroup
func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("Worker %d done\n", id)
}

// Section 3: Demonstrate unbuffered channel
func demonstrateUnbufferedChannel() {
	fmt.Println("\n=== Unbuffered Channel ===")
	ch := make(chan int)

	go func() {
		fmt.Println("Sending 42 to channel")
		ch <- 42
		fmt.Println("Sent 42")
	}()

	fmt.Println("Waiting to receive")
	value := <-ch
	fmt.Printf("Received: %d\n", value)
}

// Section 4: Demonstrate buffered channel
func demonstrateBufferedChannel() {
	fmt.Println("\n=== Buffered Channel ===")
	ch := make(chan int, 2)

	ch <- 1
	fmt.Println("Sent 1")
	ch <- 2
	fmt.Println("Sent 2")

	fmt.Printf("Received: %d\n", <-ch)
	fmt.Printf("Received: %d\n", <-ch)
}

// Section 5: Channel with range
func demonstrateChannelRange() {
	fmt.Println("\n=== Channel with Range ===")
	ch := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}
}

// Section 6: Directional channels
func sender(ch chan<- int) {
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	close(ch)
}

func receiver(ch <-chan int) {
	for val := range ch {
		fmt.Printf("Received: %d\n", val)
	}
}

// Section 7: Multiple goroutines with WaitGroup
func demonstrateWaitGroup() {
	fmt.Println("\n=== WaitGroup Synchronization ===")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("All workers completed")
}

// Section 8: Multiple goroutines with channel
func demonstrateMultipleGoroutinesWithChannel() {
	fmt.Println("\n=== Multiple Goroutines with Channel ===")
	var wg sync.WaitGroup
	results := make(chan int)

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			results <- n * n
		}(i)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for val := range results {
		fmt.Printf("Result: %d\n", val)
	}
}

// Section 9: Detecting channel closure
func demonstrateChannelClosure() {
	fmt.Println("\n=== Detecting Channel Closure ===")
	ch := make(chan int)

	go func() {
		ch <- 1
		ch <- 2
		close(ch)
	}()

	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed")
			break
		}
		fmt.Printf("Received: %d\n", val)
	}
}

// Section 10: Goroutine leaks example (commented out)
func demonstrateGoroutineLeakPrevention() {
	fmt.Println("\n=== Goroutine Leak Prevention ===")
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Printf("Received: %d\n", val)
	}()

	ch <- 42
	time.Sleep(100 * time.Millisecond)
}

func main() {
	fmt.Println("=== Day 6: Concurrency Fundamentals ===")

	// Section 1: Basic goroutines
	fmt.Println("\n--- Launching Goroutines ---")
	go greet("Alice")
	go greet("Bob")
	go greet("Charlie")
	time.Sleep(100 * time.Millisecond)

	// Section 2: Unbuffered channels
	demonstrateUnbufferedChannel()

	// Section 3: Buffered channels
	demonstrateBufferedChannel()

	// Section 4: Channel with range
	demonstrateChannelRange()

	// Section 5: Directional channels
	fmt.Println("\n--- Directional Channels ---")
	ch := make(chan int)
	go sender(ch)
	receiver(ch)

	// Section 6: WaitGroup synchronization
	demonstrateWaitGroup()

	// Section 7: Multiple goroutines with channel
	demonstrateMultipleGoroutinesWithChannel()

	// Section 8: Detecting channel closure
	demonstrateChannelClosure()

	// Section 9: Goroutine leak prevention
	demonstrateGoroutineLeakPrevention()

	fmt.Println("\n=== Day 6 Complete ===")
	fmt.Println("Next: Learn about advanced concurrency patterns on Day 7.")
}

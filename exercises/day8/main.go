package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Section 1: Context Cancellation
func demonstrateCancellation(ctx context.Context, id int, done <-chan struct{}) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
			return
		case <-done:
			fmt.Printf("Worker %d received done signal\n", id)
			return
		default:
			fmt.Printf("Worker %d working\n", id)
			time.Sleep(100 * time.Millisecond)
		}
	}
}

// Section 2: Context Timeout
func fetchDataWithTimeout(ctx context.Context, duration time.Duration) (string, error) {
	select {
	case <-time.After(duration):
		return "data retrieved", nil
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// Section 3: Context Deadline
func operationWithDeadline(ctx context.Context) error {
	select {
	case <-time.After(200 * time.Millisecond):
		fmt.Println("Operation completed successfully")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Section 4: Job and Result types for Worker Pool
type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
	Err    error
}

// Section 5: Worker Pool with Context
func worker(ctx context.Context, id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}

			select {
			case <-ctx.Done():
				return
			case <-time.After(50 * time.Millisecond):
				results <- Result{
					JobID:  job.ID,
					Output: fmt.Sprintf("Worker %d processed: %s", id, job.Data),
				}
			}
		}
	}
}

// Section 6: Pipeline Stage - Generate
func generate(ctx context.Context, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Section 7: Pipeline Stage - Square
func square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Section 8: Pipeline Stage - Cube
func cube(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n * n:
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

// Section 9: Fan-Out - distribute work to multiple workers
func fanOut(ctx context.Context, in <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		ch := make(chan int)
		channels[i] = ch
		go func(out chan<- int) {
			defer close(out)
			for val := range in {
				select {
				case out <- val:
				case <-ctx.Done():
					return
				}
			}
		}(ch)
	}
	return channels
}

// Section 10: Fan-In - merge multiple channels
func fanIn(ctx context.Context, channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	out := make(chan int)

	output := func(ch <-chan int) {
		defer wg.Done()
		for {
			select {
			case val, ok := <-ch:
				if !ok {
					return
				}
				select {
				case out <- val:
				case <-ctx.Done():
					return
				}
			case <-ctx.Done():
				return
			}
		}
	}

	wg.Add(len(channels))
	for _, ch := range channels {
		go output(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("=== Day 8: Context and Advanced Concurrency ===")

	// Section 1: Context Cancellation
	fmt.Println("\n--- Context Cancellation ---")
	ctx, cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	done := make(chan struct{})

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			demonstrateCancellation(ctx, id, done)
		}(i)
	}

	time.Sleep(300 * time.Millisecond)
	cancel()
	wg.Wait()

	// Section 2: Context Timeout
	fmt.Println("\n--- Context Timeout ---")
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	data, err := fetchDataWithTimeout(ctx, 500*time.Millisecond)
	if err != nil {
		fmt.Printf("Timeout error: %v\n", err)
	} else {
		fmt.Printf("Result: %s\n", data)
	}

	// Section 3: Context Deadline
	fmt.Println("\n--- Context Deadline ---")
	deadline := time.Now().Add(1 * time.Second)
	ctx, cancel = context.WithDeadline(context.Background(), deadline)
	defer cancel()

	err = operationWithDeadline(ctx)
	if err != nil {
		fmt.Printf("Deadline error: %v\n", err)
	}

	// Section 4: Worker Pool with Context
	fmt.Println("\n--- Worker Pool with Context ---")
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	numWorkers := 2
	numJobs := 5

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(ctx, w, jobs, results, &wg)
	}

	go func() {
		for i := 1; i <= numJobs; i++ {
			jobs <- Job{ID: i, Data: fmt.Sprintf("Job %d", i)}
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Printf("Result %d: %s\n", result.JobID, result.Output)
	}

	// Section 5: Pipeline Pattern
	fmt.Println("\n--- Pipeline Pattern ---")
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	numbers := []int{1, 2, 3, 4, 5}
	squared := square(ctx, generate(ctx, numbers...))
	cubed := cube(ctx, squared)

	for result := range cubed {
		fmt.Printf("Result: %d\n", result)
	}

	// Section 6: Fan-Out/Fan-In Pattern
	fmt.Println("\n--- Fan-Out/Fan-In Pattern ---")
	ctx, cancel = context.WithCancel(context.Background())
	defer cancel()

	input := make(chan int)
	go func() {
		for i := 1; i <= 4; i++ {
			input <- i
		}
		close(input)
	}()

	workers := fanOut(ctx, input, 2)
	merged := fanIn(ctx, workers...)

	for result := range merged {
		fmt.Printf("Merged result: %d\n", result)
	}

	// Section 7: Graceful Shutdown
	fmt.Println("\n--- Graceful Shutdown Pattern ---")
	ctx, cancel = context.WithCancel(context.Background())

	done = make(chan struct{})
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for {
				select {
				case <-done:
					fmt.Printf("Goroutine %d shutting down\n", id)
					return
				default:
					fmt.Printf("Goroutine %d working\n", id)
					time.Sleep(100 * time.Millisecond)
				}
			}
		}(i)
	}

	time.Sleep(250 * time.Millisecond)
	close(done)
	wg.Wait()

	cancel()

	fmt.Println("\n=== Day 8 Complete ===")
	fmt.Println("Next: Learn about error handling on Day 9.")
}

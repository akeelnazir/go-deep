package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Section 1: Mutex-protected Counter
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// Section 2: RWMutex-protected Cache
type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{data: make(map[string]string)}
}

func (c *Cache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Section 3: Atomic Counter
type AtomicCounter struct {
	value int64
}

func (ac *AtomicCounter) Increment() {
	atomic.AddInt64(&ac.value, 1)
}

func (ac *AtomicCounter) Value() int64 {
	return atomic.LoadInt64(&ac.value)
}

// Section 4: Worker Pool
type Job struct {
	ID   int
	Data string
}

type Result struct {
	JobID  int
	Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		result := Result{
			JobID:  job.ID,
			Output: fmt.Sprintf("Worker %d processed: %s", id, job.Data),
		}
		results <- result
	}
}

// Section 5: Rate Limiter
type RateLimiter struct {
	limiter <-chan time.Time
}

func NewRateLimiter(rps int) *RateLimiter {
	limiter := time.Tick(time.Second / time.Duration(rps))
	return &RateLimiter{limiter}
}

func (rl *RateLimiter) Wait() {
	<-rl.limiter
}

func main() {
	fmt.Println("=== Day 7: Synchronization Patterns ===")

	// Section 1: Mutex-protected Counter
	fmt.Println("\n--- Mutex-protected Counter ---")
	counter := &Counter{}
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.Value())

	// Section 2: RWMutex-protected Cache
	fmt.Println("\n--- RWMutex-protected Cache ---")
	cache := NewCache()

	cache.Set("key1", "value1")
	cache.Set("key2", "value2")

	fmt.Printf("key1: %s\n", cache.Get("key1"))
	fmt.Printf("key2: %s\n", cache.Get("key2"))

	// Section 3: Atomic Counter
	fmt.Println("\n--- Atomic Counter ---")
	atomicCounter := &AtomicCounter{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				atomicCounter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final atomic counter value: %d\n", atomicCounter.Value())

	// Section 4: Worker Pool
	fmt.Println("\n--- Worker Pool ---")
	numWorkers := 3
	numJobs := 10

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
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

	// Section 5: Rate Limiter
	fmt.Println("\n--- Rate Limiter (5 ops/sec) ---")
	limiter := NewRateLimiter(5)

	for i := 1; i <= 5; i++ {
		limiter.Wait()
		fmt.Printf("Operation %d at %v\n", i, time.Now().Format("15:04:05.000"))
	}

	fmt.Println("\n=== Day 7 Complete ===")
	fmt.Println("Next: Learn about interfaces on Day 8.")
}

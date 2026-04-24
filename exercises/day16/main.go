package main

import (
	"fmt"
	"time"
)

type CircuitBreaker struct {
	maxFailures int
	timeout     time.Duration
	failures    int
	lastFailure time.Time
	state       string
}

func NewCircuitBreaker(maxFailures int, timeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		maxFailures: maxFailures,
		timeout:     timeout,
		state:       "closed",
	}
}

func (cb *CircuitBreaker) Call(fn func() error) error {
	if cb.state == "open" {
		if time.Since(cb.lastFailure) > cb.timeout {
			cb.state = "half-open"
			cb.failures = 0
		} else {
			return fmt.Errorf("circuit breaker is open")
		}
	}

	err := fn()

	if err != nil {
		cb.failures++
		cb.lastFailure = time.Now()

		if cb.failures >= cb.maxFailures {
			cb.state = "open"
		}
		return err
	}

	if cb.state == "half-open" {
		cb.state = "closed"
		cb.failures = 0
	}

	return nil
}

func (cb *CircuitBreaker) State() string {
	return cb.state
}

func retryWithBackoff(fn func() error, maxRetries int) error {
	var err error
	backoff := time.Millisecond * 100

	for i := 0; i < maxRetries; i++ {
		err = fn()
		if err == nil {
			return nil
		}

		if i < maxRetries-1 {
			time.Sleep(backoff)
			backoff *= 2
		}
	}

	return err
}

func main() {
	fmt.Println("=== Day 16: Microservices and gRPC ===")

	fmt.Println("\n--- Circuit Breaker Pattern ---")
	cb := NewCircuitBreaker(3, 2*time.Second)

	failCount := 0
	operation := func() error {
		failCount++
		if failCount <= 3 {
			return fmt.Errorf("service unavailable")
		}
		return nil
	}

	for i := 0; i < 5; i++ {
		err := cb.Call(operation)
		fmt.Printf("Attempt %d: State=%s, Error=%v\n", i+1, cb.State(), err)
	}

	fmt.Println("\n--- Retry with Exponential Backoff ---")
	retryCount := 0
	retryOp := func() error {
		retryCount++
		if retryCount < 3 {
			return fmt.Errorf("attempt %d failed", retryCount)
		}
		return nil
	}

	err := retryWithBackoff(retryOp, 5)
	if err != nil {
		fmt.Printf("Failed after retries: %v\n", err)
	} else {
		fmt.Printf("Succeeded after %d attempts\n", retryCount)
	}

	fmt.Println("\n=== Day 16 Complete ===")
	fmt.Println("Next: Learn about testing web applications on Day 17.")
}

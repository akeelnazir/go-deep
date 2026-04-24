package main

// TODO: Implement ExerciseMutexCounter function
// Should create a Counter, increment it from 5 goroutines (100 times each)
// Use sync.WaitGroup to synchronize, return the final value
func ExerciseMutexCounter() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseRWMutexCache function
// Should create a Cache, set 3 key-value pairs, then get them back
// Return the value of "key2"
func ExerciseRWMutexCache() string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseAtomicCounter function
// Should create an AtomicCounter, increment it from 5 goroutines (100 times each)
// Use sync.WaitGroup to synchronize, return the final value
func ExerciseAtomicCounter() int64 {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseWorkerPool function
// Should create a worker pool with 2 workers processing 5 jobs
// Return the count of results received
func ExerciseWorkerPool() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseChannelWithMutex function
// Should use a channel and mutex together to safely increment a counter
// Send 10 values through the channel, increment counter for each
// Return the final counter value
func ExerciseChannelWithMutex() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseWaitGroupWithChannel function
// Should use WaitGroup and channel to collect results from 3 goroutines
// Each goroutine sends a number (1, 2, 3) to the channel
// Return the sum of all received numbers
func ExerciseWaitGroupWithChannel() int {
	// TODO: Add logic
	return 0
}

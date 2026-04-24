package main

// TODO: Implement ExerciseSimpleGoroutine function
// Should launch a goroutine that sends a message to a channel
// The goroutine should send the string "Hello from goroutine" to the channel
// Return the received message from the channel
func ExerciseSimpleGoroutine() string {
	// TODO: Add logic
	return ""
}

// TODO: Implement ExerciseWaitGroupCounter function
// Should use sync.WaitGroup to count from 1 to n
// Launch n goroutines, each incrementing a counter
// Return the final counter value
// Use a pointer to sync.Mutex to protect the counter from race conditions
func ExerciseWaitGroupCounter(n int) int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseBufferedChannelSum function
// Should create a buffered channel with capacity 3
// Send three integers (1, 2, 3) to the channel
// Receive all values and return their sum
func ExerciseBufferedChannelSum() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseChannelRange function
// Should create a channel and a goroutine that sends values 1 to 5
// Close the channel after sending all values
// Use a for range loop to receive all values
// Return the sum of all received values
func ExerciseChannelRange() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseDirectionalChannels function
// Should create a channel and use directional channel parameters
// Launch a sender goroutine that sends values 1, 2, 3
// Launch a receiver goroutine that receives all values
// Return the sum of received values
func ExerciseDirectionalChannels() int {
	// TODO: Add logic
	return 0
}

// TODO: Implement ExerciseDetectChannelClosure function
// Should create a channel and a goroutine that sends 3 values then closes
// Use the ok flag to detect when the channel is closed
// Return the count of values received before closure
func ExerciseDetectChannelClosure() int {
	// TODO: Add logic
	return 0
}

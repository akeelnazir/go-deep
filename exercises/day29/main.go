package main

import (
	"fmt"
)

type LearningProgress struct {
	Day           int
	Topic         string
	Completed     bool
	Exercises     int
	TestsPassed   int
}

var progress []LearningProgress

func init() {
	for i := 1; i <= 29; i++ {
		progress = append(progress, LearningProgress{
			Day:       i,
			Completed: false,
			Exercises: 0,
			TestsPassed: 0,
		})
	}
}

func markDayComplete(day int) {
	if day > 0 && day <= len(progress) {
		progress[day-1].Completed = true
	}
}

func recordExercise(day int, passed bool) {
	if day > 0 && day <= len(progress) {
		progress[day-1].Exercises++
		if passed {
			progress[day-1].TestsPassed++
		}
	}
}

func getCompletionPercentage() float64 {
	completed := 0
	for _, p := range progress {
		if p.Completed {
			completed++
		}
	}
	return float64(completed) / float64(len(progress)) * 100
}

func getTotalTestsPassed() int {
	total := 0
	for _, p := range progress {
		total += p.TestsPassed
	}
	return total
}

func main() {
	fmt.Println("=== Day 29: Final Review and Next Steps ===")

	fmt.Println("\n--- Learning Progress ---")
	for i := 1; i <= 5; i++ {
		markDayComplete(i)
		recordExercise(i, true)
		recordExercise(i, true)
	}

	fmt.Printf("Completion: %.1f%%\n", getCompletionPercentage())
	fmt.Printf("Total tests passed: %d\n", getTotalTestsPassed())

	fmt.Println("\n--- Completed Days ---")
	for _, p := range progress {
		if p.Completed {
			fmt.Printf("Day %d: %d exercises, %d tests passed\n", p.Day, p.Exercises, p.TestsPassed)
		}
	}

	fmt.Println("\n--- Remaining Days ---")
	remaining := 0
	for _, p := range progress {
		if !p.Completed {
			remaining++
		}
	}
	fmt.Printf("Days remaining: %d\n", remaining)

	fmt.Println("\n--- Next Steps ---")
	fmt.Println("1. Complete remaining exercises")
	fmt.Println("2. Build a capstone project")
	fmt.Println("3. Contribute to open source")
	fmt.Println("4. Explore advanced topics")
	fmt.Println("5. Engage with the Go community")

	fmt.Println("\n=== Congratulations! ===")
	fmt.Println("You've completed the 29-day Go learning curriculum!")
	fmt.Println("Continue your journey with real-world projects and community engagement.")
}

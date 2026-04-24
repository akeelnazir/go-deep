package main

import (
	"testing"
)

func TestExerciseMarkDayComplete(t *testing.T) {
	progress = []LearningProgress{}
	for i := 1; i <= 5; i++ {
		progress = append(progress, LearningProgress{Day: i, Completed: false})
	}

	ExerciseMarkDayComplete(1)
	if !progress[0].Completed {
		t.Errorf("Day 1 should be marked complete")
	}
}

func TestExerciseRecordExercise(t *testing.T) {
	progress = []LearningProgress{}
	for i := 1; i <= 5; i++ {
		progress = append(progress, LearningProgress{Day: i, Exercises: 0, TestsPassed: 0})
	}

	ExerciseRecordExercise(1, true)
	if progress[0].Exercises != 1 || progress[0].TestsPassed != 1 {
		t.Errorf("Exercise not recorded correctly")
	}

	ExerciseRecordExercise(1, false)
	if progress[0].Exercises != 2 || progress[0].TestsPassed != 1 {
		t.Errorf("Failed exercise not recorded correctly")
	}
}

func TestExerciseGetCompletionPercentage(t *testing.T) {
	progress = []LearningProgress{}
	for i := 1; i <= 10; i++ {
		progress = append(progress, LearningProgress{Day: i, Completed: false})
	}

	percentage := ExerciseGetCompletionPercentage()
	if percentage != 0 {
		t.Errorf("ExerciseGetCompletionPercentage() = %.1f, want 0", percentage)
	}

	ExerciseMarkDayComplete(1)
	ExerciseMarkDayComplete(2)

	percentage = ExerciseGetCompletionPercentage()
	if percentage != 20 {
		t.Errorf("ExerciseGetCompletionPercentage() = %.1f, want 20", percentage)
	}
}

func TestExerciseGetTotalTestsPassed(t *testing.T) {
	progress = []LearningProgress{}
	for i := 1; i <= 3; i++ {
		progress = append(progress, LearningProgress{Day: i, TestsPassed: 0})
	}

	ExerciseRecordExercise(1, true)
	ExerciseRecordExercise(1, true)
	ExerciseRecordExercise(2, true)

	total := ExerciseGetTotalTestsPassed()
	if total != 3 {
		t.Errorf("ExerciseGetTotalTestsPassed() = %d, want 3", total)
	}
}

func TestExerciseGetCompletedDaysCount(t *testing.T) {
	progress = []LearningProgress{}
	for i := 1; i <= 5; i++ {
		progress = append(progress, LearningProgress{Day: i, Completed: false})
	}

	count := ExerciseGetCompletedDaysCount()
	if count != 0 {
		t.Errorf("ExerciseGetCompletedDaysCount() = %d, want 0", count)
	}

	ExerciseMarkDayComplete(1)
	ExerciseMarkDayComplete(2)

	count = ExerciseGetCompletedDaysCount()
	if count != 2 {
		t.Errorf("ExerciseGetCompletedDaysCount() = %d, want 2", count)
	}
}

func TestExerciseGetRemainingDaysCount(t *testing.T) {
	progress = []LearningProgress{}
	for i := 1; i <= 5; i++ {
		progress = append(progress, LearningProgress{Day: i, Completed: false})
	}

	count := ExerciseGetRemainingDaysCount()
	if count != 5 {
		t.Errorf("ExerciseGetRemainingDaysCount() = %d, want 5", count)
	}

	ExerciseMarkDayComplete(1)

	count = ExerciseGetRemainingDaysCount()
	if count != 4 {
		t.Errorf("ExerciseGetRemainingDaysCount() = %d, want 4", count)
	}
}

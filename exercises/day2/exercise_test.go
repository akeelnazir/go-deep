package main

import (
	"testing"
)

func TestExerciseIsAdult(t *testing.T) {
	tests := []struct {
		name string
		age  int
		want bool
	}{
		{"adult", 25, true},
		{"exactly 18", 18, true},
		{"minor", 17, false},
		{"child", 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIsAdult(tt.age); got != tt.want {
				t.Errorf("ExerciseIsAdult(%d) = %v, want %v", tt.age, got, tt.want)
			}
		})
	}
}

func TestExerciseGetGrade(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{"A grade", 95, "A"},
		{"B grade", 85, "B"},
		{"C grade", 75, "C"},
		{"F grade", 60, "F"},
		{"boundary A", 90, "A"},
		{"boundary B", 80, "B"},
		{"boundary C", 70, "C"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseGetGrade(tt.score); got != tt.want {
				t.Errorf("ExerciseGetGrade(%d) = %q, want %q", tt.score, got, tt.want)
			}
		})
	}
}

func TestExerciseIsInRange(t *testing.T) {
	tests := []struct {
		name string
		num  int
		min  int
		max  int
		want bool
	}{
		{"in range", 5, 1, 10, true},
		{"at min", 1, 1, 10, true},
		{"at max", 10, 1, 10, true},
		{"below range", 0, 1, 10, false},
		{"above range", 11, 1, 10, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIsInRange(tt.num, tt.min, tt.max); got != tt.want {
				t.Errorf("ExerciseIsInRange(%d, %d, %d) = %v, want %v", tt.num, tt.min, tt.max, got, tt.want)
			}
		})
	}
}

func TestExerciseGetDayName(t *testing.T) {
	tests := []struct {
		name string
		day  int
		want string
	}{
		{"Monday", 1, "Monday"},
		{"Wednesday", 3, "Wednesday"},
		{"Friday", 5, "Friday"},
		{"Sunday", 7, "Sunday"},
		{"invalid 0", 0, ""},
		{"invalid 8", 8, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseGetDayName(tt.day); got != tt.want {
				t.Errorf("ExerciseGetDayName(%d) = %q, want %q", tt.day, got, tt.want)
			}
		})
	}
}

func TestExerciseIsVowel(t *testing.T) {
	tests := []struct {
		name string
		char rune
		want bool
	}{
		{"a", 'a', true},
		{"e", 'e', true},
		{"i", 'i', true},
		{"o", 'o', true},
		{"u", 'u', true},
		{"b", 'b', false},
		{"z", 'z', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIsVowel(tt.char); got != tt.want {
				t.Errorf("ExerciseIsVowel(%c) = %v, want %v", tt.char, got, tt.want)
			}
		})
	}
}

func TestExerciseCountToN(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"1", 1, 1},
		{"5", 5, 15},
		{"10", 10, 55},
		{"0", 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseCountToN(tt.n); got != tt.want {
				t.Errorf("ExerciseCountToN(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestExerciseFindMax(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		c    int
		want int
	}{
		{"first max", 10, 5, 3, 10},
		{"second max", 5, 10, 3, 10},
		{"third max", 5, 3, 10, 10},
		{"all equal", 5, 5, 5, 5},
		{"negatives", -1, -5, -3, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseFindMax(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("ExerciseFindMax(%d, %d, %d) = %d, want %d", tt.a, tt.b, tt.c, got, tt.want)
			}
		})
	}
}

func TestExerciseIsEven(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"even", 4, true},
		{"odd", 5, false},
		{"zero", 0, true},
		{"negative even", -4, true},
		{"negative odd", -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseIsEven(tt.num); got != tt.want {
				t.Errorf("ExerciseIsEven(%d) = %v, want %v", tt.num, got, tt.want)
			}
		})
	}
}

// Array Tests

func TestExerciseArraySum(t *testing.T) {
	tests := []struct {
		name string
		arr  [5]int
		want int
	}{
		{"positive numbers", [5]int{1, 2, 3, 4, 5}, 15},
		{"with zeros", [5]int{0, 1, 2, 3, 4}, 10},
		{"all zeros", [5]int{0, 0, 0, 0, 0}, 0},
		{"negative numbers", [5]int{-1, -2, -3, -4, -5}, -15},
		{"mixed", [5]int{10, -5, 20, -10, 5}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseArraySum(tt.arr); got != tt.want {
				t.Errorf("ExerciseArraySum(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

func TestExerciseArrayMax(t *testing.T) {
	tests := []struct {
		name string
		arr  [5]int
		want int
	}{
		{"positive numbers", [5]int{1, 2, 3, 4, 5}, 5},
		{"first is max", [5]int{10, 2, 3, 4, 5}, 10},
		{"last is max", [5]int{1, 2, 3, 4, 10}, 10},
		{"negative numbers", [5]int{-1, -2, -3, -4, -5}, -1},
		{"mixed", [5]int{-10, 5, 20, -5, 15}, 20},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseArrayMax(tt.arr); got != tt.want {
				t.Errorf("ExerciseArrayMax(%v) = %d, want %d", tt.arr, got, tt.want)
			}
		})
	}
}

func TestExerciseArrayReverse(t *testing.T) {
	tests := []struct {
		name string
		arr  [5]int
		want [5]int
	}{
		{"positive numbers", [5]int{1, 2, 3, 4, 5}, [5]int{5, 4, 3, 2, 1}},
		{"with zeros", [5]int{0, 1, 2, 3, 4}, [5]int{4, 3, 2, 1, 0}},
		{"negative numbers", [5]int{-1, -2, -3, -4, -5}, [5]int{-5, -4, -3, -2, -1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseArrayReverse(tt.arr); got != tt.want {
				t.Errorf("ExerciseArrayReverse(%v) = %v, want %v", tt.arr, got, tt.want)
			}
		})
	}
}

// Slice Tests

func TestExerciseSliceSum(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, 15},
		{"empty slice", []int{}, 0},
		{"single element", []int{42}, 42},
		{"negative numbers", []int{-1, -2, -3}, -6},
		{"mixed", []int{10, -5, 20, -10}, 15},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseSliceSum(tt.slice); got != tt.want {
				t.Errorf("ExerciseSliceSum(%v) = %d, want %d", tt.slice, got, tt.want)
			}
		})
	}
}

func TestExerciseSliceFilter(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, []int{2, 4}},
		{"empty slice", []int{}, []int{}},
		{"no evens", []int{1, 3, 5}, []int{}},
		{"all evens", []int{2, 4, 6}, []int{2, 4, 6}},
		{"mixed", []int{10, 15, 20, 25, 30}, []int{10, 20, 30}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseSliceFilter(tt.slice)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseSliceFilter(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

func TestExerciseSliceContains(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		value int
		want  bool
	}{
		{"contains", []int{1, 2, 3, 4, 5}, 3, true},
		{"not contains", []int{1, 2, 3, 4, 5}, 10, false},
		{"empty slice", []int{}, 1, false},
		{"first element", []int{1, 2, 3}, 1, true},
		{"last element", []int{1, 2, 3}, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseSliceContains(tt.slice, tt.value); got != tt.want {
				t.Errorf("ExerciseSliceContains(%v, %d) = %v, want %v", tt.slice, tt.value, got, tt.want)
			}
		})
	}
}

func TestExerciseSliceReverse(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"positive numbers", []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{42}, []int{42}},
		{"two elements", []int{1, 2}, []int{2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseSliceReverse(tt.slice)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseSliceReverse(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

func TestExerciseSliceDuplicate(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"positive numbers", []int{1, 2, 3}, []int{1, 1, 2, 2, 3, 3}},
		{"empty slice", []int{}, []int{}},
		{"single element", []int{5}, []int{5, 5}},
		{"with zeros", []int{0, 1, 2}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseSliceDuplicate(tt.slice)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseSliceDuplicate(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

// Map Tests

func TestExerciseMapFrequency(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want map[rune]int
	}{
		{"hello", "hello", map[rune]int{'h': 1, 'e': 1, 'l': 2, 'o': 1}},
		{"empty string", "", map[rune]int{}},
		{"single char", "a", map[rune]int{'a': 1}},
		{"repeated chars", "aaa", map[rune]int{'a': 3}},
		{"mixed", "aabbcc", map[rune]int{'a': 2, 'b': 2, 'c': 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseMapFrequency(tt.s)
			if !mapsEqual(got, tt.want) {
				t.Errorf("ExerciseMapFrequency(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestExerciseMapInvert(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want map[int]string
	}{
		{"simple", map[string]int{"a": 1, "b": 2}, map[int]string{1: "a", 2: "b"}},
		{"empty", map[string]int{}, map[int]string{}},
		{"single", map[string]int{"x": 10}, map[int]string{10: "x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseMapInvert(tt.m)
			if !intStringMapsEqual(got, tt.want) {
				t.Errorf("ExerciseMapInvert(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}

func TestExerciseMapMerge(t *testing.T) {
	tests := []struct {
		name string
		m1   map[string]int
		m2   map[string]int
		want map[string]int
	}{
		{"no overlap", map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{"with overlap", map[string]int{"a": 1, "b": 2}, map[string]int{"b": 20, "c": 3}, map[string]int{"a": 1, "b": 20, "c": 3}},
		{"empty m1", map[string]int{}, map[string]int{"a": 1}, map[string]int{"a": 1}},
		{"empty m2", map[string]int{"a": 1}, map[string]int{}, map[string]int{"a": 1}},
		{"both empty", map[string]int{}, map[string]int{}, map[string]int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseMapMerge(tt.m1, tt.m2)
			if !stringIntMapsEqual(got, tt.want) {
				t.Errorf("ExerciseMapMerge(%v, %v) = %v, want %v", tt.m1, tt.m2, got, tt.want)
			}
		})
	}
}

func TestExerciseMapKeys(t *testing.T) {
	tests := []struct {
		name string
		m    map[string]int
		want []string
	}{
		{"simple", map[string]int{"a": 1, "b": 2, "c": 3}, []string{"a", "b", "c"}},
		{"empty", map[string]int{}, []string{}},
		{"single", map[string]int{"x": 10}, []string{"x"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseMapKeys(tt.m)
			if !stringSlicesEqualUnordered(got, tt.want) {
				t.Errorf("ExerciseMapKeys(%v) = %v, want %v", tt.m, got, tt.want)
			}
		})
	}
}

// Combined Tests

func TestExerciseSliceMultiplyByTwo(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"positive numbers", []int{1, 2, 3}, []int{2, 4, 6}},
		{"empty slice", []int{}, []int{}},
		{"with zeros", []int{0, 1, 2}, []int{0, 2, 4}},
		{"negative numbers", []int{-1, -2, -3}, []int{-2, -4, -6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseSliceMultiplyByTwo(tt.slice)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseSliceMultiplyByTwo(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

func TestExerciseMapCountVowels(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want map[rune]int
	}{
		{"hello", "hello", map[rune]int{'e': 1, 'o': 1}},
		{"aeiou", "aeiou", map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}},
		{"no vowels", "bcdfg", map[rune]int{}},
		{"repeated vowels", "aaa", map[rune]int{'a': 3}},
		{"mixed", "hello world", map[rune]int{'e': 1, 'o': 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseMapCountVowels(tt.s)
			if !mapsEqual(got, tt.want) {
				t.Errorf("ExerciseMapCountVowels(%q) = %v, want %v", tt.s, got, tt.want)
			}
		})
	}
}

func TestExerciseSliceRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name  string
		slice []int
		want  []int
	}{
		{"with duplicates", []int{1, 2, 2, 3, 3, 3}, []int{1, 2, 3}},
		{"no duplicates", []int{1, 2, 3}, []int{1, 2, 3}},
		{"empty slice", []int{}, []int{}},
		{"all same", []int{5, 5, 5}, []int{5}},
		{"single element", []int{42}, []int{42}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseSliceRemoveDuplicates(tt.slice)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseSliceRemoveDuplicates(%v) = %v, want %v", tt.slice, got, tt.want)
			}
		})
	}
}

// Helper functions for comparing slices and maps

func slicesEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func mapsEqual(a, b map[rune]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func intStringMapsEqual(a, b map[int]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func stringIntMapsEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func stringSlicesEqualUnordered(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	m := make(map[string]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		m[v]--
	}
	for _, v := range m {
		if v != 0 {
			return false
		}
	}
	return true
}

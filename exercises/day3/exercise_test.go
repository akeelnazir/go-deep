package main

import (
	"testing"
)

func TestExerciseGreet(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want string
	}{
		{"Alice", "Alice", "Hello, Alice!"},
		{"Bob", "Bob", "Hello, Bob!"},
		{"single char", "J", "Hello, J!"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseGreet(tt.arg); got != tt.want {
				t.Errorf("ExerciseGreet(%q) = %q, want %q", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseAddNumbers(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{"positive", 5, 3, 8},
		{"zero", 0, 0, 0},
		{"negative", -5, -3, -8},
		{"mixed", 10, -5, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseAddNumbers(tt.a, tt.b); got != tt.want {
				t.Errorf("ExerciseAddNumbers(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestExerciseDivideWithError(t *testing.T) {
	tests := []struct {
		name    string
		a       int
		b       int
		want    int
		wantErr bool
	}{
		{"valid", 10, 2, 5, false},
		{"zero divisor", 10, 0, 0, true},
		{"negative", -10, 2, -5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExerciseDivideWithError(tt.a, tt.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseDivideWithError(%d, %d) error = %v, wantErr %v", tt.a, tt.b, err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ExerciseDivideWithError(%d, %d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestExerciseSwapStrings(t *testing.T) {
	tests := []struct {
		name   string
		first  string
		second string
		want1  string
		want2  string
	}{
		{"simple", "Hello", "World", "World", "Hello"},
		{"names", "John", "Doe", "Doe", "John"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := ExerciseSwapStrings(tt.first, tt.second)
			if got1 != tt.want1 || got2 != tt.want2 {
				t.Errorf("ExerciseSwapStrings(%q, %q) = (%q, %q), want (%q, %q)", tt.first, tt.second, got1, got2, tt.want1, tt.want2)
			}
		})
	}
}

func TestExerciseSumVariadic(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want int
	}{
		{"single", []int{5}, 5},
		{"multiple", []int{1, 2, 3, 4, 5}, 15},
		{"empty", []int{}, 0},
		{"negative", []int{-1, -2, -3}, -6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseSumVariadic(tt.args...); got != tt.want {
				t.Errorf("ExerciseSumVariadic(%v) = %d, want %d", tt.args, got, tt.want)
			}
		})
	}
}

func TestExerciseFactorial(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"0", 0, 1},
		{"1", 1, 1},
		{"5", 5, 120},
		{"6", 6, 720},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseFactorial(tt.n); got != tt.want {
				t.Errorf("ExerciseFactorial(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestExerciseFibonacci(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"0", 0, 0},
		{"1", 1, 1},
		{"5", 5, 5},
		{"6", 6, 8},
		{"7", 7, 13},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExerciseFibonacci(tt.n); got != tt.want {
				t.Errorf("ExerciseFibonacci(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}

func TestExerciseCreateMultiplier(t *testing.T) {
	tests := []struct {
		name   string
		factor int
		arg    int
		want   int
	}{
		{"double", 2, 5, 10},
		{"triple", 3, 4, 12},
		{"zero", 0, 5, 0},
		{"negative", -2, 5, -10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fn := ExerciseCreateMultiplier(tt.factor)
			if got := fn(tt.arg); got != tt.want {
				t.Errorf("ExerciseCreateMultiplier(%d)(%d) = %d, want %d", tt.factor, tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseMapIntegers(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		fn   func(int) int
		want []int
	}{
		{
			"double",
			[]int{1, 2, 3, 4, 5},
			func(x int) int { return x * 2 },
			[]int{2, 4, 6, 8, 10},
		},
		{
			"square",
			[]int{1, 2, 3},
			func(x int) int { return x * x },
			[]int{1, 4, 9},
		},
		{
			"add one",
			[]int{0, 5, 10},
			func(x int) int { return x + 1 },
			[]int{1, 6, 11},
		},
		{
			"empty slice",
			[]int{},
			func(x int) int { return x * 2 },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseMapIntegers(tt.nums, tt.fn)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseMapIntegers(%v, fn) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestExerciseFilterIntegers(t *testing.T) {
	tests := []struct {
		name      string
		nums      []int
		predicate func(int) bool
		want      []int
	}{
		{
			"even numbers",
			[]int{1, 2, 3, 4, 5, 6},
			func(x int) bool { return x%2 == 0 },
			[]int{2, 4, 6},
		},
		{
			"greater than 3",
			[]int{1, 2, 3, 4, 5},
			func(x int) bool { return x > 3 },
			[]int{4, 5},
		},
		{
			"all match",
			[]int{2, 4, 6},
			func(x int) bool { return x%2 == 0 },
			[]int{2, 4, 6},
		},
		{
			"none match",
			[]int{1, 3, 5},
			func(x int) bool { return x%2 == 0 },
			[]int{},
		},
		{
			"empty slice",
			[]int{},
			func(x int) bool { return x > 0 },
			[]int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseFilterIntegers(tt.nums, tt.predicate)
			if !slicesEqual(got, tt.want) {
				t.Errorf("ExerciseFilterIntegers(%v, predicate) = %v, want %v", tt.nums, got, tt.want)
			}
		})
	}
}

func TestExerciseReduceIntegers(t *testing.T) {
	tests := []struct {
		name    string
		nums    []int
		initial int
		reducer func(int, int) int
		want    int
	}{
		{
			"sum",
			[]int{1, 2, 3, 4, 5},
			0,
			func(acc, x int) int { return acc + x },
			15,
		},
		{
			"product",
			[]int{1, 2, 3, 4},
			1,
			func(acc, x int) int { return acc * x },
			24,
		},
		{
			"max",
			[]int{3, 1, 4, 1, 5},
			0,
			func(acc, x int) int {
				if x > acc {
					return x
				}
				return acc
			},
			5,
		},
		{
			"empty slice",
			[]int{},
			42,
			func(acc, x int) int { return acc + x },
			42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseReduceIntegers(tt.nums, tt.initial, tt.reducer)
			if got != tt.want {
				t.Errorf("ExerciseReduceIntegers(%v, %d, reducer) = %d, want %d", tt.nums, tt.initial, got, tt.want)
			}
		})
	}
}

func TestExerciseCompose(t *testing.T) {
	tests := []struct {
		name string
		f    func(int) int
		g    func(int) int
		arg  int
		want int
	}{
		{
			"add one then double",
			func(x int) int { return x * 2 },
			func(x int) int { return x + 1 },
			5,
			12, // (5+1)*2 = 12
		},
		{
			"double then add one",
			func(x int) int { return x + 1 },
			func(x int) int { return x * 2 },
			5,
			11, // (5*2)+1 = 11
		},
		{
			"square then double",
			func(x int) int { return x * 2 },
			func(x int) int { return x * x },
			3,
			18, // (3*3)*2 = 18
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			composed := ExerciseCompose(tt.f, tt.g)
			got := composed(tt.arg)
			if got != tt.want {
				t.Errorf("ExerciseCompose(f, g)(%d) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseCurry(t *testing.T) {
	tests := []struct {
		name string
		fn   func(int, int) int
		a    int
		b    int
		want int
	}{
		{
			"add",
			func(x, y int) int { return x + y },
			5,
			3,
			8,
		},
		{
			"multiply",
			func(x, y int) int { return x * y },
			4,
			7,
			28,
		},
		{
			"subtract",
			func(x, y int) int { return x - y },
			10,
			3,
			7,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			curried := ExerciseCurry(tt.fn)
			got := curried(tt.a)(tt.b)
			if got != tt.want {
				t.Errorf("ExerciseCurry(fn)(%d)(%d) = %d, want %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestExercisePartialApply(t *testing.T) {
	tests := []struct {
		name  string
		fn    func(int, int) int
		first int
		arg   int
		want  int
	}{
		{
			"add with 5",
			func(x, y int) int { return x + y },
			5,
			3,
			8,
		},
		{
			"multiply by 2",
			func(x, y int) int { return x * y },
			2,
			7,
			14,
		},
		{
			"subtract from 10",
			func(x, y int) int { return x - y },
			10,
			3,
			7,
		},
		{
			"zero first",
			func(x, y int) int { return x + y },
			0,
			5,
			5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			partial := ExercisePartialApply(tt.fn, tt.first)
			got := partial(tt.arg)
			if got != tt.want {
				t.Errorf("ExercisePartialApply(fn, %d)(%d) = %d, want %d", tt.first, tt.arg, got, tt.want)
			}
		})
	}
}

func TestExerciseCreateDecorator(t *testing.T) {
	tests := []struct {
		name string
		fn   func(int) int
		arg  int
		want int
	}{
		{
			"double",
			func(x int) int { return x * 2 },
			5,
			10,
		},
		{
			"add one",
			func(x int) int { return x + 1 },
			10,
			11,
		},
		{
			"square",
			func(x int) int { return x * x },
			3,
			9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			decorator := ExerciseCreateDecorator()
			decorated := decorator(tt.fn)
			got := decorated(tt.arg)
			if got != tt.want {
				t.Errorf("decorator(fn)(%d) = %d, want %d", tt.arg, got, tt.want)
			}
		})
	}
}

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

// ============================================================================
// PART 2: STRUCTS AND METHODS TESTS
// ============================================================================

func TestExerciseCreatePerson(t *testing.T) {
	tests := []struct {
		name       string
		personName string
		age        int
		wantName   string
		wantAge    int
	}{
		{"Alice", "Alice", 30, "Alice", 30},
		{"Bob", "Bob", 25, "Bob", 25},
		{"Zero age", "Charlie", 0, "Charlie", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseCreatePerson(tt.personName, tt.age)
			if got.Name != tt.wantName || got.Age != tt.wantAge {
				t.Errorf("ExerciseCreatePerson(%q, %d) = {%q, %d}, want {%q, %d}",
					tt.personName, tt.age, got.Name, got.Age, tt.wantName, tt.wantAge)
			}
		})
	}
}

func TestExercisePersonGreet(t *testing.T) {
	tests := []struct {
		name string
		p    ExercisePerson
		want string
	}{
		{"Alice", ExercisePerson{Name: "Alice", Age: 30}, "Hello, I'm Alice and I'm 30 years old"},
		{"Bob", ExercisePerson{Name: "Bob", Age: 25}, "Hello, I'm Bob and I'm 25 years old"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.p.ExerciseGreet()
			if got != tt.want {
				t.Errorf("ExerciseGreet() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExercisePersonHaveBirthday(t *testing.T) {
	tests := []struct {
		name    string
		initial int
		want    int
	}{
		{"age 30 to 31", 30, 31},
		{"age 0 to 1", 0, 1},
		{"age 99 to 100", 99, 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ExercisePerson{Name: "Test", Age: tt.initial}
			p.ExerciseHaveBirthday()
			if p.Age != tt.want {
				t.Errorf("After HaveBirthday(), Age = %d, want %d", p.Age, tt.want)
			}
		})
	}
}

func TestExerciseCreateEmployee(t *testing.T) {
	tests := []struct {
		name       string
		personName string
		age        int
		id         int
		salary     float64
		wantName   string
		wantAge    int
		wantID     int
		wantSalary float64
	}{
		{"Alice", "Alice", 30, 101, 75000.0, "Alice", 30, 101, 75000.0},
		{"Bob", "Bob", 35, 102, 85000.5, "Bob", 35, 102, 85000.5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseCreateEmployee(tt.personName, tt.age, tt.id, tt.salary)
			if got.Name != tt.wantName || got.Age != tt.wantAge || got.EmployeeID != tt.wantID || got.Salary != tt.wantSalary {
				t.Errorf("ExerciseCreateEmployee(%q, %d, %d, %f) = {%q, %d, %d, %f}, want {%q, %d, %d, %f}",
					tt.personName, tt.age, tt.id, tt.salary,
					got.Name, got.Age, got.EmployeeID, got.Salary,
					tt.wantName, tt.wantAge, tt.wantID, tt.wantSalary)
			}
		})
	}
}

func TestExerciseEmployeeGetInfo(t *testing.T) {
	tests := []struct {
		name string
		emp  ExerciseEmployee
		want string
	}{
		{
			"Alice",
			ExerciseEmployee{
				ExercisePerson: ExercisePerson{Name: "Alice", Age: 30},
				EmployeeID:     101,
				Salary:         75000.0,
			},
			"Alice (ID: 101, Salary: $75000.00)",
		},
		{
			"Bob",
			ExerciseEmployee{
				ExercisePerson: ExercisePerson{Name: "Bob", Age: 35},
				EmployeeID:     102,
				Salary:         85000.5,
			},
			"Bob (ID: 102, Salary: $85000.50)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.emp.ExerciseGetInfo()
			if got != tt.want {
				t.Errorf("ExerciseGetInfo() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExercisePersonToJSON(t *testing.T) {
	tests := []struct {
		name string
		p    ExercisePerson
		want string
	}{
		{"Alice", ExercisePerson{Name: "Alice", Age: 30}, `{"Name":"Alice","Age":30}`},
		{"Bob", ExercisePerson{Name: "Bob", Age: 25}, `{"Name":"Bob","Age":25}`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExercisePersonToJSON(tt.p)
			if err != nil {
				t.Errorf("ExercisePersonToJSON() error = %v", err)
				return
			}
			if got != tt.want {
				t.Errorf("ExercisePersonToJSON() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExerciseJSONToPerson(t *testing.T) {
	tests := []struct {
		name     string
		jsonStr  string
		wantName string
		wantAge  int
		wantErr  bool
	}{
		{"Alice", `{"Name":"Alice","Age":30}`, "Alice", 30, false},
		{"Bob", `{"Name":"Bob","Age":25}`, "Bob", 25, false},
		{"invalid", `invalid json`, "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExerciseJSONToPerson(tt.jsonStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExerciseJSONToPerson() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && (got.Name != tt.wantName || got.Age != tt.wantAge) {
				t.Errorf("ExerciseJSONToPerson() = {%q, %d}, want {%q, %d}",
					got.Name, got.Age, tt.wantName, tt.wantAge)
			}
		})
	}
}

// ============================================================================
// PART 4: DESIGN PATTERNS TESTS
// ============================================================================

type MockLogger struct {
	lastMsg string
}

func (ml *MockLogger) Log(msg string) string {
	ml.lastMsg = msg
	return "[LOG] " + msg
}

func TestExerciseDependencyInjection(t *testing.T) {
	tests := []struct {
		name string
		task string
		want string
	}{
		{"simple task", "build", "[LOG] Working on: build"},
		{"complex task", "deploy", "[LOG] Working on: deploy"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logger := &MockLogger{}
			service := ExerciseCreateService(logger)
			if service == nil {
				t.Errorf("ExerciseCreateService() returned nil")
				return
			}
			got := service.DoWork(tt.task)
			if got != tt.want {
				t.Errorf("DoWork(%q) = %q, want %q", tt.task, got, tt.want)
			}
		})
	}
}

func TestExerciseBuilderPattern(t *testing.T) {
	tests := []struct {
		name  string
		query string
		limit int
		want  string
	}{
		{"simple", "SELECT * FROM users", 10, "SELECT * FROM users LIMIT 10"},
		{"no limit", "SELECT * FROM products", 0, "SELECT * FROM products"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			builder := ExerciseNewQueryBuilder(tt.query)
			if builder == nil {
				t.Errorf("ExerciseNewQueryBuilder() returned nil")
				return
			}
			if tt.limit > 0 {
				builder = builder.ExerciseLimit(tt.limit)
			}
			got := builder.ExerciseBuild()
			if got != tt.want {
				t.Errorf("Build() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestExerciseFunctionalOptions(t *testing.T) {
	tests := []struct {
		name     string
		opts     []ExerciseConfigOption
		wantPort int
		wantHost string
	}{
		{
			"default",
			[]ExerciseConfigOption{},
			8080,
			"localhost",
		},
		{
			"custom port",
			[]ExerciseConfigOption{ExerciseWithPort(9000)},
			9000,
			"localhost",
		},
		{
			"custom host",
			[]ExerciseConfigOption{ExerciseWithHost("0.0.0.0")},
			8080,
			"0.0.0.0",
		},
		{
			"both custom",
			[]ExerciseConfigOption{ExerciseWithPort(3000), ExerciseWithHost("127.0.0.1")},
			3000,
			"127.0.0.1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ExerciseCreateConfig(tt.opts...)
			if got == nil {
				t.Errorf("ExerciseCreateConfig() returned nil")
				return
			}
			if got.Port != tt.wantPort || got.Host != tt.wantHost {
				t.Errorf("Config = {Port: %d, Host: %q}, want {Port: %d, Host: %q}",
					got.Port, got.Host, tt.wantPort, tt.wantHost)
			}
		})
	}
}

func TestExerciseSingletonPattern(t *testing.T) {
	s1 := ExerciseGetSingleton()
	s2 := ExerciseGetSingleton()

	if s1 == nil || s2 == nil {
		t.Errorf("ExerciseGetSingleton() returned nil")
		return
	}

	if s1 != s2 {
		t.Errorf("Singleton instances are not the same: %p != %p", s1, s2)
	}
}

func TestExerciseFactoryPattern(t *testing.T) {
	tests := []struct {
		name      string
		shapeType string
		args      []float64
		wantArea  float64
	}{
		{"circle", "circle", []float64{5}, 78.53981633974483}, // π * 5^2
		{"rectangle", "rectangle", []float64{4, 5}, 20},       // 4 * 5
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			shape := ExerciseCreateShape(tt.shapeType, tt.args...)
			if shape == nil {
				t.Errorf("ExerciseCreateShape(%q) returned nil", tt.shapeType)
				return
			}
			got := shape.Area()
			if got != tt.wantArea {
				t.Errorf("Area() = %f, want %f", got, tt.wantArea)
			}
		})
	}
}

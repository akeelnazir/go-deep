package main

import (
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// Section 1: Basic Function Declaration
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// Section 2: Function with Return Value
func add(a, b int) int {
	return a + b
}

// Section 3: Function with Multiple Return Values
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

// Section 4: Named Return Values
func swap(x, y string) (first, second string) {
	first = y
	second = x
	return
}

// Section 5: Variadic Function
func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

// Section 6: Variadic with Mixed Parameters
func printStrings(prefix string, items ...string) {
	fmt.Printf("%s: ", prefix)
	for i, item := range items {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(item)
	}
	fmt.Println()
}

// Section 7: Function with Defer
func deferExample() {
	defer fmt.Println("3. Cleanup (deferred)")
	fmt.Println("1. Start")
	defer fmt.Println("2. More cleanup (deferred)")
	fmt.Println("1.5. Middle")
}

// Section 8: Defer with File Operations (Simulated)
func deferWithResource() {
	fmt.Println("Opening resource...")
	defer fmt.Println("Closing resource...")
	fmt.Println("Using resource...")
}

// Section 9: Higher-Order Function - Takes a Function as Parameter
func applyOperation(a, b int, operation func(int, int) int) int {
	return operation(a, b)
}

// Section 10: Function Returning a Function
func makeMultiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

// Section 11: Closure Example
func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

// Section 12: Recursive Function - Factorial
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// Section 13: Recursive Function - Fibonacci
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// Higher-Order Function Helpers

// mapIntegers applies a function to each element in a slice
func mapIntegers(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

// filterIntegers returns elements that satisfy the predicate
func filterIntegers(nums []int, predicate func(int) bool) []int {
	var result []int
	for _, n := range nums {
		if predicate(n) {
			result = append(result, n)
		}
	}
	return result
}

// reduceIntegers combines all elements using a reducer function
func reduceIntegers(nums []int, initial int, reducer func(int, int) int) int {
	result := initial
	for _, n := range nums {
		result = reducer(result, n)
	}
	return result
}

// composeFunc applies g first, then f to the result
func composeFunc(f, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}

// pipe applies functions left to right
func pipe(x int, fns ...func(int) int) int {
	result := x
	for _, fn := range fns {
		result = fn(result)
	}
	return result
}

// curriedAdd returns a curried version of addition
func curriedAdd(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

// curriedMultiply returns a curried version of multiplication
func curriedMultiply(a int) func(int) int {
	return func(b int) int {
		return a * b
	}
}

// partialMultiply fixes the first argument of multiplication
func partialMultiply(factor int) func(int) int {
	return func(x int) int {
		return factor * x
	}
}

// withLogging decorates a function with logging
func withLogging(fn func(int) int) func(int) int {
	return func(x int) int {
		fmt.Printf("Calling function with argument: %d\n", x)
		result := fn(x)
		fmt.Printf("Function returned: %d\n", result)
		return result
	}
}

// withTiming decorates a function with execution timing
func withTiming(fn func(int) int) func(int) int {
	return func(x int) int {
		start := time.Now()
		result := fn(x)
		elapsed := time.Since(start)
		fmt.Printf("Execution time: %v\n", elapsed)
		return result
	}
}

// withCache decorates a function with caching
func withCache(fn func(int) int) func(int) int {
	cache := make(map[int]int)
	return func(x int) int {
		if val, exists := cache[x]; exists {
			fmt.Println("Cache hit!")
			return val
		}
		fmt.Println("Cache miss, computing...")
		result := fn(x)
		cache[x] = result
		return result
	}
}

// chainDecorators applies multiple decorators to a function
func chainDecorators(fn func(int) int, decorators ...func(func(int) int) func(int) int) func(int) int {
	result := fn
	for _, decorator := range decorators {
		result = decorator(result)
	}
	return result
}

// pipeline processes data through a series of transformations
func pipeline(data []int, transforms ...func([]int) []int) []int {
	result := data
	for _, transform := range transforms {
		result = transform(result)
	}
	return result
}

// SortStrategy is a type for sorting strategies
type SortStrategy func([]int) []int

// bubbleSort implements bubble sort
func bubbleSort(nums []int) []int {
	result := make([]int, len(nums))
	copy(result, nums)
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(result)-1-i; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
			}
		}
	}
	return result
}

// quickSort implements quick sort
func quickSort(nums []int) []int {
	if len(nums) <= 1 {
		result := make([]int, len(nums))
		copy(result, nums)
		return result
	}
	pivot := nums[len(nums)/2]
	var left, middle, right []int
	for _, n := range nums {
		if n < pivot {
			left = append(left, n)
		} else if n == pivot {
			middle = append(middle, n)
		} else {
			right = append(right, n)
		}
	}
	result := append(quickSort(left), middle...)
	result = append(result, quickSort(right)...)
	return result
}

// sortData uses a strategy to sort data
func sortData(data []int, strategy SortStrategy) []int {
	return strategy(data)
}

// Section 14: Function with Multiple Named Returns
func getCoordinates() (x, y int) {
	x = 10
	y = 20
	return
}

// ============================================================================
// PART 2: STRUCTS AND METHODS
// ============================================================================

// Section 33: Struct Definition
type Person struct {
	Name string
	Age  int
	City string
}

// Section 34: Method with Value Receiver (read-only)
func (p Person) Greet() string {
	return "Hello, my name is " + p.Name
}

// Section 35: Method with Pointer Receiver (can modify)
func (p *Person) HaveBirthday() {
	p.Age++
}

// Section 36: Struct Embedding and Composition
type Employee struct {
	Person
	EmployeeID int
	Salary     float64
}

// Section 37: Method on Embedded Type
func (e Employee) GetInfo() string {
	return fmt.Sprintf("%s (ID: %d, Salary: $%.2f)", e.Name, e.EmployeeID, e.Salary)
}

// Section 38: Struct with Tags for JSON
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email,omitempty"`
	Age   int    `json:"age"`
}

// ============================================================================
// PART 4: DESIGN PATTERNS
// ============================================================================

// Section 39: Dependency Injection Pattern
type Logger interface {
	Log(msg string)
}

type ConsoleLogger struct{}

func (cl ConsoleLogger) Log(msg string) {
	fmt.Printf("[LOG] %s\n", msg)
}

type UserService struct {
	logger Logger
}

func NewUserService(logger Logger) *UserService {
	return &UserService{logger: logger}
}

func (us *UserService) CreateUser(name string) {
	us.logger.Log("Creating user: " + name)
}

// Section 40: Builder Pattern
type QueryBuilder struct {
	query  string
	args   []interface{}
	limit  int
	offset int
}

func NewQueryBuilder(query string) *QueryBuilder {
	return &QueryBuilder{query: query}
}

func (qb *QueryBuilder) Where(condition string, args ...interface{}) *QueryBuilder {
	qb.query += " WHERE " + condition
	qb.args = append(qb.args, args...)
	return qb
}

func (qb *QueryBuilder) Limit(limit int) *QueryBuilder {
	qb.limit = limit
	return qb
}

func (qb *QueryBuilder) Offset(offset int) *QueryBuilder {
	qb.offset = offset
	return qb
}

func (qb *QueryBuilder) Build() (string, []interface{}) {
	if qb.limit > 0 {
		qb.query += fmt.Sprintf(" LIMIT %d", qb.limit)
	}
	if qb.offset > 0 {
		qb.query += fmt.Sprintf(" OFFSET %d", qb.offset)
	}
	return qb.query, qb.args
}

// Section 41: Functional Options Pattern
type ServerConfig struct {
	Port         int
	Host         string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type Option func(*ServerConfig)

func WithPort(port int) Option {
	return func(cfg *ServerConfig) {
		cfg.Port = port
	}
}

func WithHost(host string) Option {
	return func(cfg *ServerConfig) {
		cfg.Host = host
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(cfg *ServerConfig) {
		cfg.ReadTimeout = timeout
	}
}

func NewServer(opts ...Option) *ServerConfig {
	cfg := &ServerConfig{
		Port:         8080,
		Host:         "localhost",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

// Section 42: Singleton Pattern
var (
	instance *Singleton
	once     sync.Once
)

type Singleton struct {
	value string
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{value: "singleton"}
	})
	return instance
}

// Section 43: Factory Pattern
type Database interface {
	Query(sql string) error
}

type PostgresDB struct{}

func (p *PostgresDB) Query(sql string) error {
	fmt.Printf("Executing on PostgreSQL: %s\n", sql)
	return nil
}

type MySQLDB struct{}

func (m *MySQLDB) Query(sql string) error {
	fmt.Printf("Executing on MySQL: %s\n", sql)
	return nil
}

func NewDatabase(dbType string) Database {
	switch dbType {
	case "postgres":
		return &PostgresDB{}
	case "mysql":
		return &MySQLDB{}
	default:
		return nil
	}
}

// Section 44: Observer Pattern
type Observer interface {
	Update(event string)
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Notify(event string) {
	for _, o := range s.observers {
		o.Update(event)
	}
}

type ConcreteObserver struct {
	name string
}

func (co *ConcreteObserver) Update(event string) {
	fmt.Printf("%s received event: %s\n", co.name, event)
}

// Section 45: Strategy Pattern
type PaymentStrategy interface {
	Pay(amount float64) error
}

type CreditCardPayment struct {
	cardNumber string
}

func (c *CreditCardPayment) Pay(amount float64) error {
	fmt.Printf("Processing credit card payment of $%.2f\n", amount)
	return nil
}

type PayPalPayment struct {
	email string
}

func (p *PayPalPayment) Pay(amount float64) error {
	fmt.Printf("Processing PayPal payment of $%.2f\n", amount)
	return nil
}

type Order struct {
	strategy PaymentStrategy
}

func (o *Order) Checkout(amount float64) error {
	return o.strategy.Pay(amount)
}

// Section 46: Repository Pattern
type UserRepository interface {
	GetByID(id int) (*User, error)
	Save(user *User) error
	Delete(id int) error
}

type InMemoryUserRepository struct {
	users map[int]*User
}

func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{
		users: make(map[int]*User),
	}
}

func (r *InMemoryUserRepository) GetByID(id int) (*User, error) {
	if user, exists := r.users[id]; exists {
		return user, nil
	}
	return nil, fmt.Errorf("user not found")
}

func (r *InMemoryUserRepository) Save(user *User) error {
	r.users[user.ID] = user
	return nil
}

func (r *InMemoryUserRepository) Delete(id int) error {
	delete(r.users, id)
	return nil
}

// Section 15: Practical Example - Calculator with Error Handling
func calculate(operation string, a, b int) (int, error) {
	switch operation {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("unknown operation: %s", operation)
	}
}

func main() {
	fmt.Println("=== Day 3: Functions ===")

	// Section 1: Basic Function Call
	fmt.Println("\n--- Basic Function Call ---")
	greet("Alice")
	greet("Bob")

	// Section 2: Function with Return Value
	fmt.Println("\n--- Function with Return Value ---")
	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	result2 := add(10, 20)
	fmt.Printf("10 + 20 = %d\n", result2)

	// Section 3: Multiple Return Values
	fmt.Println("\n--- Multiple Return Values ---")
	quotient, err := divide(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", quotient)
	}

	quotient2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Printf("Error: %v\n", err2)
	} else {
		fmt.Printf("10 / 0 = %d\n", quotient2)
	}

	// Section 4: Named Return Values
	fmt.Println("\n--- Named Return Values ---")
	first, second := swap("Hello", "World")
	fmt.Printf("Swapped: %s, %s\n", first, second)

	// Section 5: Variadic Functions
	fmt.Println("\n--- Variadic Functions ---")
	fmt.Printf("Sum of 1, 2, 3: %d\n", sum(1, 2, 3))
	fmt.Printf("Sum of 10, 20, 30, 40: %d\n", sum(10, 20, 30, 40))
	fmt.Printf("Sum of no args: %d\n", sum())

	// Passing a slice to variadic function
	numbers := []int{5, 10, 15, 20}
	fmt.Printf("Sum of slice: %d\n", sum(numbers...))

	// Section 6: Variadic with Mixed Parameters
	fmt.Println("\n--- Variadic with Mixed Parameters ---")
	printStrings("Colors", "Red", "Green", "Blue")
	printStrings("Numbers", "One", "Two", "Three", "Four")

	// Section 7: Defer Statement
	fmt.Println("\n--- Defer Statement ---")
	deferExample()

	fmt.Println("\n--- Defer with Resource ---")
	deferWithResource()

	// Section 8: Multiple Defers (LIFO order)
	fmt.Println("\n--- Multiple Defers (LIFO) ---")
	func() {
		defer fmt.Println("First defer (executed last)")
		defer fmt.Println("Second defer (executed second)")
		defer fmt.Println("Third defer (executed first)")
		fmt.Println("Main function body")
	}()

	// Section 9: Function Values
	fmt.Println("\n--- Function Values ---")
	add2 := func(a, b int) int {
		return a + b
	}
	fmt.Printf("Anonymous function: 5 + 3 = %d\n", add2(5, 3))

	multiply := func(a, b int) int {
		return a * b
	}
	fmt.Printf("Anonymous function: 4 * 7 = %d\n", multiply(4, 7))

	// Section 10: Higher-Order Functions
	fmt.Println("\n--- Higher-Order Functions ---")
	addFunc := func(a, b int) int { return a + b }
	subtractFunc := func(a, b int) int { return a - b }

	fmt.Printf("Apply add: 10 + 5 = %d\n", applyOperation(10, 5, addFunc))
	fmt.Printf("Apply subtract: 10 - 5 = %d\n", applyOperation(10, 5, subtractFunc))

	// Section 11: Function Returning a Function
	fmt.Println("\n--- Function Returning a Function ---")
	double := makeMultiplier(2)
	triple := makeMultiplier(3)

	fmt.Printf("Double 5: %d\n", double(5))
	fmt.Printf("Triple 5: %d\n", triple(5))
	fmt.Printf("Double 10: %d\n", double(10))

	// Section 12: Closures
	fmt.Println("\n--- Closures ---")
	count1 := counter()
	count2 := counter()

	fmt.Printf("count1: %d\n", count1())
	fmt.Printf("count1: %d\n", count1())
	fmt.Printf("count1: %d\n", count1())

	fmt.Printf("count2: %d\n", count2())
	fmt.Printf("count2: %d\n", count2())

	// Section 13: Recursion - Factorial
	fmt.Println("\n--- Recursion: Factorial ---")
	for i := 0; i <= 5; i++ {
		fmt.Printf("Factorial of %d: %d\n", i, factorial(i))
	}

	// Section 14: Recursion - Fibonacci
	fmt.Println("\n--- Recursion: Fibonacci ---")
	for i := 0; i <= 7; i++ {
		fmt.Printf("Fibonacci(%d): %d\n", i, fibonacci(i))
	}

	// Section 15: Multiple Named Returns
	fmt.Println("\n--- Multiple Named Returns ---")
	x, y := getCoordinates()
	fmt.Printf("Coordinates: x=%d, y=%d\n", x, y)

	// Section 16: Practical Example - Calculator
	fmt.Println("\n--- Practical Example: Calculator ---")
	operations := []string{"add", "subtract", "multiply", "divide"}
	a, b := 20, 5

	for _, op := range operations {
		result, err := calculate(op, a, b)
		if err != nil {
			fmt.Printf("%s: Error - %v\n", op, err)
		} else {
			fmt.Printf("%s: %d %s %d = %d\n", op, a, op, b, result)
		}
	}

	// Section 17: Practical Example - Function Composition
	fmt.Println("\n--- Practical Example: Function Composition ---")
	addOne := func(x int) int { return x + 1 }
	double2 := func(x int) int { return x * 2 }
	compose := func(f, g func(int) int, x int) int {
		return f(g(x))
	}

	value := 5
	fmt.Printf("Value: %d\n", value)
	fmt.Printf("Double then add one: %d\n", compose(addOne, double2, value))
	fmt.Printf("Add one then double: %d\n", compose(double2, addOne, value))

	// Section 18: Defer with Panic Recovery (Preview)
	fmt.Println("\n--- Defer with Panic Recovery (Preview) ---")
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Recovered from panic: %v\n", r)
			}
		}()
		fmt.Println("Before panic")
		panic("Something went wrong!")
	}()
	fmt.Println("After panic recovery")

	// Section 19: Higher-Order Functions - Map Pattern
	fmt.Println("\n--- Higher-Order Functions: Map ---")
	mapNumbers := []int{1, 2, 3, 4, 5}
	doubled := mapIntegers(mapNumbers, func(x int) int { return x * 2 })
	fmt.Printf("Original: %v\n", mapNumbers)
	fmt.Printf("Doubled: %v\n", doubled)

	squared := mapIntegers(mapNumbers, func(x int) int { return x * x })
	fmt.Printf("Squared: %v\n", squared)

	// Section 20: Higher-Order Functions - Filter Pattern
	fmt.Println("\n--- Higher-Order Functions: Filter ---")
	filterNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filterIntegers(filterNumbers, func(x int) bool { return x%2 == 0 })
	fmt.Printf("Original: %v\n", filterNumbers)
	fmt.Printf("Even numbers: %v\n", evens)

	greaterThan5 := filterIntegers(filterNumbers, func(x int) bool { return x > 5 })
	fmt.Printf("Greater than 5: %v\n", greaterThan5)

	// Section 21: Higher-Order Functions - Reduce Pattern
	fmt.Println("\n--- Higher-Order Functions: Reduce ---")
	reduceNumbers := []int{1, 2, 3, 4, 5}
	sum := reduceIntegers(reduceNumbers, 0, func(acc, x int) int { return acc + x })
	fmt.Printf("Numbers: %v\n", reduceNumbers)
	fmt.Printf("Sum: %d\n", sum)

	product := reduceIntegers(reduceNumbers, 1, func(acc, x int) int { return acc * x })
	fmt.Printf("Product: %d\n", product)

	// Section 22: Function Composition
	fmt.Println("\n--- Function Composition ---")
	addOneFunc := func(x int) int { return x + 1 }
	doubleFunc := func(x int) int { return x * 2 }
	squareFunc := func(x int) int { return x * x }

	composedFunc := composeFunc(doubleFunc, addOneFunc)
	fmt.Printf("Compose(double, addOne)(5) = %d (should be (5+1)*2 = 12)\n", composedFunc(5))

	composedFunc2 := composeFunc(addOneFunc, doubleFunc)
	fmt.Printf("Compose(addOne, double)(5) = %d (should be (5*2)+1 = 11)\n", composedFunc2(5))

	// Section 23: Pipe - Left-to-Right Composition
	fmt.Println("\n--- Pipe (Left-to-Right Composition) ---")
	pipeResult := pipe(5, addOneFunc, doubleFunc, squareFunc)
	fmt.Printf("Pipe(5, addOne, double, square) = %d (should be ((5+1)*2)^2 = 144)\n", pipeResult)

	// Section 24: Currying
	fmt.Println("\n--- Currying ---")
	curriedAddFunc := curriedAdd(10)
	fmt.Printf("curriedAdd(10)(5) = %d\n", curriedAddFunc(5))
	fmt.Printf("curriedAdd(10)(20) = %d\n", curriedAddFunc(20))

	curriedMultiplyFunc := curriedMultiply(3)
	fmt.Printf("curriedMultiply(3)(4) = %d\n", curriedMultiplyFunc(4))
	fmt.Printf("curriedMultiply(3)(7) = %d\n", curriedMultiplyFunc(7))

	// Section 25: Partial Application
	fmt.Println("\n--- Partial Application ---")
	doublePartial := partialMultiply(2)
	triplePartial := partialMultiply(3)

	fmt.Printf("partialMultiply(2)(5) = %d\n", doublePartial(5))
	fmt.Printf("partialMultiply(3)(5) = %d\n", triplePartial(5))

	// Section 26: Decorator Pattern - Logging
	fmt.Println("\n--- Decorator Pattern: Logging ---")
	simpleDouble := func(x int) int { return x * 2 }
	loggedDouble := withLogging(simpleDouble)
	fmt.Printf("Result: %d\n", loggedDouble(5))

	// Section 27: Decorator Pattern - Timing
	fmt.Println("\n--- Decorator Pattern: Timing ---")
	slowFunction := func(x int) int {
		time.Sleep(10 * time.Millisecond)
		return x * x
	}
	timedFunction := withTiming(slowFunction)
	fmt.Printf("Result: %d\n", timedFunction(4))

	// Section 28: Decorator Pattern - Caching
	fmt.Println("\n--- Decorator Pattern: Caching ---")
	expensiveFunc := func(n int) int {
		fmt.Printf("Computing fibonacci(%d)...\n", n)
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}
	cachedFunc := withCache(expensiveFunc)
	fmt.Printf("First call - fibonacci(5) = %d\n", cachedFunc(5))
	fmt.Printf("Second call - fibonacci(5) = %d (should use cache)\n", cachedFunc(5))

	// Section 29: Chaining Decorators
	fmt.Println("\n--- Chaining Decorators ---")
	baseFunc := func(x int) int { return x + 10 }
	chainedFunc := chainDecorators(baseFunc, withLogging, withTiming)
	fmt.Printf("Result: %d\n", chainedFunc(5))

	// Section 30: Pipeline Processing
	fmt.Println("\n--- Pipeline Processing ---")
	pipelineNumbers := []int{1, 2, 3, 4, 5}
	doubleTransform := func(nums []int) []int {
		return mapIntegers(nums, func(x int) int { return x * 2 })
	}
	filterEvenTransform := func(nums []int) []int {
		return filterIntegers(nums, func(x int) bool { return x%2 == 0 })
	}
	pipelineResult := pipeline(pipelineNumbers, doubleTransform, filterEvenTransform)
	fmt.Printf("Original: %v\n", pipelineNumbers)
	fmt.Printf("After pipeline (double, filter even): %v\n", pipelineResult)

	// Section 31: Strategy Pattern with Higher-Order Functions
	fmt.Println("\n--- Strategy Pattern ---")
	strategyNumbers := []int{5, 2, 8, 1, 9}
	fmt.Printf("Original: %v\n", strategyNumbers)
	fmt.Printf("Using bubbleSort strategy: %v\n", sortData(strategyNumbers, bubbleSort))
	fmt.Printf("Using quickSort strategy: %v\n", sortData(strategyNumbers, quickSort))

	// Section 32: Practical Example - Data Transformation Pipeline
	fmt.Println("\n--- Practical Example: Data Transformation ---")
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	addTwo := func(x int) int { return x + 2 }
	isEven := func(x int) bool { return x%2 == 0 }
	transformed := pipeline(data,
		func(nums []int) []int { return mapIntegers(nums, addTwo) },
		func(nums []int) []int { return filterIntegers(nums, isEven) },
	)
	fmt.Printf("Original: %v\n", data)
	fmt.Printf("After adding 2 and filtering even: %v\n", transformed)

	// Section 33: Struct Creation and Methods
	fmt.Println("\n=== Part 2: Structs and Methods ===")
	fmt.Println("\n--- Struct Creation ---")
	person := Person{
		Name: "Alice",
		Age:  30,
		City: "New York",
	}
	fmt.Printf("Person: %+v\n", person)
	fmt.Println(person.Greet())

	// Section 34: Value vs Pointer Receivers
	fmt.Println("\n--- Value vs Pointer Receivers ---")
	fmt.Printf("Before birthday: Age = %d\n", person.Age)
	person.HaveBirthday()
	fmt.Printf("After birthday: Age = %d\n", person.Age)

	// Section 35: Struct Embedding
	fmt.Println("\n--- Struct Embedding and Composition ---")
	emp := Employee{
		Person: Person{
			Name: "Bob",
			Age:  35,
			City: "San Francisco",
		},
		EmployeeID: 101,
		Salary:     75000,
	}
	fmt.Println(emp.GetInfo())
	fmt.Println(emp.Greet())

	// Section 36: JSON Serialization
	fmt.Println("\n--- JSON Serialization ---")
	user := User{
		ID:    1,
		Name:  "Charlie",
		Email: "charlie@example.com",
		Age:   28,
	}
	jsonData, _ := json.Marshal(user)
	fmt.Printf("JSON: %s\n", string(jsonData))

	jsonString := `{"id":2,"name":"Diana","email":"diana@example.com","age":32}`
	var user2 User
	json.Unmarshal([]byte(jsonString), &user2)
	fmt.Printf("Unmarshaled: %+v\n", user2)

	// Section 37: Design Patterns
	fmt.Println("\n=== Part 4: Design Patterns ===")

	// Dependency Injection
	fmt.Println("\n--- Dependency Injection Pattern ---")
	logger := ConsoleLogger{}
	userService := NewUserService(logger)
	userService.CreateUser("Eve")

	// Builder Pattern
	fmt.Println("\n--- Builder Pattern ---")
	query, args := NewQueryBuilder("SELECT * FROM users").
		Where("age > ?", 18).
		Limit(10).
		Offset(0).
		Build()
	fmt.Printf("Query: %s\n", query)
	fmt.Printf("Args: %v\n", args)

	// Functional Options Pattern
	fmt.Println("\n--- Functional Options Pattern ---")
	server := NewServer(
		WithPort(9000),
		WithHost("0.0.0.0"),
		WithReadTimeout(10*time.Second),
	)
	fmt.Printf("Server Config: Port=%d, Host=%s, ReadTimeout=%v\n", server.Port, server.Host, server.ReadTimeout)

	// Singleton Pattern
	fmt.Println("\n--- Singleton Pattern ---")
	s1 := GetInstance()
	s2 := GetInstance()
	fmt.Printf("Same instance: %v\n", s1 == s2)

	// Factory Pattern
	fmt.Println("\n--- Factory Pattern ---")
	pgDB := NewDatabase("postgres")
	pgDB.Query("SELECT * FROM users")
	mysqlDB := NewDatabase("mysql")
	mysqlDB.Query("SELECT * FROM users")

	// Observer Pattern
	fmt.Println("\n--- Observer Pattern ---")
	subject := &Subject{}
	observer1 := &ConcreteObserver{name: "Observer1"}
	observer2 := &ConcreteObserver{name: "Observer2"}
	subject.Attach(observer1)
	subject.Attach(observer2)
	subject.Notify("Event occurred!")

	// Strategy Pattern
	fmt.Println("\n--- Strategy Pattern ---")
	order := &Order{strategy: &CreditCardPayment{cardNumber: "1234-5678"}}
	order.Checkout(99.99)
	order.strategy = &PayPalPayment{email: "user@example.com"}
	order.Checkout(49.99)

	// Repository Pattern
	fmt.Println("\n--- Repository Pattern ---")
	repo := NewInMemoryUserRepository()
	repo.Save(&User{ID: 1, Name: "Frank", Age: 40})
	retrievedUser, _ := repo.GetByID(1)
	fmt.Printf("Retrieved user: %+v\n", retrievedUser)

	fmt.Println("\n=== Day 3 Complete ===")
	fmt.Println("Next: Learn about arrays, slices, and maps on Day 4.")
}

package main

import "fmt"

// Optimized using memoization
func fibonacci_memoize(n int, memoizationTable map[int]int) int {
	// Base case: the 1st and 2nd value in the fibonacci sequence, which are 1
	if n == 0 || n == 1 {
		return 1
	}

	// If the nth fibonacci doesn't exist in our memo table, generate it
	if _, exists := memoizationTable[n]; !exists {
		memoizationTable[n] = fibonacci_memoize(n-1, memoizationTable) + fibonacci_memoize(n-2, memoizationTable)
	}

	return memoizationTable[n]
}

// The fibonacci sequence is 1, 1, and then the sum of the previous two elements
func fibonacci(n int) int {
	memorizationTable := make(map[int]int)
	return fibonacci_memoize(n, memorizationTable)
}

func main() {
	// Basic driver code that gets the first 20 fibonacci numbers
	for n := 0; n < 20; n++ {
		fmt.Printf("Fibonacci # at idx %d: %d\n", n+1, fibonacci(n))
	}
}

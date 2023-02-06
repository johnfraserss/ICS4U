package main

import "fmt"

// The fibonacci sequence is 1, 1, and then the sum of the previous two elements
func fibonacci(n int) int {
	// Base case: the 1st and 2nd value in the fibonacci sequence, which are 1
	if n == 0 || n == 1 {
		return 1
	}

	// The nth value in the fibonacci sequence is the sum of the previous two numbers,
	// which are both calculated recursively
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	n := 0
	fmt.Printf("Fibonacci # at idx %d: %d\n", n, fibonacci(n))
}

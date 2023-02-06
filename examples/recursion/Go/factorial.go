/**
 * Recursion
 *
 * Recursion is essentially just calling a function in itself until a certain
 * value or expected response is returned.
 *
 * In this example, we'll use factorials as an example.
 * Factorials (in case you weren't aware already) are mathematically represented by `N!`
 * where N is a positive integer, and ! is the factorial operator.
 *
 * Take `4!`, this is equal to 4 x 3 x 2 x 1, or 24
 *
 * So how do we implement this in code?
 */
package main

import "fmt"

// recursiveFactorial returns the factorial of n.
func recursiveFactorial(n int) int {
	// Base case: n <= 1
	// 1 is the last number that we should be multiplying with, so just return n back
	if n <= 1 {
		return n
	}

	// Return n * (n - 1)!, where (n - 1)! is just calling the func again but with n - 1
	return n * recursiveFactorial(n-1)
}

func main() {
	/**
	 * Naive approach
	 *
	 * In this approach, we're given N, and we store the factorial of N in a variable
	 * The factorial of N is manually calculated by typing everything out.
	 *
	 * As you may have predicted, this won't hold true for all integers N, if N was 5, then
	 * `N!` wouldn't be equivalent to 4*3*2*1
	 */
	n := 4
	nFactorial := 4 * 3 * 2 * 1
	fmt.Printf("Factorial of %d (naive): %d\n", n, nFactorial)

	/**
	 * Approach with recursion
	 *
	 * With recursion, we're also given a positive integer N.
	 * You may have noticed that there's a pattern, for any N,
	 * `N!` is equivalent to `N * N-1 * N-2 * ... * 1`.
	 *
	 * The pattern above holds true for any integer N, so, based off of that,
	 * we can make a function that recursively calculates the factorial of N
	 */
	fmt.Printf("Factorial of %d (recursive): %d\n", n, recursiveFactorial(n))
}

// Still confused? There's a lengthy explanation below

/**
 * Let's take the driver code/example there, finding the factorial of 4.
 *
 * We start off by passing 6 into `recursiveFactorial`, 6 !== 1,  so we pass
 * 4-1, which is 3, into recursiveFactorial, as well as the product, which is equal to 1*4, or just 4.
 * At N = 4:
 *    product = 4, N = 3
 *
 * `recursiveFactorial` is called again with the above parameters.
 * N !== 1, which means we pass N-1, which is 2, into `recursiveFactorial` with the product of
 * 4*3, which is 12.
 * At N = 3:
 *    product = 12, N = 2
 *
 * `recursiveFactorial` is called again with the above parameters.
 * N !== 1, which means we pass N-1, which is 1, into `recursiveFactorial` with the product of
 * 12*2, which is 24.
 * At N = 2:
 *    product = 24, N = 1
 *
 * `recursiveFactorial` is called again with the above parameters.
 * N === 1, which means we return the product, 24
 *
 * 24 === 4!
 */

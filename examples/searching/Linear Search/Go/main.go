/**
 * Linear search (a.k.a. the easier one)
 *
 * Time complexity: O(N)
 *   - slower at large scales
 *
 * Linear search is pretty self-descriptive.
 * Given an array, and an element X, loop through the entire
 * array in order until you find X.
 *
 * The benefit of linear search is that the array doesn't have to
 * be sorted since you're looping through the entire array.
 */

package main

import "fmt"

// linearSearch returns whether element x is in array arr.
func linearSearch(arr []int, x int) int {
	// Using range, each element of the array is stored as "item" and compared to x.
	// If it equals x, we set found to true and exit the loop.
	// If none of them equal x, found remains false.
	for i, item := range arr {
		if item == x {
			// Return position of the item in the array
			return i
		}
	}

	// Could not find item
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 5
	pos := linearSearch(arr, x)

	if pos != -1 {
		fmt.Printf("%v was found in the array at pos %d.\n", x, pos)
	} else {
		fmt.Printf("%v was not found in the array\n", x)
	}
}

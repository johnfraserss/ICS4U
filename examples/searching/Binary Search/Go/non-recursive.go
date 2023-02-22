/**
 * Binary Search
 *
 * Time complexity: O(logN) - significantly faster than linear search at scale
 *
 * Given an element X in an array, binary search will find whether or not X
 * exists in said array.
 *
 * NOTE: this assume that the array is sorted
 */

package main

import "fmt"

// Iteratively search sorted array arr for element x.
// Returns whether element x is within array arr or not.
func nonRecursiveSearch(arr []int, x int) int {
	start, end := 0, len(arr)-1

	// Instead of recursing, a loop is used to perform the multiplication / "halving".
	// While the starting index of the search is less than or equal to the ending index,
	// binary search is performed.
	for start <= end {
		// Binary search works by essentially splitting the stored array in half to find a value.
		//
		// To find the middle index of the array, we find the average between the start and end
		// indices. Since both start and end are ints, we perform int division, which returns an int.
		middle := (start + end) / 2

		// There are three possible cases for each iteration of this loop.
		if arr[middle] == x {
			// 1. The middle element matches the element we're looking for, in which case x was found.
			return x
		} else if arr[middle] < x {
			// 2. The middle element is **less** than X. This means that X should be **right** of the middle.
			// Knowing this, the next iter should search the right half of the part of the array
			// currently searched.
			start = middle + 1
		} else {
			// 3. The middle element is **more** than X. This means that X should be **left** of the middle.
			// Knowing this, the next iter should search the left half of the part of the array
			// currently searched.
			end = middle - 1
		}
	}

	// X is not the array at all.
	return -1
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := 5
	pos := nonRecursiveSearch(arr, x)

	if pos != -1 {
		fmt.Printf("%v was found in the array at pos %d.\n", x, pos)
	} else {
		fmt.Printf("%v was not found in the array.\n", x)
	}
}

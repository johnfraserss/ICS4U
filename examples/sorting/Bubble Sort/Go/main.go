// Bubble sort
//
// Time complexity: O(N^2)
//
// Bubble sort "bubbles" up the array by comparing each array value with the one
// to the right of it.

package main

import "fmt"

func bubbleSort(arr []int) []int {
	length := len(arr)

	// Use the go equivalent of a do while loop to minimize
	// amount of bubbling required to verify it's sorted
	for swapped := true; swapped; {
		swapped = false

		// Loop through the entire array to perform swaps
		for i := 0; i < length-1; i++ {
			// If current val > value to the right, swap
			if arr[i] > arr[i+1] {
				tmp := arr[i]
				arr[i] = arr[i+1]
				arr[i+1] = tmp
				swapped = true
			}
		}
	}

	return arr
}

func main() {
	arr := []int{2, 1, 3, 4, 5}
	fmt.Println("Sorted array:", bubbleSort(arr)) // should equal [1, 2, 3, 4, 5]
}

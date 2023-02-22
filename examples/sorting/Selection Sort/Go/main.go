// Selection sort
//
// Time complexity: O(N^2)
//   - pretty slow, but not too bad on a small scale
//
// For each item of the array, selection sort loops through the array again
// to see if there are values less than the current minimum value.
//
// This is a very roundabout way to sort an array, but it works.
package main

import "fmt"

// selectionSort returns a sorted array given arr, an array of ints.
func selectionSort(arr []int) []int {
    length := len(arr)

    for i := 0; i < length; i++ {
        // Start off with the first element as the minium
        minIdx := i

        // Find the index for the minimum in the unsorted part of the array (i -> max)
        for j := i; j < length; j++ {
            // If a lower value exists, record its index as the new minIdx
            if arr[j] < arr[minIdx] {
                minIdx = j
            }
        }

        // Swap value at i with the minimum found previously
        tmp := arr[i]
        arr[i] = arr[minIdx]
        arr[minIdx] = tmp
    }

    return arr
}

func main() {
    arr := []int{4, 2, 9, 3, 1, 8, 6, 7, 5}
    fmt.Println("Sorted array:", selectionSort(arr))
}

// Insertion sort
//
// Time complexity: O(N^2)
//   - outperforms more complex sorts with small lists due to its low hidden constant value
//
// Insertion sort is an in-place sort that performs comparisons to sort an array.
// It starts at a given item in the list, and then looks to move it to the left so that
// the sublist to the left of the chosen index is sorted in ascending order.
//
// This means that by the time the sort moves the last item in the array, the array
// would already be sorted
package main

import "fmt"

func insertionSort(arr []int) []int {
    length := len(arr)

    for i := 1; i < length; i++ {
        // use current element as point of comparison
        key := arr[i]

        j := i - 1 // set idx to start comparison with
        // find a spot in the left of the arr where key fits
        for ; j >= 0 && arr[j] > key; j-- {
            arr[j+1] = arr[j]
        }

        // put key in right spot
        arr[j+1] = key
    }

    return arr
}

func main() {
    arr := []int{3, 2, 6, 4, 1, 5}
    fmt.Println("Sorted array:", insertionSort(arr))
}

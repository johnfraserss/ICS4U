/**
 * Binary search
 *
 * Time complexity: O(logN) - significantly faster than linear search at scale
 *
 * Given an element X in an array, binary search will find whether or not X
 * exists in that array.
 *
 * NOTE: this assumes the array is sorted
 */

package main

import "fmt"

// Recursively search sorted array arr for element x.
// Returns whether element x is within array arr or not.
func recursiveSearch(arr []int, x int) int {
    // Base condition if the slice from a prev iter is empty. Empty -> no result.
    if len(arr) == 0 {
        return -1
    }

    // Binary search works by essentially splitting the sorted array in half to
    // find a value. The middle point is calculated here.
    //
    // To find the middle index of the array, we add up the starting and ending
    // index, divide by 2, and round down (auto due to int division)
    start, end := 0, len(arr)
    middle := (start + end) / 2

    // There are three possible cases.

    if arr[middle] == x {
        // 1. The middle element matches element X.
        return x
    } else if arr[middle] < x {
        // 2. The middle element is **less** than X - search the right side
        // of the array since all those values are guaranteed to be greater
        // than the middle element.
        start = middle + 1

        // No such array can exist, so element x is not in array
        if start > end {
            return -1
        }

        // Recursively call function with new domain to search
        return recursiveSearch(arr[start:end], x)
    } else {
        // 3. The middle element is **more** than X - search the left side
        // of the array since all those values are guaranteed to be less
        // than the middle element.
        end = middle - 1

        // No such array can exist, so element x is not in array
        if start > end {
            return -1
        }

        // Recursively call function with new domain to search, adding 1
        // to the end index to ensure it's included in the slice
        return recursiveSearch(arr[start:end+1], x)
    }
}

func main() {
    arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    x := 9
    pos := recursiveSearch(arr, x)

    if pos != -1 {
        fmt.Printf("%v was found in the array at pos %d.\n", x, pos)
    } else {
        fmt.Printf("%v was not found in the array.\n", x)
    }
}

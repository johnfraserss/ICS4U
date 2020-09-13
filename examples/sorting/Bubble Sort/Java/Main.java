/**
 * Bubble sort
 * 
 * Time complexity: O(N^2)
 * 
 * Bubble sort "bubbles" up the array by comparing each array value with the one
 * to the right of it.
 */

 // import java.util.Arrays to print out the array for testing
import java.util.Arrays;

class Main {
  /**
   * Sorts an array in ascending order
   * 
   * @param array the array to sort
   * @returns the sorted array
   */
  public static int[] bubbleSort(int[] array) {
    int length = array.length;
    boolean swapped;

    /**
     * Using a do while loop to minimize the amount of bubbling needed to verify
     * that an array is sorted
     */
    do {
      swapped = false;
      // loop through the entire array to perform swaps
      for (int i = 0; i < length-1; i++) {
        // if the current value is greater than the value to the right of it, swap
        if (array[i] > array[i + 1]) {
          int tempValue = array[i];
          array[i] = array[i + 1];
          array[i + 1] = tempValue;
          swapped = true;
        }
      }
    } while (swapped);
    return array;
  }

  public static void main(String[] args) {
    // driver code for testing
    int[] array = {2, 1, 3, 4, 5};
    System.out.println(Arrays.toString(bubbleSort(array)));
  }
}
/**
 * Selection sort
 * 
 * Time complexity: O(N^2)
 *   - pretty slow, but not too bad on a small scale
 * 
 * For each item of the array, selection sort loops through the array again
 * to see if there are values less than the current minimum value.
 * 
 * This is a very roundabout way to sort an array, but it works. 
 */

// import java.util.Arrays to print out the array for testing
import java.util.Arrays;

class Main {
  /**
   * Sorts a given array
   * 
   * @param array the array to sort
   * @returns the sorted array
   */
  public static int[] selectionSort(int[] array) {
    // storing the length of the array here
    int length = array.length;

    /**
     * Loop through the array twice (nested loops)
     * 
     * This will check each item in the array against the rest to see if they're
     * lesser than the current minimum value, if they are, we swap their positions
     * in the array since we're sorting by ascending order and continue looking for
     * items of lesser value to swap
     */
    for (int i = 0; i < length; i++) {
      // we hold the minimum value's index here
      int minimumIndex = i;
      for (int j = i + 1; j < length; j++) {
        // if a value lesser than the current value of array[minimumIndex] is found, we
        // swap it
        if (array[j] < array[minimumIndex])
          minimumIndex = j;
      }

      // if the minimumIndex isn't the active one
      if (minimumIndex != i) {
        // swap the value at the minimumIndex to the front of the array
        int tempValue = array[i];
        array[i] = array[minimumIndex];
        array[minimumIndex] = tempValue;
      }
    }
    return array;
  }

  public static void main(String[] args) {
    // driver code for testing
    int[] array = { 2, 1, 3, 4, 5 };
    System.out.println(Arrays.toString(selectionSort(array)));
  }
}
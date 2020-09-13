
/**
 * Insertion sort
 * 
 * Time complexity: O(N^2)
 *   - outperforms more complex sorts with small lists due to its low hidden constant value
 * 
 * Insertion sort is an in-place sort that performs comparisons to sort an array.
 * It starts at a given item in the list, and then looks to move it to the left so that 
 * the sublist to the left of the chosen index is sorted in ascending order.
 * 
 * This means that by the time the sort moves the last item in the array, the array
 * would already be sorted
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
  public static int[] insertionSort(int[] array) {
    int length = array.length;
    for (int i = 1; i < length; i++) {
      // use the current array item as a key to compare
      int key = array[i];

      // set index to start comparison with
      int j = i - 1;
      // find a spot in the left sublist for the key to go
      while (j >= 0 && array[j] > key) {
        array[j + 1] = array[j];
        j = j - 1;
      }
      // set the key in the right spot
      array[j + 1] = key;
    }
    return array;
  };

  public static void main(String[] args) {
    // driver code for testing
    int[] array = { 2, 1, 3, 4, 5 };
    System.out.println(Arrays.toString(insertionSort(array)));
  }
}

/**
 * Binary Search
 *
 * Time complexity: O(logN) - significantly faster than linear search at scale
 *
 * Given an element X in an array, binary search will find whether or not X
 * exists in that array
 *
 * NOTE: this assumes the array is sorted
 */


class Recursive {
  /**
   * Recursively searches the array for the element x
   * 
   * @param array sorted array to search in
   * @param x the element to find
   * @param start starting index of search
   * @param end ending index of search
   * @return does the element x exist in the array
   */
  public static boolean recursiveSearch(int[] array, int x, int start, int end) {
    /**
     * Base condition if the starting index is greater than the ending index then
     * the search is complete with no results (you can't start at an index greater
     * than the end of the searchable array)
     */
    if (start > end)
      return false;

    /**
     * Binary search works by essentially splitting the sorted array in half to find
     * a value, we split the array in half here
     *
     * To find the middle index of the array, we add up the starting and ending
     * index, divide that by 2, and round down (Java integer division rounds down automatically)
     */
    int middleOfArray = (start + end) / 2;

    /**
     * Before we do anything else, let's see if the middle item in the array matches
     * the one we're searching for
     */
    if (array[middleOfArray] == x)
      return true;

    /**
     * Now for some actual searching There are two possible cases here:
     *
     * 1. The element in the middle is greater than `X` - search on the left side of
     * the array since all those values are guaranteed to be less than
     * `array[middleOfArray]` 
     * 2. The element in the middle is less than `X` - search
     * on the right side of the array since all those values are guaranteed to be
     * greater than `array[middleOfArray]`
     */

    if (array[middleOfArray] > x) {
      // we recursively call the binary search function with a new domain to search
      return recursiveSearch(array, x, start, middleOfArray - 1);
    } else {
      // we recursively call the binary search function with a new domain to search
      return recursiveSearch(array, x, middleOfArray + 1, end);
    }
  }

  public static void main(String[] args) {
    // driver code for testing
    int[] array = {1, 3, 5, 8, 10};
    int x = 8;

    if(recursiveSearch(array, x, 0, array.length - 1)) System.out.println(x + " was found in the array");
    else System.out.println(x + " was not found in the array");
  }
}
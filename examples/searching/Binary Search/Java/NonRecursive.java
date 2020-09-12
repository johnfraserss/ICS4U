
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


class NonRecursive {
  /**
   * Recursively searches the array for the element x
   * 
   * @param array sorted array to search in
   * @param x     the element to find
   * @return does the element x exist in the array
   */
  public static boolean nonRecursiveSearch(int[] array, int x) {
    int start = 0, end = array.length - 1;

    /**
     * Instead of recursing, we use a while loop to perform the multiplication While
     * the starting index of the search is less than or equal to the ending index,
     * we perform binary search
     */
    while (start <= end) {
      /**
       * Binary search works by essentially splitting the sorted array in half to find
       * a value, we split the array in half here
       *
       * To find the middle index of the array, we add up the starting and ending
       * index, divide that by 2, and round down (Java does this automatically)
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
       * 1. The element in the middle is less than `X` - search on the right side of
       * the array since all those values are guaranteed to be greater than
       * `array[middleOfArray]` 
       * 2. The element in the middle is greater than `X` -
       * search on the left side of the array since all those values are guaranteed to
       * be less than `array[middleOfArray]`
       */
      else if (array[middleOfArray] < x) {
        start = middleOfArray + 1;
      } else {
        end = middleOfArray - 1;
      }
    }

    return false;
  }

  public static void main(String[] args) {
    // driver code for testing
    int[] array = { 1, 3, 5, 8, 10 };
    int x = 8;

    if (nonRecursiveSearch(array, x))
      System.out.println(x + " was found in the array");
    else
      System.out.println(x + " was not found in the array");
  }
}
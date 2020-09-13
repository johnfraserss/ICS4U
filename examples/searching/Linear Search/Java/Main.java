/**
 * Linear search (a.k.a. the easier one)
 * 
 * Time complexity: O(N) - slower at large scales
 * 
 * Linear search is pretty self-descriptive. Given an array, and an element X,
 * loop through the entire array in order until you find X.
 * 
 * The benefit of linear search is that the array doesn't have to be sorted
 * since you're looping through the entire array.
 */

class Main {

  /**
   * Searches the array for element x
   * 
   * @param array array to search in
   * @param x     the element to find
   * @return does the element x exist in the array
   */
  public static boolean linearSearch(int[] array, int x) {

    // set found to false by default, if nothing matches, found will be unchanged
    // and returned
    boolean found = false;

    /**
     * Using the for (item : array) loop to take each element of the array as `item`
     * and comparing it to x.
     * 
     * If it equals x, we set found to true, if none of them equal x, then the
     * function returns false.
     */
    for (int item : array) {
      if (item == x)
        found = true;
    }

    return found;
  }

  public static void main(String[] args) {
    // driver code for testing
    int[] array = { 1, 3, 5, 8, 10 };
    int x = 8;

    if (linearSearch(array, x))
      System.out.println(x + " was found in the array");
    else
      System.out.println(x + " was not found in the array");
  }
}
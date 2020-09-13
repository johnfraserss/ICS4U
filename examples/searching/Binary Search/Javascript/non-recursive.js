/**
 * Binary Search
 *
 * Time complexity: O(logN)
 *   - significantly faster than linear search at scale
 *
 * Given an element X in an array, binary search will find whether
 * or not X exists in that array
 *
 * NOTE: this assumes the array is sorted
 */

 /**
  * Recursively searches the array for the element x
  * @param {array} array sorted array to search in
  * @param {number} x the element to find
  * @returns {boolean} does the element x exist in the array
  */
const nonRecursiveSearch = (array, x) => {
  let start = 0,
    end = array.length - 1;

  /**
   * Instead of recursing, we use a while loop to perform the multiplication
   * While the starting index of the search is less than or equal to the ending index,
   * we perform binary search
   */
  while (start <= end) {

    /**
     * Binary search works by essentially splitting the sorted
     * array in half to find a value, we split the array in half here
     *
     * To find the middle index of the array, we add up the starting and ending
     * index, divide that by 2, and round down
     */
    let middleOfArray = Math.floor((start + end) / 2);

    /**
     * Before we do anything else, let's see if the
     * middle item in the array matches the one we're searching for
     *
     * NOTE: Javascript uses triple equality - if you're interested in why,
     * google to learn more about truthy/falsy in Javascript
     */
    if (array[middleOfArray] === x) return true;

   /**
    * Now for some actual searching
    * There are two possible cases here:
    *
    *    1. The element in the middle is less than `X`
    *       - search on the right side of the array since all those values are
    *         guaranteed to be greater than `array[middleOfArray]`
    *    2. The element in the middle is greater than `X`
    *       - search on the left side of the array since all those values are
    *         guaranteed to be less than `array[middleOfArray]`
    */ 
    else if (array[middleOfArray] < x)  {
      start = middleOfArray + 1;
    } else {
      end = middleOfArray - 1;
    }
  }

  return false;
};

// driver code for testing
const array = [1, 3, 5, 8, 10]; // side note: use const for objects you won't modify. It's good practice
const x = 8;

nonRecursiveSearch(array, x, 0, array.length - 1)
  ? console.log(`${x} was found in the array`)
  : console.error(`${x} was not found in the array`);

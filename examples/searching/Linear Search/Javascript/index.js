/**
 * Linear search (a.k.a. the easier one)
 * 
 * Time complexity: O(N)
 *   - slower at large scales
 * 
 * Linear search is pretty self-descriptive.
 * Given an array, and an element X, loop through the entire 
 * array in order until you find X. 
 * 
 * The benefit of linear search is that the array doesn't have to 
 * be sorted since you're looping through the entire array.
 */

/**
 * Searches the array for element x
 * @param {array} array array to search in
 * @param {any} x the element to find
 * @returns {boolean} does the element x exist in the array
 */
const linearSearch = (array, x) => {

  // set found to false by default, if nothing matches, found will be unchanged and returned
  let found = false;
  
  /**
   * Using the `forEach` array method to take each element of the array
   * as `item` and comparing it to x.
   * 
   * If it equals x, we set found to true, if none of them equal x,
   * then the function returns false.
   */
  array.forEach(item => {
    if (item === x) found = true;
  })

  return found;
}

// driver code for testing
const array = [1, 3, 5, 8, 10]; // side note: use const for objects you won't modify. It's good practice
const x = 8;

linearSearch(array, x)
  ? console.log(`${x} was found in the array`)
  : console.error(`${x} was not found in the array`);
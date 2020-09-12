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


/**
 * Sorts a given array
 * @param {array} array the array to sort
 * @returns {array} the sorted array
 */
const selectionSort = (array) => {
  // storing the length of the array here
  let length = array.length;

  /**
   * Loop through the array twice (nested loops)
   * 
   * This will check each item in the array against the rest to see if 
   * they're lesser than the current minimum value, if they are, we swap their
   * positions in the array since we're sorting by ascending order
   * and continue looking for items of lesser value to swap
   */
  for (let i = 0; i < length; i++) {
    // we hold the minimum value's index here
    let minimumIndex = i;
    for (let j = i+1; j < length; j++) {
      // if a value lesser than the current value of array[minimumIndex] is found, we swap it
      if (array[j] < array[minimumIndex]) minimumIndex = j;
    }

    // if the minimumIndex isn't the active one
    if (minimumIndex !== i) {
      // swap the value at the minimumIndex to the front of the array
      const tempValue = array[i];
      array[i] = array[minimumIndex];
      array[minimumIndex] = tempValue;
    }
  }
  return array;
}

// driver code to test
const array = [2, 1, 3, 4, 5];
console.log(selectionSort(array)); // should equal [1, 2, 3, 4, 5]
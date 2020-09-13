/**
 * Bubble sort
 * 
 * Time complexity: O(N^2)
 * 
 * Bubble sort "bubbles" up the array by comparing each
 * array value with the one to the right of it.
 */


/**
 * Sorts an array in ascending order
 * @param {array} array the array to sort
 * @returns {array} the sorted array
 */
const bubbleSort = (array) => {
  const length = array.length;
  let swapped;

  /**
   * Using a do while loop to minimize the amount 
   * of bubbling needed to verify that an array
   * is sorted
   */
  do {
      swapped = false;
      // loop through the entire array to perform swaps
      for (let i = 0; i < length; i++) {
        // if the current value is greater than the value to the right of it, swap
        if (array[i] > array[i + 1]) {
            let tempValue = array[i];
            array[i] = array[i + 1];
            array[i + 1] = tempValue;
            swapped = true;
        }
      }
  } while (swapped);
  return array;
};

// driver code to test
const array = [2, 1, 3, 4, 5];
console.log(bubbleSort(array)); // should equal [1, 2, 3, 4, 5]
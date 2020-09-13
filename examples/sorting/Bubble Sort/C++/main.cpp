/**
 * Bubble sort
 * 
 * Time complexity: O(N^2)
 * 
 * Bubble sort "bubbles" up the array by comparing each array value with the one
 * to the right of it.
 *
 * NOTE: confused about what a vector is? Look here: https://www.geeksforgeeks.org/vector-in-cpp-stl/ 
 */


#include <bits/stdc++.h> 
#include <vector> 
using namespace std; 

/**
  * Sorts a vector in ascending order
  * 
  * @param vector<int> the vector to sort
  * @returns the sorted vector
  */
vector<int> bubbleSort(vector<int> array) {
  int length = array.size();
  bool swapped;

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

int main() {
  // driver code for testing
  vector<int> array{2, 1, 3, 4, 5};
  vector<int> sortedArray = bubbleSort(array);
  for (int item : sortedArray) {
    cout << item << " ";
  }
}
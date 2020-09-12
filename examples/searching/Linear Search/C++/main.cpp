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
 *
 * NOTE: confused about what a vector is? Look here: https://www.geeksforgeeks.org/vector-in-cpp-stl/
 */


#include <bits/stdc++.h> 
#include <vector> 
using namespace std; 

/**
  * Searches the vector for element x
  * 
  * @param vector<int> vector to search in
  * @param x     the element to find
  * @return does the element x exist in the vector
  */
bool linearSearch(vector<int> array, int x) {

  // set found to false by default, if nothing matches, found will be unchanged
  // and returned
  bool found = false;

  /**
    * Using the for (item : vector) loop to take each element of the vector as `item`
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

int main() {
  // driver code for testing
  vector<int> array{1, 3, 5, 8, 10};
  int x = 8;

  if(linearSearch(array, x)) cout << x << " was found in the array";
  else cout << x << " was not found in the array";
}
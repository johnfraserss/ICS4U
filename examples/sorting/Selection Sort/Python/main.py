'''
Selection sort
 
Time complexity: O(N^2)
  - pretty slow, but not too bad on a small scale
 
For each item of the array, selection sort loops through the array again
to see if there are values less than the current minimum value.
 
This is a very roundabout way to sort an array, but it works. 
'''


def selectionSort(array):
  '''
  Sorts a given array

      Parameters:
        array (array): the array to sort

      Returns:
        (array): the sorted array
  '''

  # storing the length of the array here
  length = len(array)

  '''
  Loop through the array twice (nested loops)
  
  This will check each item in the array against the rest to see if 
  they're lesser than the current minimum value, if they are, we swap their
  positions in the array since we're sorting by ascending order
  and continue looking for items of lesser value to swap
  '''
  for i in range(0, length-1):
    # we hold the minimum value's index here
    minimumIndex = i
    for j in range(i+1, length-1):
      # if a value lesser than the current value of array[minimumIndex] is found, we swap it
      if (array[j] < array[minimumIndex]): minimumIndex = j

    # if the minimumIndex isn't the active one
    if (minimumIndex != i):
      # swap the value at the minimumIndex to the front of the array
      tempValue = array[i]
      array[i] = array[minimumIndex]
      array[minimumIndex] = tempValue
  return array

# driver code to test
array = [2, 1, 3, 4, 5]
print(selectionSort(array)) # should equal [1, 2, 3, 4, 5]
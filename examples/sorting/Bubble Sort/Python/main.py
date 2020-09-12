'''
Bubble sort
 
Time complexity: O(N^2)
 
Bubble sort "bubbles" up the array by comparing each
array value with the one to the right of it.
'''


def bubbleSort(array):
  '''
  Sorts an array in ascending order

      Parameters:
        array (array): the array to sort

      Returns:
        (array): the sorted array
  '''
  length = len(array)
  swapped = ''

  '''
  Using the python equivalent of a do while loop to minimize the amount 
  of bubbling needed to verify that an array is sorted
  '''
  while True:
      swapped = False
      # loop through the entire array to perform swaps
      for i in range(0, length-1):
        # if the current value is greater than the value to the right of it, swap
        if (array[i] > array[i + 1]):
            tempValue = array[i]
            array[i] = array[i + 1];
            array[i + 1] = tempValue;
            swapped = True
      if (swapped):
        break
  return array

# driver code to test
array = [2, 1, 3, 4, 5]
print(bubbleSort(array)) # should equal [1, 2, 3, 4, 5]
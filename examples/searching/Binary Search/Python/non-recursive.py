'''
 Binary Search
 
 Time complexity: O(logN)
    - significantly faster than linear search at scale
 
  Given an element X in an array, binary search will find whether
  or not X exists in that array
 
  NOTE: this assumes the array is sorted
'''

import math

def nonRecursiveSearch(array, x):
    '''
    Recursively searches the array for the element x
  
    Parameters:
        array (array): sorted array to search in
        x (any): the element to find

    Returns:
        boolean: does the element x exist in the array
    '''
  
    start = 0
    end = len(array) - 1

    '''
    Instead of recursing, we use a while loop to perform the multiplication
    While the starting index of the search is less than or equal to the ending index,
    we perform binary search
    '''
    while (start <= end):

        '''
        Binary search works by essentially splitting the sorted
        array in half to find a value, we split the array in half here
    
        To find the middle index of the array, we add up the starting and ending
        index, divide that by 2, and round down
        '''
        middleOfArray = math.floor((start + end) / 2)

        '''
        Before we do anything else, let's see if the
        middle item in the array matches the one we're searching for
        '''
        if (array[middleOfArray] == x): 
            return True
    
        # Now for some actual searching
        # There are two possible cases here:

        #   1. The element in the middle is less than X
        #     - search on the right side of the array since all those values are
        #       guaranteed to be greater than array[middleOfArray]    
        #   2. The element in the middle is greater than X
        #     - search on the left side of the array since all those values are
        #       guaranteed to be less than array[middleOfArray]
        
        elif (array[middleOfArray] < x):
            start = middleOfArray + 1
          
        else:
            end = middleOfArray - 1

# driver code for testing
array = [1, 3, 5, 8, 10]
x = 8

print('{x} was found in the array'.format(x = x)) if nonRecursiveSearch(array, x) else print('{x} was not found in the array'.format(x = x))

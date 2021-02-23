'''
Linear search (a.k.a. the easier one)
 
Time complexity: O(N)
  - slower at large scales
 
Linear search is pretty self-descriptive.
Given an array, and an element X, loop through the entire 
array in order until you find X. 
 
The benefit of linear search is that the array doesn't have to 
be sorted since you're looping through the entire array.
'''

def linearSearch(array, x):
    '''
    Searches the array for element x

    Parameters:
        array (array): array to search in
        x (any): the element to find
      
    Returns:
        boolean: does the element x exist in the array
    '''

    # set found to false by default, if nothing matches, found will be unchanged and returned
    found = False
  
    '''
    Using the `for item in array` loop to take each element of the array
    as `item` and comparing it to x.
   
    If it equals x, we set found to true, if none of them equal x,
    then the function returns false.
    '''
    for item in array:
        if (item == x):
            found = True

    return found

# driver code for testing
array = [1, 3, 5, 8, 10] # side note: use const for objects you won't modify. It's good practice
x = 8

print('{x} was found in the array'.format(x = x)) if linearSearch(array, x) else print('{x} was not found in the array'.format(x = x))

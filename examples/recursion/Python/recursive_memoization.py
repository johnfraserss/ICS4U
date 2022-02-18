'''
Recursive example using memoization with Fibonacci
'''

def fibonacci_memoize(n, database):
    '''
    Helper function to calculates the nth Fibonacci number
    where 0 is the 0th, and 1 is the 1st using memoization
    
    Parameters
    ----------
    n : int
        The position in the Fibonacci sequence to calculate
    db : dict
        The dictionary to hold all previously calculated Fibonacci values
    
    Returns
    -------
    int
        The nth Fibonacci number
    '''
    if n == 0:
        return 0
    elif n == 1:
        return 1
    elif n in db:
        return db[n]
    else:
        calculated_value = fibonacci_memoize( n - 1, db) + fibonacci_memoize( n - 2, db)
        db[n] = calculated_value
        return calculated_value
    
def fibonacci(n):
    '''
    Function to use to calculate the nth Fibonacci number
    
    Parameters
    ----------
    n : int
        The position in the Fibonacci sequence to calculate
    
    Returns
    -------
    int
        The nth Fibonacci number
    
    '''
    database = {}
    response = fibonacci_memoize(n, database)
    database.clear()
    return response

to_find = 11
answer = fibonacci(to_find)
print(f'The {to_find}th Fibonacci number is {answer}')
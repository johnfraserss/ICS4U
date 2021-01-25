'''
STRING INTERPOLATION
'''

name = "John Doe"
age = 20

#These are different ways you can combine variables together


#Concatenation (The one you are most familiar with)
string = "Hello! My name is " + name + " and I am " + str(age) + " years old."
print(string)

#f-strings (Works only for Python 3.6+)
# placing 'f' before initializing a string, denotes the start of an f-string
string = f"Hello! My name is {name} and I am {age} years old."
print(string)

#Formatting (Works for all versions of Python)
# {} denotes a placeholder for a variable
string = "Hello! My name is {} and I am {} years old.".format(name, age) #Variables are put in order at which they appear
string = "Hello! My name is {x} and I am {y} years old.".format(x=name, y=age) #Can also be written using keyword arguments (kwargs)
print(string)

#% String Formatting (Works for all versions of Python)
# %s denotes a placeholder for a variable
string = "Hello! My name is %s and I am %s years old." % (name, age) #Variables are put in order at which they appear
print(string)



'''
Concatenation CONS

- Lack of Readability
- Have to convert all variables to strings
- Difficult to type if you have more variables
- Lengthy

Concatenation PROS 

- Works on all versions of Python
'''



'''
Formatting CONS 
- Lengthy
- Difficult to type if you have more variables

Formatting PROS 

- Decent Readability
- Don't have to convert variables into strings
- Works on all versions of Python
'''



'''
%s String Formatting CONS 
- Lengthy
- Difficult to type if you have more variables
- Poor Readability

%s String Formatting PROS 

- Don't have to convert variables into strings
- Works on all versions of Python
'''


'''
f-string CONS 
- Only works on Python versions 3.6 and up

f-string PROS 

- Excellent Readability
- Don't have to convert variables into strings
- Not Lengthy
- Easy to type even with multiple variables
'''


#Conclusion
# - Formatting is the best to use if you are on a python version lower than 3.6
# - F-Strings are the best to use if you are on python 3.6 or above
# - Concatenation should be avoided when possible
# - %s String Formatting can be used but should be avoided if not needed

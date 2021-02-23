#-----------------------------------------------------------------------------
# Name:        Class and Documentation i.e. docstrings (python.py)
# Purpose:     To provide an example of how to create a class and the associated
#				docstrings in Python for classes.
#
# References: 	This program uses the NumPy/SciPy style of documentation as found
#				here: https://numpydoc.readthedocs.io/en/latest/format.html with
#				some minor modifications based on Python 3 function annotations
#				(the -> notation).
#
# Author:      Mr. Seidel
# Created:     17-Sep-2018
# Updated:     13-Apr-2020 (updated methods to be "get" instead of "print" for better encapsulation)
#-----------------------------------------------------------------------------

from book import Book

bookOne = Book("Terry Pratchett", "Guards! Guards!", 5.99)
bookTwo = Book("Robert Jordan", "The Eye of the World", 8.99)

print(bookOne.getAuthor());
print(bookOne.getPrice());
print(bookOne.getTitle());
bookOne.increasePrice(1.00);
print(bookOne.getPrice());

print(bookTwo.getAuthor());
print(bookTwo.getPrice());
print(bookTwo.getTitle());
bookTwo.increasePrice(-6.00);
print(bookTwo.getPrice());
try:
    bookTwo.increasePrice("Hello!")
except TypeError:
    print("Raised a TypeError as expected")
except Exception:
    print("Some other Error popped up?")

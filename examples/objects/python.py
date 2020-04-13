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

class Book():
	'''
	A book object that hold the price, author, and title of book
	
	Attributes
	----------
	price : float
		The price of the book in dollars and cents (example format ###.##)
	author : str
		The full name of the author of the book
	title : str
		The full title of the book
		
	Methods
	-------
	# note, do not list private methods in this section.  Do not include this line
	getAuthor() -> str
		Prints the name of the author to the console
	getPrice() -> float
		Prints the price of the book to the console
	getTitle() -> str
		Prints the name of the book to the console
	increasePrice(increase : float) -> None
		Attempts to increase the price by a float value
		
	'''
	
	def __init__(self, author, title, price=0.00):
		'''
		Constructor to build a book object
		
		
		Parameters
		----------
		author : str
			The name of the author of the book
		price : float, optional
			The initial price of the book
			The default price of a book is 0.00 if none is entered
		title : str
			The title of the book
			
		..note:: that the "self" parameter is not listed in this section.
			
		
			
		'''
		
		self.author = author
		self.price = price
		self.title = title
		
		
	def getAuthor(self) -> str:
		'''
		Returns the author of the book
		
		Returns
		-------
		The author of the book
		
		'''
		return self.author
	
	
	def getPrice(self) -> float:
		'''
		Returns the price of the book
		
		Returns
		-------
		The price of the book as a float
				
		'''
		return float(self.price)
		
		
	def getTitle(self) -> str:
		'''
		Returns the title of the book to the console
		
		Returns
		-------
		The title of the book
		'''
		return self.title
	
	
	def increasePrice(self, increase : float) -> None:
		'''
		Attempts to increase the price of the book
		
		This will add the value of 'increase' to self.price.  If there is any issues, 
		it will return False if it attempts to lower the price below zero.
		
		Parameters
		----------
		increase : float
			The value to increase the price by.  This value can be negative;
			however, it will never lower the value below zero.  If this happens
			the function will return False.
				
		Returns
		-------
		bool
			True if the method was successful
			False if the method attempted to bring the value below zero
		
		Raises
		------
		AttributeException
			If the attribute is non-existant and Python cannot add to it
		TypeError
			If the value of increase is not a float and it is trying to
			be added to self.price : float.
		
		'''
		
		if self.price + increase > 0:
			self.price += increase
			return True
		return False

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
	print ("Raised a TypeError as expected")
except Exception:
	print ("Some other Error popped up?")

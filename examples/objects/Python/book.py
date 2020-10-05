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
	
	
	def increasePrice(self, increase : float) -> bool:
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

